package main

import (
	"device/avr"
	"machine"
)

var (
	i2cFrequency uint32 = 100000
	dataTest            = []byte{0x57, 0x00, 0x01, 0x2, 0x03, 0x04}
)

func main() {

	avr.TWSR.ClearBits((avr.TWSR_TWPS0 | avr.TWSR_TWPS1))   // BUG correction
	avr.TWBR.Set(uint8(((machine.CPUFrequency() / i2cFrequency) - 16) / 2))
	avr.TWCR.Set(avr.TWCR_TWEN)

	// Clear TWI interrupt flag, put start condition on SDA, and enable TWI.
	avr.TWCR.Set((avr.TWCR_TWINT | avr.TWCR_TWSTA | avr.TWCR_TWEN))
	//avr.TWCR.Set(avr.TWCR_TWINT)

	for i := 0; i < 7; i++ {
		writeByte(dataTest[i])
	}
}

// writeByte writes a single byte to the I2C bus.
func writeByte(data byte) error {
	// Write data to register.
	avr.TWDR.Set(data)

	// Clear TWI interrupt flag and enable TWI.
	avr.TWCR.Set(avr.TWCR_TWEN | avr.TWCR_TWINT)

	// Wait till data is transmitted.
	for !avr.TWCR.HasBits(avr.TWCR_TWINT) {
	}
	return nil

}
