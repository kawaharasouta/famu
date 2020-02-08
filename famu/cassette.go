package famu

import (
	"fmt"
	"os"
	"io/ioutil"
	"bytes"
)


const (
	nes_header = 0x0010
	PRG_unit = 0x4000
	CHR_unit = 0x2000
)


type Cassette struct {
	prgRom []byte
	chrRom []byte
}


func NewCassette(filepath string) (*Cassette, error) {
	rom, err := ioutil.ReadFile(filepath)
	if err != nil || !bytes.HasPrefix(rom, []byte{0x4e, 0x45, 0x53, 0x1a}) {
		//fmt.Fprintf(os.Stderr, "Failed to readfile: %s\n", err)
		fmt.Fprintf(os.Stderr, "Failed to readfile\n")
	}

	PRG_size := rom[4]
	CHR_size := rom[5]

	CHR_start := nes_header + int(PRG_size) * PRG_unit		//(top + size * 16KB)
	CHR_end := CHR_start + int(CHR_size) * CHR_unit

	prgRom := rom[nes_header:CHR_start]
	chrRom := rom[CHR_start:CHR_end]

	return &Cassette {
		prgRom:prgRom,
		chrRom:chrRom,
	}, nil
}


type Rom interface {
	PrgRom() []byte
	ChrRom() []byte
}
func PrgRom(c *Cassette) []byte {
	return c.prgRom
}
func ChrRom(c *Cassette) []byte {
	return c.chrRom
}
