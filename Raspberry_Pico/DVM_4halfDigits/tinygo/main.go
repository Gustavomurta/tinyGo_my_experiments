/*
Dual Slope ADC Example for Raspberry Pi Pico - RP2040 - Version 10
Gustavo Murta 2025_09_15
tinygo version 0.39.0 windows/amd64 (using go version go1.25.1 and LLVM version 19.1.2)
C:\Users\jgust\tinygo\programas\Raspberry_Pico\Dual_Slope_ADC_Pico
tinygo flash -target pico main.go

74HC4051 S0	= GPIO10
74HC4051 S1	= GPIO11
VOUT Q1	= GPIO15  (interrupt pin) - Comparator Output

A4 - GND input - Auto Zero
A5 - V input integration
A6 - V Reference -2.500 V integration
A7 - Reset Integrator
*/

package main

import (
	"device/rp"
	"fmt"
	"machine"
	"time"
	//"tinygo.org/x/tinygo/src/machine"
)

var (
	mux_S0        = machine.GPIO10 // mux adress bit 0
	mux_S1        = machine.GPIO11 // mux adress bit 1
	comparatorPin = machine.GPIO15 // Comparator Output - Interrupt Pin
	tpA           = machine.GPIO13 // Test Point A

	startPeriod, endPeriod = uint32(0), uint32(0)

	autoZeroPeriod = uint32(0)

	voltageRead                = false
	vrefPeriod, vrefPeriodTime = uint32(0), uint32(0)
	vinPeriod, vinPeriodTime   = uint32(0), uint32(0)
	voltage, voltageZero       = float32(0), float32(0)
)

const MAX_RELOAD = 0xFFFFFF // Maximum reload value (24 bits): 16,777,216 ticks 0xFFFFFF

// Enable SysTick counter
func enableSysTick() {
	rp.PPB.SetSYST_RVR_RELOAD(MAX_RELOAD) // Set maximum reload value (24 bits)
	rp.PPB.SetSYST_CSR_CLKSOURCE(1)       // SysTick clock source: processor clock
	rp.PPB.SetSYST_CSR_ENABLE(1)          // Enable SysTick counter
	rp.PPB.SetSYST_CVR_CURRENT(0)         // Clear current value (24 bits)
}

// Read current SysTick value
func readSysTick() uint32 {
	return rp.PPB.GetSYST_CVR_CURRENT() // Read current SysTick value (24 bits)
}

// Delay with Systick Timer - processor clock 200 MHz - max: 83 ms
func delaySystick(ticks uint32) {
	startSystick := readSysTick()  // Read current systick timer
	target := startSystick - ticks // Calculate target delay
	for readSysTick() > target {   // Wait target timeout
	}
}

func testPointA() { // used only for testing with logic analyzer
	tpA.High()      // Set GPIO13 High
	delaySystick(5) // 5 ticks = 25 ns
	tpA.Low()       // Set GPIO13 Low
}

// Phase I - Reset Integrator
// 74HC4051 - set Address 7 => Reset Integrator
func resetIntegrator() {
	mux_S0.High()            // MUX bit0
	mux_S1.High()            // MUX bit1
	delaySystick((10e6 / 5)) // Delay to Capacitor discharge 5 ms = ((5e6 / 5) - 125)
}

// Phase II - Integrate V Input
// 74HC4051 - set Address 5 => Read VIN
func integVIN() {
	mux_S0.High()                  // MUX bit0
	mux_S1.Low()                   // MUX bit1
	delaySystick((10e6 / 5) - 240) // Read VIN 10 ms = ((10e6 / 5) - 240)
}

// Phase III - Integrate - VREF
// 74HC4051 - set Address 6 => Read -VREF
func integVREF() {
	mux_S0.Low()                // MUX bit0
	mux_S1.High()               // MUX bit1
	startPeriod = readSysTick() // Read SysTick before start integration
	delaySystick((11e6 / 5))    // Read VREF 11 ms - VREF time must be bigger than VIN time
}

// Phase IV - GND at Input - Auto Zero
// 74HC4051 - set Address 4
func gndInput() {
	mux_S0.Low()           // MUX bit0
	mux_S1.Low()           // MUX bit1
	delaySystick(10e6 / 5) // auto Zero at VIN 10 ms = minimum time for zeroing
}

func autoZeroADC() { // read Zero Volts value
	rp.PPB.SetSYST_CVR_CURRENT(0) // Clear current value (24 bits)
	resetIntegrator()             // Reset Integrator
	gndInput()                    // auto Zero Integrator
	integVREF()                   // Integrate V REF
}

func dualSlopeADC() {
	rp.PPB.SetSYST_CVR_CURRENT(0) // Clear current value (24 bits)
	resetIntegrator()             // Reset Integrator
	integVIN()                    // Integrate V Input
	integVREF()                   // Integrate V REF
}

func adcComparatorISR(pin machine.Pin) {
	endPeriod = readSysTick() // Read SysTick after rising interrupt
	voltageRead = true
}

func main() {

	machine.InitSerial()        // Initialize serial for debug output
	time.Sleep(3 * time.Second) // Sleep to catch prints on serial
	fmt.Printf("Dual Slope ADC RP2040 V10 \n")

	// MUX pins configuration - as output
	mux_S0.Configure(machine.PinConfig{Mode: machine.PinOutput})
	mux_S1.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Test Point A configuration
	tpA.Configure(machine.PinConfig{Mode: machine.PinOutput}) // Set GPIO13 as output

	// Comparator pin configuration with interrupt
	comparatorPin.Configure(machine.PinConfig{Mode: machine.PinInput}) // Set GPIO15 as input - interrupt pin

	cpuFreq := machine.CPUFrequency() // Get CPU frequency
	fmt.Println("CPU Frequency:", int(cpuFreq), "Hz")

	nanoSECperTicks := 1e9 / cpuFreq // Calculate nanoseconds per ticks
	fmt.Println("Nanoseconds per tick:", nanoSECperTicks)

	enableSysTick()                                                 // enable SysTick timer
	comparatorPin.SetInterrupt(machine.PinRising, adcComparatorISR) // enable PinRising interrupt

	for {
		autoZeroADC()

		if voltageRead {
			comparatorPin.SetInterrupt(machine.PinRising, nil) // Disable interrupt to avoid multiple triggers

			autoZeroPeriod = startPeriod - endPeriod                        // period of VIN = 0 Volts integration
			time.Sleep(time.Millisecond * 20)                               // delay of 20 ms between measurements
			comparatorPin.SetInterrupt(machine.PinRising, adcComparatorISR) // enable PinRising interrupt
			voltageRead = false
		}

		dualSlopeADC()

		if voltageRead {
			comparatorPin.SetInterrupt(machine.PinRising, nil) // Disable interrupt to avoid multiple triggers

			vrefPeriod = startPeriod - endPeriod // period of VREF integration
			//vrefPeriodTime = vrefPeriod * 5      // VREF Period time * 5 ns per tick

			fmt.Printf("Auto Zero Period: %d ticks // ", autoZeroPeriod)
			voltageZero = float32(autoZeroPeriod) * 1.25e-6 // Auto zero voltage calculation
			fmt.Printf("Voltage Auto Zero: %.3f V # ", voltageZero)

			fmt.Printf("VREF Period: %d ticks // ", vrefPeriod)
			//fmt.Printf("VREF Period Time: %d ns # ", vrefPeriodTime)

			// Calculation =>  VIN Voltage  = VIN period * (2.500 V / 2e6 ticks)

			vinPeriod = (vrefPeriod - autoZeroPeriod) // period of VIN integration
			vinPeriodTime = vinPeriod * 5             // VIN Period time * 5 ns per tick
			voltage = float32(vinPeriod) * 1.25e-6    // VIN voltage calculation
			fmt.Printf("VIN Voltage: %.3f V\n", voltage)

			time.Sleep(time.Millisecond * 125)                              // delay 125 ms => 5 measurements/second
			comparatorPin.SetInterrupt(machine.PinRising, adcComparatorISR) // enable PinRising interrupt
			voltageRead = false
		}
	}
}




