package main
import (
	"fmt"
	"os"
	"io/ioutil"
	"bytes"

	"github.com/kawaharasouta/famu/famu"
)

const (
	nes_header = 0x0010
	PRG_unit = 0x4000
	CHR_unit = 0x2000
)

func main() {

	//fmt.Println(len(os.Args))
	//arg1 := os.Args[0]
	////message := os.Sprintf("%s, %s", arg1, arg2)
	//fmt.Println(arg1, arg2)
	////fmt.Println(mes)

	//args
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}
	filepath := os.Args[1]

	rom, err := ioutil.ReadFile(filepath)
	if err != nil || !bytes.HasPrefix(rom, []byte{0x4e, 0x45, 0x53, 0x1a}) {
		//fmt.Fprintf(os.Stderr, "Failed to readfile: %s\n", err)
		fmt.Fprintf(os.Stderr, "Failed to readfile\n")
	}

	PRG_size := rom[4]
	CHR_size := rom[5]
	//fmt.Println(PRG_size)
	//fmt.Println(CHR_size)

	CHR_start := nes_header + int(PRG_size) * PRG_unit		//(top + size * 16KB)
	CHR_end := CHR_start + int(CHR_size) * CHR_unit
	sprites_num := int(CHR_size) * CHR_unit / 16

	//tyousei
	//CHR_start = CHR_start + 0x0210

	fmt.Println(CHR_start)
	fmt.Println(CHR_end)
	fmt.Println(sprites_num)

	sprites := make([]famu.Sprite, 0)
	var (
		i = 0
		j = 0
		n = 0
	)
	cur := CHR_start
	fmt.Println(rom[cur+1])

	for	n = 0; n < sprites_num; n++ {
		sprites = append(sprites, famu.Sprite{})
		//for i = 0; i < 8; i++ {
		//	for j = 0; j < 8; j++ {
		//		sprites[n].pixel[i][j] = 0
		//	}
		//}
		for i = 0; i < 16; i++ {
			for j = 0; j < 8; j++ {
				if (rom[cur + i] & (0x01 << j)) != 0 {
					sprites[n].Pixel[i % 8][j] += 0x01 << (i / 8)
				}
			}
		}
		fmt.Println(sprites[n])
		cur += 0x10
	}

	ret := famu.Sdl_init(sprites, sprites_num)

	if ret == 1 {
		fmt.Println("ret1")
	}
}


func usage() {
	fmt.Println("usage: $ famu [nes file]")
}


