

The go mod init command creates a go.mod file to track your code's dependencies.

* go mod init main.go

The go mod tidy ensures that the go.mod file matches the source code in the module.

* go mod tidy

To compile program: 
* tinygo build -target=pico main.go

To compile and flash to Raspberry Pico: 
* tinygo flash -target=pico main.go

================================================

C:\Users\jgust\tinygo\programas\Raspberry_Pico\i2c_interface\i2c_scanner>go mod init main.go

go: creating new go.mod: module main.go

go: to add module requirements and sums:

        go mod tidy

C:\Users\jgust\tinygo\programas\Raspberry_Pico\i2c_interface\i2c_scanner>go mod tidy

C:\Users\jgust\tinygo\programas\Raspberry_Pico\i2c_interface\i2c_scanner>tinygo build -target=pico main.go

C:\Users\jgust\tinygo\programas\Raspberry_Pico\i2c_interface\i2c_scanner>dir

 O volume na unidade C não tem nome.
 
 O Número de Série do Volume é 709A-B58F

 Pasta de C:\Users\jgust\tinygo\programas\Raspberry_Pico\i2c_interface\i2c_scanner

16/10/2023  23:20                26 go.mod

16/10/2023  23:20           795.884 main.elf

16/10/2023  22:52             1.599 main.go

               3 arquivo(s)        797.509 bytes
               
               2 pasta(s)   30.254.874.624 bytes disponíveis

C:\Users\jgust\tinygo\programas\Raspberry_Pico\i2c_interface\i2c_scanner>

==========================================================================

VS Code - Serial Monitor (Baud rate - 115200 Bps)

---- Reopened serial port COM9 ----

Scanning I2C devices...

I2C device found at address 0X40 (64) 

I2C device found at address 0X40 (64) 

I2C device found at address 0X40 (64) 

---- Closed the serial port COM9 ----


---- Opened the serial port COM9 ----

No I2C devices found

No I2C devices found

No I2C devices found

---- Closed the serial port COM9 ----
