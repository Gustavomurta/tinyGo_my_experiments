
/*

Connects to an MCP3204 ADC - 4 Channel SPI 12 bits ADC

 MCP3204 Datasheet:https://ww1.microchip.com/downloads/en/DeviceDoc/21298e.pdf

https://github.com/tinygo-org/tinygo/blob/release/src/examples/mcp3008/mcp3008.go

Default pins for SPI interface Bus 0 - Raspberry Pico
ref: tinygo/src/machine/board_pico.go

	SPI0_SCK_PIN = GPIO18
	SPI0_SDO_PIN = GPIO19    Tx
	SPI0_SDI_PIN = GPIO16    Rx

Below is a list of GPIO pins corresponding to SPI0 bus on the rp2040:
ref: tinygo/src/machine/machine_rp2040_spi.go

    SI : 0, 4, 17  a.k.a RX and MISO (if rp2040 is master)
    SO : 3, 7, 19  a.k.a TX and MOSI (if rp2040 is master)
    SCK: 2, 6, 18

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
	cs   = machine.Pin(5) //  CS  GPIO5
	vRef = 3.27           // VRef 3.27 Volts
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
		Frequency: 1000000, // clock 1 MHz
		Mode:      0})      // SPI modes 0,0 or 1,1

	tx = make([]byte, 3) // 3 bytes command
	rx = make([]byte, 3) // 3 bytes received data

	lsbVolt = vRef / 4096 // LSB unit in Volts

	for {
		val, _ = Read(1) // read Channel 1
		println(" CH 0 = ", float32(val)*lsbVolt)
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
	cs.Low()
	machine.SPI0.Tx(tx, rx)
	result = uint16((rx[1]&0x1F))<<8 + uint16(rx[2]) // second and third bytes
	cs.High()

	return result, nil
}
