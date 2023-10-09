/*
Digital Signal Generator - Two ports with control of the duty cycle
Raspberry Pico - Tinygo - PWM
Gustavo Murta 2023/10/08

References:
https://github.com/tinygo-org/tinygo/blob/release/src/examples/pwm/pwm.go
https://github.com/tinygo-org/tinygo/blob/release/src/examples/pwm/pico.go
https://github.com/tinygo-org/tinygo/blob/release/src/machine/machine_rp2040.go
https://github.com/tinygo-org/tinygo/blob/release/src/machine/machine_rp2040_pwm.go
*/

package main

import (
	"fmt"
	"machine"
	"time"
)

var (
	pwm                              = machine.PWM2  // PWM Channel 2
	pinA                             = machine.GPIO4 // GPIO4 (pin 06) =>  peripherals: PWM2 channel A
	pinB                             = machine.GPIO5 // GPIO5 (pin 07) =>  peripherals: PWM2 channel B
	pwmPeriod                        uint64
	frequency                        uint32
	dutyCycleA, dutyCycleB           uint
	divider                          uint32
	ratioDutyCycleA, ratioDutyCycleB uint32
)

func main() {

	frequency := 1250000 // frequency in Hz / min: 4 Hz and max: 31.25 MHz

	pwmPeriod := uint64(1e9 / frequency) // period in nanoseconds  min:32 ns / max:250 ms (250,000,000 ns)

	err := pwm.Configure(machine.PWMConfig{Period: pwmPeriod}) // configure PWM period

	if err != nil {
		println("failed to configure PWM")
		return
	}

	// Duty Cycle is a percentage of the ratio of pulse duration
	dutyCycleA := 27 //  Duty Cycle Channel A => 0 to 100 %
	dutyCycleB := 63 //  Duty Cycle Channel B => 0 to 100 %

	divider := pwm.Top() + 1                              // complete period
	ratioDutyCycleA := divider * uint32(dutyCycleA) / 100 // period / duty cycle => Channel A
	ratioDutyCycleB := divider * uint32(dutyCycleB) / 100 // period / duty cycle => Channel B

	time.Sleep(time.Second * 5) // Delay a bit on startup to easily catch the first messages

	cpuFreq := machine.CPUFrequency() // Microcontroller clock frequency
	cpuPeriod := 1e9 / cpuFreq        // clock period
	fmt.Println("CPU Frequency:", cpuFreq, "Hz", "  CPU Period:", cpuPeriod, "ns")

	fmt.Printf("Frequency: %d Hz\n", frequency)
	fmt.Printf("PMW Period: %d ns\n", pwmPeriod)

	// Configure the two channels as outputs
	channelA, err := pwm.Channel(pinA)
	if err != nil {
		println("failed to configure channel A")
		return
	}
	channelB, err := pwm.Channel(pinB)
	if err != nil {
		println("failed to configure channel B")
		return
	}

	// Invert Channel B to demonstrate output polarity
	// pwm.SetInverting(channelB, true)

	pwm.Set(channelA, ratioDutyCycleA) // Channel A running at duty cycle 0 to 100 %
	pwm.Set(channelB, ratioDutyCycleB) // Channel B running at duty cycle 0 to 100 %

	for {
		// loop
	}
}
