// go:generate pioasm -o go blink.pio blink_pio.go
// https://github.com/tinygo-org/pio/tree/main/rp2-pio/examples/blinky
// tinygo flash -target pico

package main

import (
	"machine"
	"time"

	pio "github.com/tinygo-org/pio/rp2-pio"
)

var clockFreq uint

func main() {
	time.Sleep(3 * time.Second) // Sleep to catch prints.
	Pio := pio.PIO0

	clockFreq := uint(machine.CPUFrequency())
	println("CPU Frequency:", clockFreq, "Hz")

	offset, err := Pio.AddProgram(blinkInstructions, blinkOrigin)
	if err != nil {
		panic(err.Error())
	}
	println("Loaded program at", offset)

	blinkPinForever(Pio.StateMachine(0), offset, machine.LED, 3, clockFreq)      // Led on board - 3 Hz
	blinkPinForever(Pio.StateMachine(1), offset, machine.GPIO6, 4, clockFreq)    // GPIO6 - 4 Hz
	blinkPinForever(Pio.StateMachine(2), offset, machine.GPIO11, 1, clockFreq)   // GPIO11 - 1 Hz
}

func blinkPinForever(sm pio.StateMachine, offset uint8, pin machine.Pin, freq uint, clockFreq uint) {
	blinkProgramInit(sm, offset, pin)
	sm.SetEnabled(true)
	println("Blinking", int(pin), "at", freq, "Hz")
	sm.TxPut(uint32(clockFreq/(2*freq)) - 3) // for frequency accuracy - PIO counter program takes 3 more cycles in total
}
