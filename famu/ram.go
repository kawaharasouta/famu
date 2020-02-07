package famu

type Mem interface {
	load(addr uint8) byte
	store(addr uint8, b byte)
	slice(begin int, end int) []byte
}

type Ram struct {
	data []byte
}
