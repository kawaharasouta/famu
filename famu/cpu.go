package famu

import (
)

type Cpu struct {
	A  byte		//accumulator
	X  byte		//Index register
	Y  byte		//Index register
	S  byte		//Stack pointer
	P  byte		//Status register
	PC uint16		//Program counter

	cycle uint64
	irq func()
}

const ( //Status register
	Carry    = 0x01
	Zero     = 0x02
	Irq      = 0x04
	Decimal  = 0x08
	Break    = 0x10
	Reserved = 0x20
	Overflow = 0x40
	Negative = 0x80
)

