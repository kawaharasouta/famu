package main
import (
	"fmt"
	"os"
	"io/ioutil"
	"bytes"
)


func main() {

	fmt.Println(len(os.Args))
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
	if err != nil {
		fmt.Println("readfile error.")
	}

	//header := []byte{0x4e, 0x45, 0x53, 0x1a}
	if !bytes.HasPrefix(rom, []byte{0x4e, 0x45, 0x53, 0x1a}) {
		//not nes file
		fmt.Println("hasprefix")
	}

}


func usage() {
	fmt.Println("usage: $ famu [nes file]")
}
