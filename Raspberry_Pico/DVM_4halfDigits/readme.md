## **Digital Voltmeter with 4 ½ digits - Raspberry Pico**

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

### **Dual-Slope ADC**

The Dual-Slope ADC (or a variant) is at the heart of many of the most accurate digital voltmeters. This architecture has a few useful characteristics: only a few precision components are required as most error sources cancel out, it can be configured to reject particular noise frequencies such as 50Hz or 60Hz line noise, and it is insensitive to high-frequency noise.

<img width="500" height="261" alt="image" src="https://github.com/user-attachments/assets/5a9aa2d3-17cf-4f11-afbc-765fba1d9b4d" />

**Dual-Slope ADC structure**

The converter operates by applying the unknown input voltage to an integrator for a fixed time period (called “runup”), after which a known reference voltage, of opposite polarity to the input, is applied to the integrator (called “rundown”). Thus, the input voltage can be calculated from the reference voltage and the ratio of the rundown to runup times:

Vin = Vref * (T_rundown/T_runup)

<img width="500" height="384" alt="image" src="https://github.com/user-attachments/assets/a53ad6bb-165f-4c9d-8431-6e42d636e004" />

Source: https://wiki.analog.com/university/courses/electronics/electronics-lab-adc


