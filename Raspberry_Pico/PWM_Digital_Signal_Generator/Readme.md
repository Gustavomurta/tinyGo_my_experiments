Raspberry Pico - Tinygo - PWM 

Digital Signal Generator - with control of the duty cycle - 0 to 100% 

Gustavo Murta 2023/10/08

RP2040 Datasheet - A microcontroller by Raspberry Pi:

https://datasheets.raspberrypi.com/rp2040/rp2040-datasheet.pdf

References:

https://github.com/tinygo-org/tinygo/blob/release/src/examples/pwm/pwm.go

https://github.com/tinygo-org/tinygo/blob/release/src/examples/pwm/pico.go

https://github.com/tinygo-org/tinygo/blob/release/src/machine/machine_rp2040.go

https://github.com/tinygo-org/tinygo/blob/release/src/machine/machine_rp2040_pwm.go


pwm  = machine.PWM2  // PWM Channel 2

pinA = machine.GPIO4 // GPIO4 (pin 06) =>  peripherals: PWM2 channel A

pinB = machine.GPIO5 // GPIO5 (pin 07) =>  peripherals: PWM2 channel B
