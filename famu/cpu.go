package famu

//import "fmt"

type Cpu struct {
	A  byte		//accumulator
	X  byte		//Index register
	Y  byte		//Index register
	S  byte		//Stack pointer
	P  byte		//Status register
	PC uint16		//Program counter

	cycle uint64
	bus *Bus
	intrrupt func()
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




func (c *Cpu) jmp(addr uint16) {
	c.PC = addr
}





/* warikomi */

func (c *Cpu) nmi() {

}
func (c *Cpu) reset() { // sample ha koredake toriaezu yaruyo
	//push toka iroiro wakarannkedo


	c.jmp(c.bus.Loadw(0xfffc))


}
func (c *Cpu) irq() {

}
func (c *Cpu) brk() {

}


/* execute */
func (c *Cpu) fetch() byte {
	addr := c.PC
	c.PC++
	return c.bus.Load(addr)
}

func (c *Cpu) decode(op byte) (func(uint16), func()) {
	var (
		inst func(uint16)
		addr_mode func()
	)

	return inst, addr_mode
}

func (c *Cpu) exec(inst func(uint16), addr_mode func()) {
	//switch inst {
	//case "lda":
	//	fmt.Println("lda\n")
	//default:
	//	fmt.Println("no match instruction.\n")
	//}
}

func run(c *Cpu) {
	c.reset()

	for {
		op := c.fetch()
		inst, addr_mode := c.decode(op)
		c.exec(inst, addr_mode)
	}
}
