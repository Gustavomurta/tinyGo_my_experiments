/*
Connects Raspberry Pico to an MCP4922 - 12-Bit Dual Output DAC with SPI Interface

MCP4922 Datasheet: www.microchip.com/en-us/product/mcp4922

Based on github:
tinygo-org/tinygo/blob/release/src/examples/mcp3008/mcp3008.go
helgenodland/MCP4922-Arduino-SPI-Library/blob/master/MCP4922.cpp
pico-playground/audio/sine_wave/sine_wave.c

Connections for SPI interface Bus 0 - Raspberry Pico

	SPI0_SCK_PIN = GPIO 18 (PIN 24) =========> MCP4922 CLK  (PIN 04)
    TX SPI0_SDO_PIN = GPIO 19 (PIN 25) ======> MCP4922 SDI  (PIN 05)
    RX SPI0_SDI_PIN = GPIO 16 (PIN 21) ======> No connection
    SPIO_CS_PIN = GPIO 05 (PIN 07) ==========> MCP4922 -CS  (PIN 03)

	MCP4922 VREFA (pin 13) =====> +3.3V (only bellow 3,3V!)
	MCP4922 VREFB (pin 11) =====> +3.3V (only bellow 3,3V!)
	MCP4922 -SHDN (pin 09) =====> +3.3V
	MCP4922 -LDAC (pin 08) =====> GND
	MCP4922   VDD (pin 01) =====> +3.3V
	Rasp Pico GND ======> MCP4922 GND - pin 12 (do not forget this)
SPI Modes : analog.com/en/analog-dialogue/articles/introduction-to-spi-interface.html
*/

package main

import (
	"machine"
	"math"
)

const (
	cs                = machine.Pin(5)     //  SPI0_CS  PICO-GPIO5
	raw_table_size    = 256                // 256 values
	DAC_config_chan_A = 0b0111000000000000 // CH A - Buffered - Gain 1X - OUT enabled
	DAC_config_chan_B = 0b1111000000000000 // CH B - Buffered - Gain 1X - OUT enabled
)

var (
	tx         []byte
	rx         []byte
	sine_float float32
	multiplier uint16
	channelA   uint16
	channelB   uint16
)

func main() {
	cs.Configure(machine.PinConfig{Mode: machine.PinOutput})

	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 4000000, // clock 4 MHz
		Mode:      0})      // SPI mode 0,0

	tx = make([]byte, 2) // 2 bytes write data
	rx = make([]byte, 2) // 2 bytes received data

	// Sine table
	raw_table := make([]uint16, raw_table_size)

	// Table of values to be sent to DAC
	DAC_data_A := make([]uint16, raw_table_size)
	DAC_data_B := make([]uint16, raw_table_size)

	// Build wave table and DAC data tables
	for i := 0; i < (raw_table_size); i++ {

		/* Ramp wave
		multiplier = 10
		raw_table[i] = uint16(i) // ramp wave table*/

		// Sine wave
		multiplier = 1
		raw_table[i] = uint16(2047 + 2048*math.Cos(float64(float32(i)*(2*math.Pi/raw_table_size)))) // sine wave table */

		/* triangle wave
		multiplier = 20
		if i < 128 {
			raw_table[i] = uint16(i)
		} else {
			raw_table[i] = 256 - uint16(i)
		}*/

		/* square wave
		multiplier = 10
		if i < 128 {
			raw_table[i] = 256 // High level
		} else {
			raw_table[i] = 0 // Low level
		}*/

		DAC_data_A[i] = uint16(DAC_config_chan_A | multiplier*raw_table[i]) // Channel A  x multiplier
		DAC_data_B[i] = uint16(DAC_config_chan_B | multiplier*raw_table[i]) // Channel B  x multiplier
	}

	for {
		for i := 0; i < (raw_table_size); i++ {
			Write(DAC_data_A[i]) // Channel A
			// Write(DAC_data_B[i])   // Channel B
		}
	}
}

func Write(spi_data uint16) {

	tx[0] = byte(spi_data & 0xFF00 >> 8) // first byte
	tx[1] = byte(spi_data & 0xFF)        // second byte

	cs.Low()
	machine.SPI0.Tx(tx, rx) // transmit spi data
	cs.High()
}
