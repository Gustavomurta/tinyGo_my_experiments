/*
Raspberry Pico - RP2040 - internal Frequency Counter V2
Gustavo Murta 2025/07/17
tinygo version 0.38.0 windows/amd64 (using go version go1.24.2 and LLVM version 19.1.2)
tinygo flash -target pico main.go
C:\Users\jgust\tinygo\programas\Raspberry_Pico\frequency_counter

https://github.com/raspberrypi/pico-sdk/blob/master/src/rp2_common/hardware_clocks/clocks.c#L147
https://github.com/raspberrypi/pico-sdk/blob/master/src/rp2040/hardware_structs/include/hardware/structs/clocks.h

2.15.6.2. Using the frequency counter
To use the frequency counter, the programmer must:
• Set the reference frequency: clk_ref
• Set the mux position of the source they want to measure. See FC0_SRC
• Wait for the DONE status bit in FC0_STATUS to be set
• Read the result
*/

package main

import (
	"device/rp"
	"fmt"
	"machine"
	"time"
)

func freq_count_kHz() {

	fmt.Printf("Clock Reference Control is 12 MHz (XOSC) \n")

	for rp.CLOCKS.GetFC0_STATUS_RUNNING() == 1 {
	} // If frequency counter is running need to wait for it

	rp.CLOCKS.SetFC0_REF_KHZ(12000) // Reference Clock = 12000 kHz  Max = 1048575
	freqCounterRefKHz := rp.CLOCKS.GetFC0_REF_KHZ()
	fmt.Printf("Reference clock for Frequency Counter: %d kHz\n", freqCounterRefKHz)

	rp.CLOCKS.SetFC0_INTERVAL(10) // Set Test interval (2** interval) * 0.98us / Max value: 15
	freqCounterInterval := 1 << rp.CLOCKS.GetFC0_INTERVAL()
	fmt.Printf("Test interval of the frequency counter: %.2f us\n", (float32(freqCounterInterval) * 0.98)) // Print test interval in microseconds

	rp.CLOCKS.SetFC0_MIN_KHZ(0) // Set minimum frequency = 0 KHz
	freqCounterMinKHz := rp.CLOCKS.GetFC0_MIN_KHZ()
	fmt.Printf("Minimum pass frequency: %d kHz\n", freqCounterMinKHz) // This is optional. Set to 0 if you are not using

	rp.CLOCKS.SetFC0_MAX_KHZ(0x1ffffff) // Set maximum frequency = 33,554,431 KHz
	freqCounterMaxKHz := rp.CLOCKS.GetFC0_MAX_KHZ()
	fmt.Printf("Maximum pass frequency: %d kHz\n", freqCounterMaxKHz) // This is optional. Set to 0x1ffffff if you are not using

	rp.CLOCKS.SetFC0_DELAY(0) // Max Value: 7
	freqCounterDelay := rp.CLOCKS.GetFC0_DELAY()
	fmt.Printf("Delays the start of the frequency counting: %d \n", freqCounterDelay) // Delay is measured in multiples of the reference clock period

	rp.CLOCKS.SetFC0_SRC(0x0B) // 0x01 → PLL_SYS_CLKSRC_PRIMARY / 0x02 → PLL_USB_CLKSRC_PRIMARY / 0x05 → XOSC_CLKSRC
	// 0x08 → CLK_REF / 0x09 → CLK_SYS / 0x0b → CLK_USB /

	for rp.CLOCKS.GetFC0_STATUS_DONE() == 0 {
		// Wait for Test complete of the frequency counter
	}

	freqCounterResult := rp.CLOCKS.GetFC0_RESULT_KHZ()                         // Get the frequency counter result in kHz
	freqCounterResultFrac := rp.CLOCKS.GetFC0_RESULT_FRAC()                    // Get the fractional part of the frequency counter result
	fmt.Printf("Result of frequency measurement: %d kHz\n", freqCounterResult) // only valid when status DONE=1
	fmt.Printf("Result of frequency measurement frac: %d \n", freqCounterResultFrac)
}

func main() {

	machine.InitSerial()        // Initialize serial for debug output
	time.Sleep(3 * time.Second) // Sleep to catch prints on serial
	fmt.Printf("Raspberry Pico internal Frequency Counter V2: \n")

	freq_count_kHz() // Frequency Counter in kHz
}
