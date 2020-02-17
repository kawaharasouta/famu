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



/**
	* memory map
	* addr						| nanigaarunoka
	*	0x0000 - 0x07ff	| wram
	*	0x0800 - 0x1fff	| wram mirror
	* 0x2000 - 0x2007	| ppu register
	* 0x2008 - 0x3xff	| puu register mirror
	*	0x4000 - 0x401f | APU I/O, PAD ??
	* 0x4020 - 0x5fff | Expanded rom
	* 0x6000 - 0x7fff	| Expanded ram
	* 0x8000 - 0xbfff	| program rom
	* 0xc000 - 0xffff	| program rom
	**/
func (b *Bus) Load(addr uint16) byte {
	if addr < 0x0800 {	//wram
		return b.wram.load(addr)
	}	else if addr < 0x1fff {
		return b.wram.load(addr % 0x0800)
	}	else if addr < 0x8000 {		//atodeiroiro
		return 0
	}	else if addr < 0xffff {
		return b.cassette[addr - 0x8000]
	}

	return 0
}
