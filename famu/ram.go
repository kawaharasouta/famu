package famu

import (
	"fmt"
)

type Mem interface {
	load(addr uint16)
	store(addr uint16, b byte)
	slice(begin int, end int)
}



type Ram struct {
	data []byte
}

func NewRam(size int) (Mem, error) {
	return &Ram {
		data:make([]byte, size),
	}, nil
}

func (r *Ram) load(addr uint16) {
	fmt.Println("load\n")
}
func (r *Ram) store(addr uint16, b byte) {
	fmt.Println("store\n")
}
func (r *Ram) slice(begin int, end int) {
	fmt.Println("slice\n")
}

