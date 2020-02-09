package main
import (
	"fmt"
	"os"
//	"bytes"

	"github.com/kawaharasouta/famu/famu"
)

func main() {
	//args
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}
	filepath := os.Args[1]

	c, err := famu.NewCassette(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cassette error\n")
	}
	rom := c.ChrRom()
	sprites_num := len(rom) /  16
	fmt.Println(sprites_num)

	sprites := make([]famu.Sprite, 0)
	var (
		i = 0
		j = 0
		n = 0
	)
	cur := 0

	for	n = 0; n < sprites_num; n++ {
		sprites = append(sprites, famu.Sprite{})
		for i = 0; i < 16; i++ {
			for j = 0; j < 8; j++ {
				if (rom[cur + i] & (0x01 << j)) != 0 {
					sprites[n].Pixel[i % 8][j] += 0x01 << (i / 8)
				}
			}
		}
		//fmt.Println(sprites[n])
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


