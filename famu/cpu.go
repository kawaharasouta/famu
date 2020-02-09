package famu

import "fmt"

type Cpu struct {
	A  byte		//accumulator
	X  byte		//Index register
	Y  byte		//Index register
	S  byte		//Stack pointer
	P  byte		//Status register
	PC uint16		//Program counter

	cycle uint64
	bus *Bus
	irq func()
}

const ( //Status register
	Negative = 0x80
	Overflow = 0x40
	Reserved = 0x20
	Break    = 0x10
	Decimal  = 0x08
	Irq      = 0x04
	Zero     = 0x02
	Carry    = 0x01
)


func NewCpu(bus *Bus) *Cpu{
	return &Cpu{
		bus:bus,
	}
}



/* status set */



/* status unset */



/* status isset* /



/* instruction */



/* exec */
func exec(inst, ope, mode string) {
	switch inst {
	case "lda":
		fmt.Println("lda\n")
	default:
		fmt.Println("no match instruction.\n")
	}
}
