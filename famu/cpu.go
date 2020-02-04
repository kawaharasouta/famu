package famu

import (
	"os"
)

type regs struct {
	A  uint8
	X  uint8
	Y  uint8
	S  uint8
	P  uint8
	PC uint16
}

const ( //state register
	Carry    = 0x01
	Zero     = 0x02
	Irq      = 0x04
	Decimal  = 0x08
	Break    = 0x10
	Reserved = 0x20
	Overflow = 0x40
	Negative = 0x80
)

