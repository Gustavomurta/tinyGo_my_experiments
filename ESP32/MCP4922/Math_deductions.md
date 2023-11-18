**Analyzing the waveforms and doing some mathematical calculations:**

The datasheet of the MCP4922, the maximum frequency of the SPI clock is 20 MHz.

Changing the SPI clock frequency to 80 MHz in the program:

```	machine.SPI.Configure(machine.SPI2, machine.SPIConfig{
		Frequency: 80000000, // clock 80 MHz MAX
		SCK:       18,
		SDO:       19,
		SDI:       23,
		Mode:      0}) // SPI mode 0,0
```
