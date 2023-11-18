### ESP32 - SPI Controller:

ESP32 integrates four SPI controllers which can be used to communicate with external devices that use the SPI protocol.

SPI0, SPI1, SPI2 (commonly referred to as HSPI), and SPI3 (commonly referred to as VSPI).

SP0 and SP1 are used internally to communicate with the built-in flash memory, and you should not use them for other tasks.

Controllers SPI2 and SPI3 can be configured as either a master or a slave.

When used as a master, each SPI controller can drive multiple CS signals (CS0~CS2) to activate multiple slaves.

Controllers SPI1~SPI3 share two DMA channels.

Many ESP32 boards come with default SPI pins pre-assigned. The pin mapping for most boards is as follows:

![image](https://github.com/Gustavomurta/tinyGo_my_experiments/assets/4587366/65c45f31-9a0c-4cdd-81a2-e025cccf2a32)


**References:**

_https://www.espressif.com/sites/default/files/documentation/esp32_technical_reference_manual_en.pdf_

_https://randomnerdtutorials.com/esp32-spi-communication-arduino/_
