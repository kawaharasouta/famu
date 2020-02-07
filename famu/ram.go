package famu

type Mem interface {
	load(addr uint16) byte
	store(addr uint16, b byte)
	slice(begin int, end int) []byte
}

type Ram struct {
	data []byte
}



