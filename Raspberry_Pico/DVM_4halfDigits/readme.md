**Digital Voltmeter with 4 ½ dígits - Raspberry Pico**

Using **Tinygo** - easy, fast and amazing language to control microcontrollers. 

https://github.com/tinygo-org/tinygo

Using an old technology, but still used in Voltmeters today - **Dual slope ADC**.
The counter/timer uses the RP2040 processor clock frequency of 200 MHz. 
Therefore, the resolution is 5 nanoseconds. This allows for greater precision and resolution for the DVM.

The circuit was assembled with the Raspberry Pico - RP2040. 
But it can be modified for the Raspberry Pico 2 including WIFI and Bluetooth - Raspberry Pico W.

To vary the range of voltages to be measured, voltage dividers with precision resistors can be implemented at input of the DVM. 
The same circuit is used in modern voltmeters.

**I'm already getting measurements from 0 to 2.5000 V**. But the project is ongoing to improve performance.
Currently the DVM measures positive voltages, but it could be implemented to measure negative voltages as well.
In the current version it is possible to take 5 measurements per second, but this can be improved.

Good reference for **Dual Slope ADC** at the end of the article. But I didn't use the **Analog Devices** circuit.
https://wiki.analog.com/university/courses/electronics/electronics-lab-adc

