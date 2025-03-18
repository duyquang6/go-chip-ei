package main

import (
	"log"
	"os"
)

type Chip8VM struct {
	mem  [4096]byte
	vReg [16]byte
	iReg uint16
	pc   uint16

	// support depth 16 level callstack
	stack []uint16

	delayTimer byte
	// trigger beep
	soundTimer byte

	// 64 x 32 pixels, monochrome pixel (1 pixel = 1 bit)
	screen [256]byte

	keys uint16
}

var fontSet = [...]byte{
	0xF0, 0x90, 0x90, 0x90, 0xF0, // 0: 1111, 1001, 1001, 1001, 1111, skip 4 low bit, only used 4 high bit
	0x20, 0x60, 0x20, 0x20, 0x70, // 1: 0010, 0110, 0010, 0010, 0111
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2: 1111, 0001, 1111, 1000, 1111
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x40, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
	0xF0, 0x80, 0x80, 0x80, 0xF0, // C
	0xE0, 0x90, 0x90, 0x90, 0xE0, // D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
	0xF0, 0x80, 0xF0, 0x80, 0x80, // F
}

func NewWithROMPath(romPath string) *Chip8VM {
	payload, err := os.ReadFile(romPath)
	if err != nil {
		log.Println("load rom failed, err: ", err)
		return nil
	}
	return New(payload)
}

func New(romPayload []byte) *Chip8VM {
	vm := &Chip8VM{}
	copy(vm.mem[:80], fontSet[:])
	copy(vm.mem[0x200:len(romPayload)+0x200], romPayload)

	return vm
}
