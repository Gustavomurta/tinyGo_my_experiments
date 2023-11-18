Tinygo tests with MCP4922 using ESP32 - waveform generator (256 points)

max frequency:  Hz 

max SPI clock:  MHz

MCP4922 - 12-Bit Dual Voltage Output Digital-to-Analog Converter with SPI Interface

MCP4922 Datasheet: https://www.microchip.com/en-us/product/mcp4922

```
Connections for SPI interface Bus 2 - ESP32

SPI_SCK_PIN = GPIO 18 =========> MCP4922 CLK  (PIN 04)
SPI_SDO_PIN = GPIO 19 =========> MCP4922 SDI  (PIN 05)
SPI_SDI_PIN = GPIO 23 =========> No connection
SPI_CS_PIN = GPIO 05  ==========> MCP4922 -CS  (PIN 03)

MCP4922 VREFA (pin 13) =====> +3.3V (only bellow 3,3V!)
MCP4922 VREFB (pin 11) =====> +3.3V (only bellow 3,3V!)
MCP4922 -SHDN (pin 09) =====> +3.3V
MCP4922 -LDAC (pin 08) =====> GND
MCP4922   VDD (pin 01) =====> +3.3V
ESP32 GND ======> MCP4922 GND - pin 12 (do not forget this)
```


**Compiling and writing to ESP32:**

PS C:\Users\jgust\tinygo\programas\esp32\spi_examples\mcp4922> **tinygo flash -target=esp32 main.go**
```
esptool.py v4.6.2
Serial port COM6
Connecting....
Chip is ESP32-D0WDQ6 (revision v1.0)
Features: WiFi, BT, Dual Core, Coding Scheme None
Crystal is 40MHz
MAC: 30:ae:a4:XX:XX:XX
Uploading stub...
Running stub...
Stub running...
Configuring flash size...
Flash will be erased from 0x00001000 to 0x00001fff...
Warning: Image file at 0x1000 is protected with a hash checksum, so not changing the flash mode setting. Use the --flash_mode=keep option instead of --flash_mode=dout in order to remove this warning, or use the --dont-append-digest option for the elf2image command in order to generate an image file without a hash checksum
Compressed 3728 bytes to 2785...
Wrote 3728 bytes (2785 compressed) at 0x00001000 in 0.3 seconds (effective 97.3 kbit/s)...
Hash of data verified.

Leaving...
Hard resetting via RTS pin...
```
