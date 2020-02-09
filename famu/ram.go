package famu

type Mem interface {
	load(addr uint16) byte
	store(addr uint16, b byte)
	slice(begin int, end int) []byte
}



type Ram struct {
	data []byte
}
func NewRam(size int) (Mem, error) {
	return &Ram {
		data:make([]byte, size),
	}, nil
}
func (r *Ram) load(addr uint16) byte {
	return  r.data[addr]
}
func (r *Ram) store(addr uint16, b byte) {
	r.data[addr] = b
}
func (r *Ram) slice(begin int, end int) []byte {
	return r.data[begin:end]
}


