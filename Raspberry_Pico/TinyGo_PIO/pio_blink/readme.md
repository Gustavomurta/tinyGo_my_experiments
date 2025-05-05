Raspberry Pico - testing frequency generation using PIO with TinyGo. 

References:

https://github.com/tinygo-org/pio/tree/main/rp2-pio/examples/blinky  (TinyGo version) 

https://github.com/raspberrypi/pico-examples/tree/master/pio/pio_blink   (C++ Version) 

generate blink_pio.go : pioasm -o go blink.pio blink_pio.go

Compile main.go program : tinygo build -target pico

Compile and flash main.go program : tinygo flash -target pico
