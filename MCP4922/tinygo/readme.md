
Connects to an MCP4922 - 12-Bit Dual Output DAC with SPI Interface

MCP4922 Datasheet: www.microchip.com/en-us/product/mcp4922

Based on github:

.[tinygo-org/tinygo/src/examples/mcp3008/mcp3008.go](https://github.com/tinygo-org/tinygo/tree/731532cd2b6353b60b443343b51296ec0fafae09/src/examples/mcp3008)

.[helgenodland/MCP4922-Arduino-SPI-Library/](https://github.com/helgenodland/MCP4922-Arduino-SPI-Library)

.[pico-playground/audio/sine_wave/sine_wave.c](https://github.com/raspberrypi/pico-playground/tree/master/audio/sine_wave)

Connections for SPI interface Bus 0 - Raspberry Pico

	SPI0_SCK_PIN = GPIO 18 (PIN 24) =========> MCP4922 CLK  (PIN 04)
    TX SPI0_SDO_PIN = GPIO 19 (PIN 25) ======> MCP4922 SDI  (PIN 05)
    RX SPI0_SDI_PIN = GPIO 16 (PIN 21) ======> No connection
    SPIO_CS_PIN = GPIO 05 (PIN 07) ==========> MCP4922 -CS  (PIN 03)

	MCP4922 VREFA (pin 13) =====> +3.3V (only bellow 3,3V!)
	MCP4922 VREFB (pin 11) =====> +3.3V (only bellow 3,3V!)
	MCP4922 -SHDN (pin 09) =====> +3.3V
	MCP4922 -LDAC (pin 08) =====> GND
	MCP4922   VDD (pin 01) =====> +3.3V
	Rasp Pico GND ======> MCP4922 GND - pin 12 (do not forget this)
 
SPI Modes : [analog.com/en/analog-dialogue/articles/introduction-to-spi-interface.html](https://www.analog.com/en/analog-dialogue/articles/introduction-to-spi-interface.html)

