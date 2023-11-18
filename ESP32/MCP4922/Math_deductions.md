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

My cheap logic analyzer cannot capture frequencies above 8 MHz.

![image](https://github.com/Gustavomurta/tinyGo_my_experiments/assets/4587366/694633a3-a6f6-4a00-bdea-3c544c24d5c4)

And the maximum frequency response on my Oscilloscope is 60 Mhz.

![image](https://github.com/Gustavomurta/tinyGo_my_experiments/assets/4587366/1581a5c8-ce6d-44a5-a5ba-e7e00fe2159e)




