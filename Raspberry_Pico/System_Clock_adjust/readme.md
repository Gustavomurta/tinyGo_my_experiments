
```
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
```
