/*
Raspberry Pico - RP2040 - overclock V3
Gustavo Murta 2025/07/17
tinygo version 0.38.0 windows/amd64 (using go version go1.24.2 and LLVM version 19.1.2)
C:\Users\jgust\tinygo\programas\Raspberry_Pico\cpu_clock_set
tinygo flash -target pico main.go

RP2040 Datasheet.pdf
2.10. Core Supply Regulator
2.15. Clocks
2.18. PLL

 * There are two PLLs in RP2040. They are:
 *   - pll_sys - Used to generate up to a 133MHz system clock
 *   - pll_usb - Used to generate a 48MHz USB reference clock

 The programming sequence for the PLL is as follows:
• Program the reference clock divider (is a divide by 1 in the RP2040 case)
• Program the feedback divider
• Turn on the main power and VCO
• Wait for the VCO to lock (i.e. keep its output frequency stable)
• Set up post dividers and turn them on

Default PLL configuration:
                  REF     FBDIV VCO            POSTDIV
PLL SYS: 12 / 1 = 12MHz * 125 = 1500MHz / 6 / 2 = 125MHz
PLL USB: 12 / 1 = 12MHz * 100 = 1200MHz / 5 / 5 =  48MHz

Reference clock frequency min=5MHz, max=800MHz
Feedback divider min=16, max=320
VCO frequency min=750MHz, max=1600MHz

https://github.com/tinygo-org/tinygo/blob/release/src/machine/machine_rp2_pll.go
*/

package main

import (
	"device/rp"
	"fmt"
	"machine"
	"time"
)

type celsius float32

func (c celsius) String() string {
	return fmt.Sprintf("%2.2f℃", c)
}

func decodeVsel(vregVsel uint32) float32 {

	switch vregVsel {
	case 0x6:
		return 0.85
	case 0x7:
		return 0.90
	case 0x8:
		return 0.95
	case 0x9:
		return 1.00
	case 0xA:
		return 1.05
	case 0xB:
		return 1.10
	case 0xC:
		return 1.15
	case 0xD:
		return 1.20
	case 0xE:
		return 1.25
	case 0xF:
		return 1.30
	default:
		return 0.80
	}
}

// Read Voltage Regulators Registers ///////////////////////////////////
func vregRegisters() {

	fmt.Println("Read Voltage Regulators Registers")

	vregCtrlStatus := rp.VREG_AND_CHIP_RESET.VREG.Get() // Voltage regulator control and status
	fmt.Printf("Voltage Regulator Register Value: 0x%08X \n", vregCtrlStatus)

	vregVsel := (vregCtrlStatus >> 4) & 0xF // Voltage select bits
	fmt.Printf("Voltage Select = 0x%X -> %.2f V\n", vregVsel, decodeVsel(vregVsel))

	fmt.Println()
}

// Read PLL SYS registers ////////////////////////////////////////////////
func pllSysRegisters() {

	fmt.Println("Read PLL SYS Registers")

	refClockDivider := rp.PLL_SYS.GetCS_REFDIV()
	//refClockDivider := (pllSysControlStatus >> rp.PLL_CS_REFDIV_Pos) & rp.PLL_CS_REFDIV_Msk // Get the reference clock divider
	fmt.Printf("Reference Clock Divider Value: %d \n", refClockDivider)

	pllSysFdbkDivisor := rp.PLL_SYS.FBDIV_INT.Get() // For XOSC=12 MHz Feedback divisor min=63, max=133
	fmt.Printf("PLL SYS Feedback divisor: %d \n", pllSysFdbkDivisor)
	fmt.Printf("PLL VCO frequency: %d MHz \n", (pllSysFdbkDivisor * 12)) // Divisor * 12 MHz

	postDiv1 := rp.PLL_SYS.GetPRIM_POSTDIV1() // Get the first Post divider
	fmt.Printf("PLL SYS post divider 1: %d \n", postDiv1)

	postDiv2 := rp.PLL_SYS.GetPRIM_POSTDIV2() // Get the second Post divider
	fmt.Printf("PLL SYS post divider 2: %d \n", postDiv2)

	// PLL SYS = (XOSC * FedbkDivisor) / postDiv1 / postDiv2
	pllSysFrequency := (12 * pllSysFdbkDivisor) / postDiv1 / postDiv2
	fmt.Printf("PLL SYS - System Clock frequency: %d MHz \n", pllSysFrequency)

	fmt.Println()

}

// Adjust the PLL SYS clock settings //////////////////////////////////////
func pllAdjustClock() {

	fmt.Println("PLL SYS - Adjusting clock settings")

	rp.PLL_SYS.SetCS_REFDIV(1)   // Set the reference clock divider to 1 (12 MHz / 1 = 12 MHz)
	rp.PLL_SYS.SetFBDIV_INT(125) // Set the feedback divider to 125 (VCO = 12 MHz * 125 = 1500 MHz)
	rp.PLL_SYS.SetPWR_PD(0)      // PD: PLL power up
	rp.PLL_SYS.SetPWR_VCOPD(0)   // VCOPD: PLL VCO power up

	for rp.PLL_SYS.GetCS_LOCK() == 0 { // Wait for the PLL lock = 1
	}

	rp.PLL_SYS.SetPRIM_POSTDIV1(6) // Set the first Post divider to 6 (1500 MHz / 6 = 250 MHz)
	rp.PLL_SYS.SetPRIM_POSTDIV2(2) // Set the second Post divider to 2 (250 MHz / 2 = 125 MHz)
	rp.PLL_SYS.SetPWR_POSTDIVPD(0) // Post divider power up

	fmt.Println()
}

func main() {

	machine.InitSerial()        // Initialize serial for debug output
	time.Sleep(3 * time.Second) // Sleep to catch prints on serial
	fmt.Printf("Raspberry Pico Overclock V3: \n")

	vregRegisters() // Read Voltage Regulators Registers

	pllSysRegisters() // Read PLL SYS registers

	pllAdjustClock() // Adjust the PLL SYS clock settings

	pllSysRegisters() // Read PLL SYS registers

	machine.ReadTemperature() // Read the temperature sensor to initialize it
	time.Sleep(1 * time.Second)

	for {
		temp := celsius(float32(machine.ReadTemperature()) / 1000)
		println("Core temperature:", temp.String())
		time.Sleep(2 * time.Second)
	}
}
