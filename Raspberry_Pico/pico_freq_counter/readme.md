
**RP2040 Datasheet.pdf**

**2.15.4. Frequency Counter:**

The frequency counter measures the frequency of internal and external clocks by counting the clock edges seen over a
test interval. 
The interval is defined by counting cycles of clk_ref which must be driven either from XOSC or from a
stable external source of known frequency.

**2.15.6.2. Using the frequency counter:**

To use the frequency counter, the programmer must:
- Set the reference frequency: clk_ref
- Set the mux position of the source they want to measure. See FC0_SRC
- Wait for the DONE status bit in FC0_STATUS to be set
- Read the result

**Reference:**

https://github.com/raspberrypi/pico-sdk/blob/master/src/rp2_common/hardware_clocks/clocks.c#L146-L175
