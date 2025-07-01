/*
Raspberry Pico - RP2040 - Pico Voltage Regulators 2
Gustavo Murta 2025/06/30
tinygo version 0.38.0 windows/amd64 (using go version go1.24.2 and LLVM version 19.1.2)
tinygo flash -target pico main.go
C:\Users\jgust\tinygo\programas\Raspberry_Pico\cpu_clock_set\
Using cmsis-svd-data/data/RaspberryPi/rp2040.svd
RP2040 datasheet.pdf - Chapter 2.10. Core Supply Regulator
*/

package main

import (
	"device/rp"
	"fmt"
	"machine"
	"time"
)

func main() {

	machine.InitSerial()        // Initialize serial for debug output
	time.Sleep(3 * time.Second) // Sleep to catch prints on serial
	println("Get PICO Voltage Regulators registers:")
	println()

	// --- Read VREG Register ---///////////////////////////////////
	vselValue := rp.VREG_AND_CHIP_RESET.VREG.Get() // Voltage regulator control and status
	fmt.Printf("VREG (Voltage Regulator) Register Value: 0x%08X \n", vselValue)

	// --- Read BOD Register ---////////////////////////////////////
	bodValue := rp.VREG_AND_CHIP_RESET.BOD.Get() // brown-out detection control
	fmt.Printf("BOD (brown-out detection) Register Value: 0x%08X \n", bodValue)

	// --- Read CHIP_RESET Register ---/////////////////////////////
	chipResetValue := rp.VREG_AND_CHIP_RESET.CHIP_RESET.Get() // Chip reset control and status
	fmt.Printf("CHIP_RESET (Chip Reset Status) Register Value: 0x%08X \n", chipResetValue)

	for {
		// Keep the program running indefinitely
	}
}
