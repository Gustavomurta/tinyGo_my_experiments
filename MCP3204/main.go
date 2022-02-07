
/*
Raspberry PICO - TinyGo experiment
Connects to an MCP3204 ADC - 4 Channel SPI 12 bits ADC

MCP3204 Datasheet:https://ww1.microchip.com/downloads/en/DeviceDoc/21298e.pdf

Based on: github.com/tinygo-org/tinygo/blob/release/src/examples/mcp3008/mcp3008.go

Connections for SPI interface Bus 0 - Raspberry Pico
ref: tinygo/src/machine/board_pico.go

	SPI0_SCK_PIN = GPIO 18 (PIN 24) ======> MCP3204 CLK  (PIN 11)
     TX SPI0_SDO_PIN = GPIO 19 (PIN 25) ======> MCP3204 DIN  (PIN 09)  
     RX SPI0_SDI_PIN = GPIO 16 (PIN 21) ======> MCP3204 DOUT (PIN 10)
         SPIO_CS_PIN = GPIO 05 (PIN 03) ======> MCP3204 -CS  (PIN 08)
	 
	 MCP3204 VREF (13) =====> +3.3V (only voltages bellow 3.3V at ADC inputs!) 

	SPI Modes : https://www.analog.com/en/analog-dialogue/articles/introduction-to-spi-interface.html
*/

package main

import (
	"errors"
	"machine"
	"time"
)

// cs is the pin used for Chip Select (CS). Change to whatever is in use on your board.
const (
	cs   = machine.Pin(5) // Chip Select  GPIO5
	vRef = 3.27           // VRef 3.27 Volts- reading with voltmeter
)

var (
	tx          []byte
	rx          []byte
	lsbVolt     float32
	val, result uint16
)

func main() {
	cs.Configure(machine.PinConfig{Mode: machine.PinOutput})

	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 1000000, // SPI clock 1 MHz
		Mode:      0})      // SPI modes 0,0 or 1,1

	tx = make([]byte, 3) // 3 bytes command
	rx = make([]byte, 3) // 3 bytes received data

	lsbVolt = vRef / 4096 // LSB unit in Volts

	for {
		val, _ = Read(1) // read Channel 1 for example (options: 0 to 3) 
		println(" CH 1 = ", float32(val)*lsbVolt)
		time.Sleep(100 * time.Millisecond) // timeout 100 ms
	}
}

// Read analog data from channel 0 to 3
func Read(channel int) (uint16, error) {
	if channel < 0 || channel > 3 {
		return 0, errors.New("Invalid channel for read")
	}

	// send command to ADC
	tx[0] = 0x06                 // start bit and single(1) or diff(0) bit + D2 channel bit
	tx[1] = byte(8+channel) << 6 // D1 and D0 channel bits
	tx[2] = 0x00                 // byte not used

	// receive ADC Data - first byte not used
	cs.Low()                                         // select ADC
	machine.SPI0.Tx(tx, rx)
	result = uint16((rx[1]&0x1F))<<8 + uint16(rx[2]) // second and third bytes
	cs.High()                                        // deselect ADC

	return result, nil
}
