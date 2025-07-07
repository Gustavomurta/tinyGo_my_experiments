
**W8782 CONFIGURATION PINS:**

PIN 07 - WORD LENGTH:  0=> 16 BITS   /  Z=> 24 BITS

PIN08 - FSAMPEN:       0=> 48K       /  1=> 96K       /  Z=> 192K

PIN09 - FORMAT : 0=> RIGHT JUSTIFIED / 1=> LEFT JUSTIFIED  / Z=> I2S 

PIN20 - MODE SEL: 0=> SLAVE  /  1=> MASTER


The module is powered by the VCC pin with 5V. This module has two 3.3V voltage regulators. One for the analog part (AVDD) and one for the digital part (DVDD).

Therefore the digital signals from the I2S interface are compatible with the Raspberry Pico (3.3V). They can be connected directly.
