

Running Pico Frequency Counter V2: 

```
Raspberry Pico internal Frequency Counter V2: 
Clock Reference Control is 12 MHz (XOSC) 
Reference clock for Frequency Counter: 12000 kHz
Test interval of the frequency counter: 1003.52 us
Minimum pass frequency: 0 kHz
Maximum pass frequency: 33554431 kHz
Delays the start of the frequency counting: 0 
Result of frequency measurement: 200000 kHz       for PLL SYS 200 MHz
Result of frequency measurement frac: 0 
```


**CLOCKS: FC0_SRC Register:** 
```
Clock sent to frequency counter, set to 0 when not required.
Writing to this register initiates the frequency count. 
(RP2040 Datasheet.pdf)

0x01 → PLL_SYS_CLKSRC_PRIMARY
0x02 → PLL_USB_CLKSRC_PRIMARY
0x03 → ROSC_CLKSRC
0x04 → ROSC_CLKSRC_PH
0x05 → XOSC_CLKSRC
0x06 → CLKSRC_GPIN0
0x07 → CLKSRC_GPIN1
0x08 → CLK_REF
0x09 → CLK_SYS
0x0a → CLK_PERI
0x0b → CLK_USB
0x0c → CLK_ADC
0x0d → CLK_RTC

```





