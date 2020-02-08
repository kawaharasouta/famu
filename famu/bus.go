package famu

type Bus struct {
	cpu *Cpu
	ppu *Ppu
	cassette []byte
	wram Mem
	vram Mem
}
