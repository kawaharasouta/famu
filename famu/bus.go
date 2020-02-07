package nes

type Bus struct {
	cpu *Cpu
	ppu *Ppu
	rom []byte
	wram Mem
	vram Mem
}
