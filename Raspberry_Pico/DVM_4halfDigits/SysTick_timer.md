**Some theory about Systick Timer (RP2040):**

reference: https://datasheets.raspberrypi.com/rp20 ... asheet.pdf

Rp2040 Processor features :
• An ARMv6-M compliant 24-bit SysTick timer.

A 24-bit SysTick system timer, extends the functionality of both the processor and the NVIC and provides:
• A 24-bit system timer (SysTick).
• Additional configurable priority SysTick interrupt.

The SysTick timer uses a 1μs pulse as a clock enable. This is generated in the watchdog block as timer_tick. Accuracy
of SysTick timing depends upon accuracy of this timer_tick. The SysTick timer can also run from the system clock (see
SYST_CALIB).

**In my project I used the processor clock - 200 MHz.**

**2 ^ 24 bits = 16,777,216 counts
Systick timer frequency = 200 MHz (period = 5 ns)
Maximum timer count = 16,777,216 * 5 ns =~ 83 miliseconds**

**Therefore the timer is reset after each measurement process.**
• func autoZeroADC() // read Zero Volts value
• func dualSlopeADC() // read VIN voltage


**SyStick timer**

**2.4.4.1. System control register summary:**

**SYST_CSR = SysTick Control and Status Register**
Use the SysTick Control and Status Register to enable the SysTick features
Selects the SysTick timer clock source:
0 = External reference clock.
1 = Processor clock.

**SYST_RVR = SysTick Reload Value Register**
Use the SysTick Reload Value Register to specify the start value to load into the current value register when the
counter reaches 0. It can be any value between 0 and 0x00FFFFFF. A start value of 0 is possible, but has no effect
because the SysTick interrupt and COUNTFLAG are activated when counting from 1 to 0. The reset value of this
register is UNKNOWN.
To generate a multi-shot timer with a period of N processor clock cycles, use a RELOAD value of N-1. For example,
if the SysTick interrupt is required every 100 clock pulses, set RELOAD to 99

**SYST_CVR = SysTick Current Value Register**
Use the SysTick Current Value Register to find the current value in the register. The reset value of this register is
UNKNOWN.

**SYST_CALIB = SysTick Calibration value Register**
Use the SysTick Calibration Value Register to enable software to scale to any required speed using divide and
multiply.
