/*
Raspberry Pico 2 - RP2350 - teste de SRAM - DMA + Spinlock via SIO - Teste
Gustavo Murta 2025/06/24
tinygo version 0.37.0 windows/amd64 (using go version go1.24.2 and LLVM version 19.1.2)
C:\Users\jgust\tinygo\programas\Raspberry_Pico\sram_pico2
tinygo flash -target pico2 main.go
3.1.4. Hardware Spinlocks - Exemplo Muito Bom !
*/

package main

import (
	"fmt"
	"time"
	"unsafe"
)

const (
	// DMA Registers
	DMA_BASE            = 0x50000000       // Base address for DMA controller
	DMA_CH0_READ_ADDR   = DMA_BASE + 0x000 // DMA Channel 0 Read Address pointer
	DMA_CH0_WRITE_ADDR  = DMA_BASE + 0x004 // DMA Channel 0 Write Address pointer
	DMA_CH0_TRANS_COUNT = DMA_BASE + 0x008 // DMA Channel 0 Transfer Count
	DMA_CH0_CTRL_TRIG   = DMA_BASE + 0x00C // DMA Channel 0 Control and Status
	DMA_CH0_CTRL_ENABLE = 1 << 0           // Enable DMA Channel 0
	DMA_CH0_DATA_SIZE_8 = 0 << 2           // Data size 8 bits
	DMA_CH0_INCR_READ   = 1 << 4           // Increment read address
	DMA_CH0_INCR_WRITE  = 1 << 5           // Increment write address
	DMA_CH0_CHAIN_TO_0  = 0 << 11          // Chain to Channel 0
	DMA_CH0_BUSY_FLAG   = 1 << 24          // DMA Channel 0 Busy Flag

	BUFFER_SIZE = 32 // Size of the buffer for DMA transfer (32 bytes)

	// SIO Spinlock base
	SIO_BASE       = 0xD0000000       // Base address for SIO (Single cycle I/O)
	SPINLOCK_BASE  = SIO_BASE + 0x100 // Spinlock base address
	SPINLOCK_INDEX = 0                // Spinlock index (0-31)
)

var (
	sourceBuffer [BUFFER_SIZE]uint8 // Source buffer for DMA transfer
	destBuffer   [BUFFER_SIZE]uint8 // Destination buffer for DMA transfer

	dmaDone = false // Flag to indicate DMA completion
)

func main() {

	time.Sleep(3 * time.Second) // Sleep to catch prints.
	println("SRAM - DMA + Spinlock testes")

	go core1() // Core 1 will read the data

	for {
		acquireLock(SPINLOCK_INDEX) // Acquire the lock to write data
		println("Core 0: Preenchendo o Source buffer")

		for i := 0; i < BUFFER_SIZE; i++ { // Fill the source buffer with data
			sourceBuffer[i] = uint8(i + 1)       // Fill with values 1 to BUFFER_SIZE
			fmt.Printf("%02X ", sourceBuffer[i]) // Print the data being written to the source buffer
		}
		fmt.Printf("\n")

		releaseLock(SPINLOCK_INDEX) // Release the lock after writing data

		startDMA() // Start DMA transfer

		for !dmaDone { // Wait for DMA to complete
		}
		dmaDone = false // Reset the DMA done flag for the next transfer

		println("Core 0: DMA finalizada - buffer filled.")
		time.Sleep(1 * time.Second) // Sleep to allow Core 1 to read the data
	}
}

func core1() {
	for {
		time.Sleep(1100 * time.Millisecond) // Sleep to allow Core 0 to write data

		acquireLock(SPINLOCK_INDEX) // Acquire the lock to read data
		print("Core 1: Dados recebidos: ")
		for i := 0; i < BUFFER_SIZE; i++ { // Read the data from the destination buffer
			print(destBuffer[i], " ") // Print the data received
		}
		println()
		println()
		releaseLock(SPINLOCK_INDEX) // Release the lock after reading data
	}
}

// DMA – transferência por canal 0
func startDMA() {
	readAddr := uint32(uintptr(unsafe.Pointer(&sourceBuffer[0]))) // Read address for DMA transfer
	writeAddr := uint32(uintptr(unsafe.Pointer(&destBuffer[0])))  // Write address for DMA transfer
	// fmt.Printf(" DMA read adress = 0x%X ", readAddr)
	// fmt.Printf(" DMA write adress = 0x%X \n", writeAddr)

	*(*uint32)(unsafe.Pointer(uintptr(DMA_CH0_READ_ADDR))) = readAddr      // Set the read address for DMA Channel 0
	*(*uint32)(unsafe.Pointer(uintptr(DMA_CH0_WRITE_ADDR))) = writeAddr    // Set the write address for DMA Channel 0
	*(*uint32)(unsafe.Pointer(uintptr(DMA_CH0_TRANS_COUNT))) = BUFFER_SIZE // Set the transfer count for DMA Channel 0

	ctrl := DMA_CH0_CTRL_ENABLE | // Enable DMA Channel 0
		DMA_CH0_DATA_SIZE_8 | // Set data size to 8 bits
		DMA_CH0_INCR_READ | // Increment read address
		DMA_CH0_INCR_WRITE | // Increment write address
		DMA_CH0_CHAIN_TO_0 // Chain to Channel 0

	*(*uint32)(unsafe.Pointer(uintptr(DMA_CH0_CTRL_TRIG))) = uint32(ctrl) // Set the control and trigger register for DMA Channel 0

	for *(*uint32)(unsafe.Pointer(uintptr(DMA_CH0_CTRL_TRIG)))&DMA_CH0_BUSY_FLAG != 0 { // Wait until DMA Channel 0 is not busy
	}
	dmaDone = true // Set the DMA done flag to true to indicate completion
}

// Spinlock via SIO
func spinlockAddr(index uint8) uintptr { // Calculate the address of the spinlock based on the index
	return SPINLOCK_BASE + uintptr(index)*4 // Each spinlock is 4 bytes, so we multiply the index by 4
}

func acquireLock(index uint8) { // Acquire the spinlock at the specified index
	for {
		val := *(*uint32)(unsafe.Pointer(spinlockAddr(index))) // Read the value of the spinlock
		if val == 1 {                                          // If the spinlock is already acquired
			break
		}
	}
}

func releaseLock(index uint8) { // Release the spinlock at the specified index
	*(*uint32)(unsafe.Pointer(spinlockAddr(index))) = 0 // Set the spinlock value to 0 to release it
}
