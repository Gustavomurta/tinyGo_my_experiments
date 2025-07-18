
**Use vcocalc.py to configure registers:**

https://github.com/raspberrypi/pico-sdk/tree/master/src/rp2_common/hardware_clocks/scripts

```
Examples:

Requested: 125.0 MHz
Achieved:  125.0 MHz
REFDIV:    1
FBDIV:     125 (VCO = 1500.0 MHz)             x 12 MHz
PD1:       6
PD2:       2

Requested: 133.0 MHz
Achieved:  133.0 MHz
REFDIV:    1
FBDIV:     133 (VCO = 1596.0 MHz)
PD1:       6
PD2:       2

Requested: 150.0 MHz
Achieved:  150.0 MHz
REFDIV:    1
FBDIV:     125 (VCO = 1500.0 MHz)
PD1:       5
PD2:       2

Requested: 200.0 MHz
Achieved:  200.0 MHz
REFDIV:    1
FBDIV:     100 (VCO = 1200.0 MHz)
PD1:       6
PD2:       1

Requested: 240.0 MHz
Achieved:  240.0 MHz
REFDIV:    1
FBDIV:     120 (VCO = 1440.0 MHz)
PD1:       6
PD2:       1
```

```

**Running tests with Raspberry Pico - RP2040:**


Raspberry Pico Overclock V3: 
Read Voltage Regulators Registers
Voltage Regulator Register Value: 0x000010C1 
Voltage Select = 0xC -> 1.15 V

Read PLL SYS Registers
Reference Clock Divider Value: 1 
PLL SYS Feedback divisor: 100 
PLL VCO frequency: 1200 MHz 
PLL SYS post divider 1: 6 
PLL SYS post divider 2: 1 
PLL SYS - System Clock frequency: 200 MHz 

PLL SYS - Adjusting clock settings

Read PLL SYS Registers
Reference Clock Divider Value: 1 
PLL SYS Feedback divisor: 125 
PLL VCO frequency: 1500 MHz 
PLL SYS post divider 1: 6 
PLL SYS post divider 2: 2 
PLL SYS - System Clock frequency: 125 MHz 

Core temperature: 21.52℃
Core temperature: 21.05℃
Core temperature: 21.05℃

```
