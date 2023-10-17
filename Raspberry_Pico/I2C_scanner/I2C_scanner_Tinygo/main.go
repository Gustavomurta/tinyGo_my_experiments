/*
I2C Scanner for TinyGo - Raspberry Pico
Tool to verify the correct functioning of the I2C interface and identify the addresses of I2C devices.

Inspired by: https://playground.arduino.cc/Main/I2cScanner/
Based on: https://github.com/ysoldak/tinygo-stuff/blob/master/examples/common/i2c-scan/main.go

Gustavo Murta 2023/10/15

Raspberry Pico GPIO BUS uses 3.3V level. It is not recommended to connect directly 5V I2C Devices.
If becessary, use Level converters.
For the I2C bus to function properly, pullup resistors (3.3K ohms) are necessary on the SCL and SDA lines.
Check if the i2C device already has these pullup resistors.

I2C Default pins on Raspberry Pico:

I2C0_SDA_PIN = GP4
I2C0_SCL_PIN = GP5

I2C1_SDA_PIN = GP2
I2C1_SCL_PIN = GP3
*/

package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {

	machine.I2C0.Configure(machine.I2CConfig{
		Frequency: 100 * machine.KHz, // I2C clock frequency 100 KHz
	})

	w := []byte{}
	r := []byte{0} // shall pass at least one byte for I2C code to at all try to communicate

	time.Sleep(3 * time.Second) // Delay to enable USB monitor time to attach

	fmt.Println("Scanning I2C devices...")

	for {
		i2cDevices := 0
		for address := 1; address < 127; address++ {
			err := machine.I2C0.Tx(uint16(address), w, r) // try read a byte from the current address

			if err == nil {
				fmt.Printf("I2C device found at address %#X (%d) \n", address, address)
				i2cDevices++
			}
		}
		if i2cDevices == 0 {
			fmt.Println("No I2C devices found")
		}
		time.Sleep(2 * time.Second)
	}
}
