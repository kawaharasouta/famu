package famu

type Bus struct {
	cpu *Cpu
	ppu *Ppu
	cassette []byte
	wram Mem
	vram Mem
}

func NewBus(wram Mem, PrgRom []byte) *Bus {
	return &Bus {
		wram: wram,
		cassette: PrgRom,
	}
}



