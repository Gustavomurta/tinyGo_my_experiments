### Tinygo - ESP32 information:

https://tinygo.org/docs/reference/microcontrollers/esp32-coreboard-v2/

#### Documentation for the machine package for the ESP32:

https://tinygo.org/docs/reference/microcontrollers/machine/esp32-coreboard-v2/

#### Do you want to learn Tinygo and Golang?

https://github.com/Gustavomurta/tinyGo_my_experiments/tree/main

#### ESP32 - Tinygo my Experiments

- MCP4922 - 12-Bit Dual Voltage Output Digital-to-Analog Converter with SPI Interface

#### Compiling TinyGo with ESP32: 

**tinygo build -target=esp32-coreboard-v2 main.go**  (to compile) 

**tinygo flash -target=esp32-coreboard-v2 main.go**  (to compile and write to ESP32)

I copied the esptool program from this link:

https://github.com/espressif/esptool/releases/

To be able to use it on Windows, I renamed the program to **esptool.py.exe**

I found where the original tool was in Windows:

C:\Users\jgust\AppData\Local\Programs\Python\Python312\Scripts

How to find where a Program is installed in Windows 11/10:

https://www.thewindowsclub.com/how-to-find-where-a-program-is-installed-in-windows-10

I copied the renamed tool into this folder.
