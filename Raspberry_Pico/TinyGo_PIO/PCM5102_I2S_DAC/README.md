Testing I2S DAC PCM5102 with Raspberry Pico 

Based on: **https://github.com/tinygo-org/pio/tree/main**

```
PS C:\Users\jgust\tinygo\programas\Raspberry_Pico\PCM5102> go mod init main.go
go: creating new go.mod: module main.go
go: to add module requirements and sums:
        go mod tidy

PS C:\Users\jgust\tinygo\programas\Raspberry_Pico\PCM5102> go mod tidy
go: finding module for package github.com/tinygo-org/pio/rp2-pio
go: finding module for package github.com/tinygo-org/pio/rp2-pio/piolib
go: found github.com/tinygo-org/pio/rp2-pio in github.com/tinygo-org/pio v0.0.0-20231101233832-892dc73511e3
go: found github.com/tinygo-org/pio/rp2-pio/piolib in github.com/tinygo-org/pio v0.0.0-20231101233832-892dc73511e3
PS C:\Users\jgust\tinygo\programas\Raspberry_Pico\PCM5102>
```

```
PS C:\Users\jgust\tinygo\programas\Raspberry_Pico\PCM5102> tinygo build -target=pico main.go
PS C:\Users\jgust\tinygo\programas\Raspberry_Pico\PCM5102> tinygo flash -target=pico main.go
PS C:\Users\jgust\tinygo\programas\Raspberry_Pico\PCM5102> 
```
