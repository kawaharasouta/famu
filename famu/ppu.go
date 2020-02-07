package famu


type Ppu struct {
	PpuCtrl uint8   //ctrl register1 (0x2000)
	PpuMask uint8   //ctrl register2 (0x2001)
	PpuStatus uint8   //ppu status (0x2002)
	OamAddr uint8   //splite addr (0x2003)
	OamData uint8   //xplite data (0x2004)
//	scrollFirst bool   // for 0x2005
//	PpuScrollX  byte   // 0x2005(1)
//	PpuScrollY  byte   // 0x2005(2)
	PpuAddr uint16   //ppu mem addr (0x2006)
//	isLowerAddr bool   // for 0x2006
	PpuData uint8   //ppu mem data (0x2007)

	cycle       uint64
	vram        Mem
	bus         *Bus
	renderer    *Renderer

	// Sprite RAM
	spriteBuffer [64]*Sprite
	spriteRam Mem
	vramBuf byte
}
