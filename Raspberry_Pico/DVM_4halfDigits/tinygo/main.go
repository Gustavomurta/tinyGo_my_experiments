/*
Dual Slope ADC Example for Raspberry Pi Pico - RP2040 - Version 9
Gustavo Murta 2025_09_09 - MIT license
tinygo version 0.38.0 windows/amd64 (using go version go1.24.2 and LLVM version 19.1.2)
C:\Users\jgust\tinygo\programas\Raspberry_Pico\Dual_Slope_ADC_Pico
tinygo flash -target pico main.go

https://github.com/tinygo-org/tinygo/tree/release/src/examples/pininterrupt

74HC4051 S0	= GPIO10
74HC4051 S1	= GPIO11
Comparator Q1 = GPIO15 (interrupt input pin)

A4 - GND at input
A5 - V input
A6 - V Reference -2.500 V
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
	mux_S0        = machine.GPIO10
	mux_S1        = machine.GPIO11
	comparatorPin = machine.GPIO15
	tpA           = machine.GPIO13

	startCounter1, endCounter1 = uint32(0), uint32(0)
	counter1                   = uint32(0)
	startCounter2, endCounter2 = uint32(0), uint32(0)
	counter2                   = uint32(0)

	startPeriod, startPeriodVIN = uint32(0), uint32(0)
	endPeriod, endPeriodVIN     = uint32(0), uint32(0)
	periodVIN                   = uint32(0)
	width                       = uint32(0)
	widthnS                     = uint32(0)
	widthZero                   = uint32(0)

	voltageRead            = false
	vinPeriod, vinPeriodns = uint32(0), uint32(0)
	//startTime              = time.Now() // Start measuring time
	voltage = float32(0)
)

const MAX_RELOAD = 0xF42400 // Maximum reload value (24 bits): 16,777,216 ticks 0xFFFFFF

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

// Delay with Systick Timer - processor clock
func delayMSsystick(ticks uint32) {
	startSystick := readSysTick()  // Read current systick timer
	target := startSystick - ticks // Calculate target delay
	for readSysTick() > target {   // Wait target timeout
	}
}

func testPointA() {
	tpA.High()         // Set GPIO13
	delayMSsystick(10) // 10 sticks = 50 ns
	tpA.Low()          // Reset GPIO13
}

// Phase I - Reset Integrator - T0
// 74HC4051 - set Address 7 => Reset Integrator
func resetIntegrator() {
	mux_S0.High()                   // MUX bit1
	mux_S1.High()                   // MUX bit2
	delayMSsystick((5e6 / 5) - 125) // Delay to Capacitor discharge 5 ms = ((5e6 / 5) - 125)
}

// Phase II - Integrate V Input - T1
// 74HC4051 - set Address 5 => Read VIN
func integVIN() {
	mux_S0.High()                   // MUX bit1
	mux_S1.Low()                    // MUX bit2
	startPeriodVIN = readSysTick()  // Read SysTick before start integration
	delayMSsystick((10e6 / 5) - 30) // Read VIN 10 ms = ((10e6 / 5) - 240)
	endPeriodVIN = readSysTick()    // Read SysTick after end integration
}

// Phase III - Integrate - VREF - T2
// 74HC4051 - set Address 6 => Read -VREF
func integVREF() {
	mux_S0.Low()  // MUX bit1
	mux_S1.High() // MUX bit2
	//testPointA()              // Set GPIO13
	startPeriod = readSysTick() // Read SysTick before start integration
	delayMSsystick((15e6 / 5))  // Read VREF 6 ms = 6e6 ns ((6e6 / 5) - 18)
}

// Phase IV - GND at Input
// 74HC4051 - set Address 4
func gndInput() {
	mux_S0.Low()                     // MUX bit1
	mux_S1.Low()                     // MUX bit2
	delayMSsystick((10e6 / 5) - 230) // Read VIN 10 ms = ((10e6 / 5) - 240)
}

func zeroADC() {
	rp.PPB.SetSYST_CVR_CURRENT(0) // Clear current value (24 bits) SysTick timer - max time 83 ns
	resetIntegrator()             // Reset Integrator
	gndInput()                    // zero Integrator
	integVREF()                   // Integrate V REF
}

func dualSlopeADC() {
	rp.PPB.SetSYST_CVR_CURRENT(0) // Clear current value (24 bits) SysTick timer - max time 83 ns
	resetIntegrator()             // Reset Integrator
	integVIN()                    // Integrate V Input
	integVREF()                   // Integrate V REF
}

func adcComparatorISR(pin machine.Pin) {
	endPeriod = readSysTick() // Read SysTick after rising interrupt
	//width = startPeriod - endPeriod
	voltageRead = true
}

func main() {

	machine.InitSerial()        // Initialize serial for debug output
	time.Sleep(3 * time.Second) // Sleep to catch prints on serial
	fmt.Printf("Dual Slope ADC RP2040 V9 \n")

	// MUX Address pins configuration
	mux_S0.Configure(machine.PinConfig{Mode: machine.PinOutput})
	mux_S1.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Test Point A configuration 
	tpA.Configure(machine.PinConfig{Mode: machine.PinOutput})

	// Comparator interrupt PIN configuration
	comparatorPin.Configure(machine.PinConfig{Mode: machine.PinInputPullup}) //PinInputPullup
	comparatorPin.SetInterrupt(machine.PinRising, adcComparatorISR)

	cpuFreq := machine.CPUFrequency() // Get CPU frequency
	fmt.Println("CPU Frequency:", int(cpuFreq), "Hz")	

	nanoSECperTick := 1e9 / cpuFreq // Calculate nanoseconds per tick
	fmt.Println("Nanoseconds per tick:", nanoSECperTick)

	enableSysTick() // enable SysTick Timer

	for {

		zeroADC()

		if voltageRead {
			comparatorPin.SetInterrupt(machine.PinRising, nil) // Disable interrupt to avoid multiple triggers 

			widthZero = startPeriod - endPeriod
			//fmt.Printf("Zero Width: %d ticks \n", widthZero)

			//time.Sleep(time.Millisecond * 20)                               // delay 20 ms
			delayMSsystick(20e6 / 5)
			comparatorPin.SetInterrupt(machine.PinRising, adcComparatorISR) // enable PinRising interrupt
			voltageRead = false
		}

		dualSlopeADC()

		if voltageRead {
			comparatorPin.SetInterrupt(machine.PinRising, nil) // Disable interrupt to avoid multiple triggers 

			width = startPeriod - endPeriod
			widthnS = width * 5 // 5 ns per tick
			periodVIN = startPeriodVIN - endPeriodVIN

			// fmt.Printf("startPeriod: %d ticks / ", startPeriod)
			// fmt.Printf("endPeriod: %d ticks / ", endPeriod)
			// width = startPeriod - periodWidth

			// fmt.Printf("Period VIN: %d ticks / ", periodVIN) //expected VIN period = 2e6 ticks
			fmt.Printf("Width: %d ticks / ", width)
			// fmt.Printf("Width ns: %d ns * ", widthnS)

			// VREF / VIN period  = 2.500 / 2e6
			vinPeriod = (width - widthZero)        // VREF Period width
			vinPeriodns = vinPeriod * 5            // 5 ns per tick
			voltage = float32(vinPeriod) * 1.25e-6 // Period VIN width * V bit 2.465e-6

			fmt.Printf("VIN Period: %d ticks / ", vinPeriod)
			// fmt.Printf("VIN Period: %d ns * ", vinPeriodns)
			fmt.Printf("Voltage: %.4f V\n", voltage)
			//fmt.Printf("\n")

			time.Sleep(time.Millisecond * 100)                              // delay 100 ms
			comparatorPin.SetInterrupt(machine.PinRising, adcComparatorISR) // enable PinRising interrupt
			voltageRead = false
		}
	}
}


