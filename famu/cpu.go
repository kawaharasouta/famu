package famu

import (
	"fmt"
)

type Cpu struct {
	A  uint8		//accumulator
	X  uint8		//Index register
	Y  uint8		//Index register
	S  uint8		//Stack pointer
	P  uint8		//Status register
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

