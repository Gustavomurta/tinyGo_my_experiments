/*
Raspberry Pico 2 - RP2350 - rp2350 CPU Clock configuration and temperature reading
Gustavo Murta 2026_05_23
tinygo version 0.39.0 windows/amd64 (using go version go1.25.1 and LLVM version 19.1.2)
C:\Users\jgust\tinygo\programas\Raspberry_Pico2\cpu_clock_set
tinygo flash -target pico2 main.go

RP2350 Datasheet.pdf
6.3. Core voltage regulator:
voltage can be set in the range 0.55 V to 3.30 V, and the regulator can supply up to 200mA.
6.4. Power management (POWMAN) registers
8.2. Crystal oscillator (XOSC)
8.6. PLL
 There are two PLLs in RP2350. They are:
• pll_sys - used to generate up to a 150 MHz system clock
• pll_usb - used to generate a 48 MHz USB reference clock

The programming sequence for the PLL is as follows:
1. Program the reference clock divider (is a divide by 1 in the RP2350 case).
2. Program the feedback divider.
3. Turn on the main power and VCO.
4. Wait for the VCO to achieve a stable frequency, as indicated by the LOCK status flag.
5. Set up post dividers and turn them on.

Default PLL configuration:
                  REF     FBDIV VCO            POSTDIV
PLL SYS: 12 / 1 = 12MHz * 125 = 1500MHz / 6 / 1 = 150MHz
PLL USB: 12 / 1 = 12MHz * 100 = 1200MHz / 5 / 5 =  48MHz

Reference clock frequency min=5MHz, max=800MHz
Feedback divider min=16, max=320
VCO frequency min=400MHz, max=1600MHz

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
	case 0x60:
		return 0.85
	case 0x70:
		return 0.90
	case 0x80:
		return 0.95
	case 0x90:
		return 1.00
	case 0xA0:
		return 1.05
	case 0xB0:
		return 1.10
	case 0xC0:
		return 1.15
	case 0xD0:
		return 1.20
	case 0xE0:
		return 1.25
	case 0xF0:
		return 1.30
	default:
		return 0.00 // decode error
	}
}

// Read Power management Registers ///////////////////////////////////
func powmanRegisters() {

	fmt.Println("Read Power management Registers")

	vregCtrlStatus := rp.POWMAN.VREG_CTRL.Get() // Voltage regulator control and status
	fmt.Printf("Power management VREG CTRL Value: 0x%08X \n", vregCtrlStatus)

	vregSet := rp.POWMAN.VREG.Get() // Voltage Regulator settings
	fmt.Printf("Voltage Regulator settings = 0x%04X -> %.2f V\n", vregSet, decodeVsel(vregSet))

	fmt.Println()
}

// Read PLL SYS registers ////////////////////////////////////////////////
func pllSysRegisters() {

	fmt.Println("Read PLL SYS Registers")

	refClockDivider := rp.PLL_SYS.GetCS_REFDIV()
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
/*
Reference Clock Divider Value: 1
PLL SYS Feedback divisor: 125
PLL VCO frequency: 1500 MHz
PLL SYS post divider 1: 5
PLL SYS post divider 2: 2
PLL SYS - System Clock frequency: 150 MHz

Reference Clock Divider Value: 1
PLL SYS Feedback divisor: 100
PLL VCO frequency: 1200 MHz
PLL SYS post divider 1: 6
PLL SYS post divider 2: 1
PLL SYS - System Clock frequency: 200 MHz
*/

func pllAdjustClock() {

	fmt.Println("PLL SYS - Adjusting clock settings")

	rp.PLL_SYS.SetCS_REFDIV(1)   // Set the reference clock divider to 1 (12 MHz / 1 = 12 MHz)
	rp.PLL_SYS.SetFBDIV_INT(100) // Set the feedback divider to 100 (VCO = 12 MHz * 100 = 1200 MHz)
	rp.PLL_SYS.SetPWR_PD(0)      // PD: PLL power up
	rp.PLL_SYS.SetPWR_VCOPD(0)   // VCOPD: PLL VCO power up

	for rp.PLL_SYS.GetCS_LOCK() == 0 { // Wait for the PLL lock = 1
	}

	rp.PLL_SYS.SetPRIM_POSTDIV1(6) // Set the first Post divider to 6 (1200 MHz / 6 = 200 MHz)
	rp.PLL_SYS.SetPRIM_POSTDIV2(1) // Set the second Post divider to 1 (200 MHz / 1 = 200 MHz)
	rp.PLL_SYS.SetPWR_POSTDIVPD(0) // Post divider power up

	fmt.Println()
}

func main() {

	machine.InitSerial()        // Initialize serial for debug output
	time.Sleep(3 * time.Second) // Sleep to catch prints on serial
	fmt.Printf("Raspberry Pico 2 CPU Clock config: \n")

	powmanRegisters() // Read Power management Registers

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
