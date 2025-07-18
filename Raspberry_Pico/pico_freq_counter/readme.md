
**RP2040 Datasheet.pdf**

**2.15.4. Frequency Counter:**

The frequency counter measures the frequency of internal and external clocks by counting the clock edges seen over a
test interval. 
The interval is defined by counting cycles of clk_ref which must be driven either from XOSC or from a
stable external source of known frequency.

Reference:
https://github.com/raspberrypi/pico-sdk/blob/master/src/rp2_common/hardware_clocks/clocks.c#L146-L175
