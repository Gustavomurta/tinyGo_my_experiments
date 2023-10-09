Raspberry Pico - Tinygo - PWM 

Digital Signal Generator - two channels with control of the duty cycle - 0 to 100% 

Configure frequency in Hz / min: 4 Hz and max: 31.25 MHz

But, some restrictions: 

As the frequency is generated through CPU clock frequency dividers, the higher the frequency, the lower the precision.

If the chosen frequency is proportional to the CPU clock, you can obtain very good accuracy.

In the future I will include programming to change the CPU clock, for better options.

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

Each PWM Channel is equipped with the following:

• 16-bit counter

• 8.4 fractional clock divider

• Two independent output channels, duty cycle from 0% to 100% inclusive

![image](https://github.com/Gustavomurta/tinyGo_my_experiments/assets/4587366/b5ea8897-5f1b-49bd-b195-0381b2e7086a)




================================================================================

In my program, I chose PWM Channel 2:

pwm  = machine.PWM2  // PWM Channel 2

pinA = machine.GPIO4 // GPIO4 (pin 06) =>  peripherals: PWM2 channel A output

pinB = machine.GPIO5 // GPIO5 (pin 07) =>  peripherals: PWM2 channel B output
