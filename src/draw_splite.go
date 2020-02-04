package draw_splite
import (
	"fmt"
	"os"
	"io/ioutil"
	"bytes"

	"github.com/veandco/go-sdl2/sdl"
)


const (
	WindowTitle = "famu"
	WindowWidth = 800
	WindowHeight = 600
)

func draw_one() {

}

func draw() int {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var rect sdl.Rect
	var rects []sdl.Rect
	var err error

	window, err = sdl.CreateWindow(WindowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WindowWidth, WindowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}

	renderer.Clear()
	defer renderer.Destroy()
}





