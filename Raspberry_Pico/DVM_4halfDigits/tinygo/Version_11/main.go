/*
Dual Slope ADC Project for Raspberry Pi Pico - RP2040 - Version 11
Gustavo Murta 2025_09_27
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

	ticks, nanoSECperTicks = uint32(0), uint32(0)
	ACfrequency            = uint32(60)                           // AC main frequency 50 Hz or 60 Hz
	ACperiod               = uint32(1e9 / ACfrequency)            // AC period in nanoseconds
	voltageLSB             = float32(2.5) / float32(ACperiod/5.0) // voltage LSB calculation 750 nV for 60 Hz

	calibrate = uint32(130e3) // calibration timer value to adjust voltage accuracy (130 us in nanoseconds)

	startPeriod, endPeriod     = uint32(0), uint32(0)
	autoZeroPeriod             = uint32(0)
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

func processorClock() uint32 {
	cpuFreq := machine.CPUFrequency() // Get CPU frequency
	fmt.Println("CPU Frequency:", int(cpuFreq), "Hz")
	nanoSECperTicks := 1e9 / cpuFreq // Calculate nanoseconds per ticks
	fmt.Println("Nanoseconds per tick:", nanoSECperTicks)
	return nanoSECperTicks
}

// Read current SysTick value
func readSysTick() uint32 {
	return rp.PPB.GetSYST_CVR_CURRENT() // Read current SysTick value (24 bits)
}

// Delay with Systick Timer - processor clock 200 MHz - max: 83 ms (83,886,080 ns) - multiples of 5 ns
func delaynanoSec(ns uint32) {
	ticks = ns / 5                 //  5 nanoSECperTicks
	startSystick := readSysTick()  // Read current systick timer
	target := startSystick - ticks // Calculate target delay
	for readSysTick() > target {   // Wait target timeout
	}
}

func testPointA() { // used only for testing with logic analyzer
	tpA.High()       // Set GPIO13 High
	delaynanoSec(25) // delay 25 ns
	tpA.Low()        // Set GPIO13 Low
}

// Phase I - Reset Integrator
// 74HC4051 - set Address 7 => Reset Integrator
func resetIntegrator() {
	mux_S0.High()     // MUX bit0
	mux_S1.High()     // MUX bit1
	delaynanoSec(5e6) // Delay to Capacitor discharge 5 ms - do not change
}

// Phase II - Integrate V Input
// 74HC4051 - set Address 5 => Read VIN
func integVIN() {
	mux_S0.High()                      // MUX bit0
	mux_S1.Low()                       // MUX bit1
	delaynanoSec(ACperiod + calibrate) // Read VIN 16.666 ms - do not change
}

// Phase III - Integrate - VREF
// 74HC4051 - set Address 6 => Read -VREF
func integVREF() {
	mux_S0.Low()                // MUX bit0
	mux_S1.High()               // MUX bit1
	startPeriod = readSysTick() // Read SysTick before start integration
	delaynanoSec(20e6)          // Read VREF 20 ms - VREF time must be bigger than VIN time
}

// Phase IV - GND at Input - Auto Zero
// 74HC4051 - set Address 4
func gndInput() {
	mux_S0.Low()                       // MUX bit0
	mux_S1.Low()                       // MUX bit1
	delaynanoSec(ACperiod + calibrate) // auto Zero at VIN 16.666 ms - do not change
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
	fmt.Printf("Dual Slope ADC RP2040 V11 \n")

	fmt.Printf("AC main period: %d ns\n", ACperiod)
	fmt.Printf("LSB Voltage: %.9f V\n", voltageLSB)

	// MUX pins configuration - as output
	mux_S0.Configure(machine.PinConfig{Mode: machine.PinOutput})
	mux_S1.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Test Point A configuration
	tpA.Configure(machine.PinConfig{Mode: machine.PinOutput}) // Set GPIO13 as output

	// Comparator pin configuration with interrupt
	comparatorPin.Configure(machine.PinConfig{Mode: machine.PinInput}) // Set GPIO15 as input - interrupt pin

	processorClock() // Display processor clock and nanoseconds per tick

	enableSysTick()                                                 // enable SysTick timer
	comparatorPin.SetInterrupt(machine.PinRising, adcComparatorISR) // enable PinRising interrupt

	for {
		autoZeroADC()

		if voltageRead {
			comparatorPin.SetInterrupt(machine.PinRising, nil) // Disable interrupt to avoid multiple triggers

			autoZeroPeriod = startPeriod - endPeriod                        // period of VIN = 0 Volts integration
			delaynanoSec(8.3e6)                                             // delay of 8.3 ms between measurements - do not change
			comparatorPin.SetInterrupt(machine.PinRising, adcComparatorISR) // enable PinRising interrupt
			voltageRead = false
		}

		dualSlopeADC()

		if voltageRead {
			comparatorPin.SetInterrupt(machine.PinRising, nil) // Disable interrupt to avoid multiple triggers

			vrefPeriod = startPeriod - endPeriod // period of VREF integration
			// vrefPeriodTime = vrefPeriod * 5      // VREF Period time * 5 ns per tick

			fmt.Printf("Calibration: %d // ", calibrate)
			fmt.Printf("Auto Zero Period: %d ticks // ", autoZeroPeriod)
			voltageZero = float32(autoZeroPeriod) * voltageLSB // Auto zero voltage calculation 7.5e-7
			fmt.Printf("Voltage Auto Zero: %.3f V # ", voltageZero)

			// fmt.Printf("VREF Period: %d ticks // ", vrefPeriod)
			// fmt.Printf("VREF Period Time: %d ns # ", vrefPeriodTime)

			// Calculation =>  VIN Voltage  = VIN period * (2.500 V / 2e6 ticks)

			vinPeriod = (vrefPeriod - autoZeroPeriod) // period of VIN integration
			// vinPeriodTime = vinPeriod * 5             // VIN Period time * 5 ns per tick
			fmt.Printf("VIN Period: %d ticks // ", vinPeriod)
			voltage = float32(vinPeriod) * voltageLSB // VIN voltage calculation 7.5e-7
			fmt.Printf("VIN Voltage: %.3f V\n", voltage)

			delaynanoSec(8.3e6)                                             // delay of 8.3 ms -  10 measurements/second
			comparatorPin.SetInterrupt(machine.PinRising, adcComparatorISR) // enable PinRising interrupt
			voltageRead = false
		}
	}
}


