
**W8782 CONFIGURATION PINS:**

In the WM8782 Datasheet:
															

![image](https://github.com/user-attachments/assets/eedbfa51-1048-4045-95e3-7a6ce711a798)



For my tests (Master mode):
															
![image](https://github.com/user-attachments/assets/e50d4a35-fb67-4c6e-b887-cd11dd324014)



![image](https://github.com/user-attachments/assets/22dd715e-e34a-4a58-9567-ec4add3da20b)




The module is powered by the VCC pin with 5V. This module has two 3.3V voltage regulators. One for the analog part (AVDD) and one for the digital part (DVDD).

Therefore the digital signals from the I2S interface are compatible with the Raspberry Pico (3.3V). They can be connected directly.

**W8782 Module - ADC I2S SCHEMATIC :** 

Using reverse engineering, I created this schematic.

![image](https://github.com/user-attachments/assets/3c1b162c-6e4d-4e85-b12e-8f6e02dfb5c0)
