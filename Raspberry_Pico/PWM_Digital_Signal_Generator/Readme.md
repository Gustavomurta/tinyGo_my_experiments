Raspberry Pico - Tinygo - PWM 

Digital Signal Generator - with control of the duty cycle - 0 to 100% 

Gustavo Murta 2023/10/08

References:

https://github.com/tinygo-org/tinygo/blob/release/src/examples/pwm/pwm.go

https://github.com/tinygo-org/tinygo/blob/release/src/examples/pwm/pico.go

https://github.com/tinygo-org/tinygo/blob/release/src/machine/machine_rp2040.go

https://github.com/tinygo-org/tinygo/blob/release/src/machine/machine_rp2040_pwm.go

RP2040 Datasheet - A microcontroller by Raspberry Pi:

https://datasheets.raspberrypi.com/rp2040/rp2040-datasheet.pdf

==============================================================================

Chapter 4.5. PWM 

The RP2040 PWM block has 8 identical PWM Channels. 

Each channel can drive two PWM output signals, or measure the frequency or duty cycle of an input signal.

This gives a total of up to 16 controllable PWM outputs. All 30 GPIO pins can be driven by the PWM block.

Each PWM slice is equipped with the following:

• 16-bit counter

• 8.4 fractional clock divider

• Two independent output channels, duty cycle from 0% to 100% inclusive


================================================================================

pwm  = machine.PWM2  // PWM Channel 2

pinA = machine.GPIO4 // GPIO4 (pin 06) =>  peripherals: PWM2 channel A

pinB = machine.GPIO5 // GPIO5 (pin 07) =>  peripherals: PWM2 channel B
