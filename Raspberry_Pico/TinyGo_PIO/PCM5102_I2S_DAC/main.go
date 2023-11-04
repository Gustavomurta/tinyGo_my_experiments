/*
Raspberry Pico - PCM5102 I2S DAC - Tinygo PIO
Gustavo Murta - 2023_11_03

Testing of: https://github.com/tinygo-org/pio/tree/main
https://github.com/tinygo-org/pio/tree/main/rp2-pio/piolib

I2S SCK:  Rasp Pico GND (pin 38)
I2S BCK:  Rasp Pico GPIO26 (pin 31)
I2S LRCK: Rasp Pico GPIO27 (pin 32)
I2S DIN:  Rasp Pico GPIO28 (pin 34)
*/

package main

import (
	"machine"

	pio "github.com/tinygo-org/pio/rp2-pio"
	"github.com/tinygo-org/pio/rp2-pio/piolib"
)

const (
	i2sBCK  = machine.GPIO26 // I2S BCK - Raspberry Pico GPIO26 (pin31)
	i2sData = machine.GPIO28 // I2S DIN - Raspberry Pico GPIO28 (pin34)

	// frequency    = 1000           // square wave 1KHz stereo
	// amplitude    = 10000          // amplitude of square wave  1V pp
	dacTableSize = 256 // 256 values
)

var multiplier uint32
var audioData uint32

func main() {

	dacData := make([]uint32, dacTableSize) // Table of values to be sent to DAC

	// Build DAC data table
	for i := 0; i < (dacTableSize); i++ {

		// square wave
		multiplier = 100
		if i < 128 {
			dacData[i] = multiplier * 256 // High level
		} else {
			dacData[i] = multiplier * 0 // Low level
		}
	}

	sm, _ := pio.PIO0.ClaimStateMachine()
	i2sDAC, err := piolib.NewI2S(sm, i2sData, i2sBCK)
	if err != nil {
		panic(err.Error())
	}
	i2sDAC.SetSampleFrequency(384000) // I2S BCK PCM5102 384KHz - 32 bits per sample

	for {
		for i := 0; i < (dacTableSize); i++ {
			i2sDAC.WriteStereo(dacData) // send wave data to DAC
		}
	}
}
