package famu
import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	WindowTitle = "famu"
	WindowWidth = 256
	WindowHeight = 240
	WindowSpriteWidth = 32
	WindowSpriteHeight = 30
)

type Sprite struct {
	Pixel [8][8]int
}
type screen struct {
	sprite_index [WindowSpriteWidth][WindowSpriteHeight]Sprite
}

func draw_splite_one(index int, ) {

}

func Sdl_init(sprites []Sprite, sprites_num int) int {
	var window *sdl.Window
	var renderer *sdl.Renderer
	//var rect sdl.Rect
	//var rects []sdl.Rect
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

	//renderer.SetDrawColor(255, 255, 255, 255)
	//renderer.DrawLine(0, 0, 150, 200)
	//renderer.Present()
	//sdl.PollEvent()
	//sdl.Delay(4000)

	//renderer.SetDrawColor(255, 255, 255, 255)
	//renderer.DrawPoint(150, 200)
	//renderer.Present()
	////sdl.PollEvent()
	//sdl.Delay(2000)

	var (
		index = 0
		width_index int
		height_index int
		sprite_width int
		sprite_height int
		//draw_width uint8
		//draw_height uint8
	)
	var cur_sprite *Sprite
	width_index = 0

	for index, width_index = 0, 0; index < sprites_num; index, width_index = index +1, width_index + 8 {
		cur_sprite = &sprites[index]
		if width_index >= WindowWidth {
			width_index = 0
			height_index += 3
		}

		for sprite_width = 0; sprite_width < 8; sprite_width++ {
			for sprite_height = 0; sprite_height < 8; sprite_height++ {
				//draw_width = uint8(width_index) + uint8(sprite_width)
				//draw_height = uint8(width_height) + uint8(sprite_height)
				switch cur_sprite.Pixel[sprite_width][sprite_height] {
				case 0:
					renderer.SetDrawColor(100, 100, 100, 255)
					renderer.DrawPoint(int32(width_index) + int32(sprite_width), int32(height_index) + int32(sprite_height))
				case 1:
					renderer.SetDrawColor(150, 150, 150, 255)
					renderer.DrawPoint(int32(width_index) + int32(sprite_width), int32(height_index) + int32(sprite_height))
				case 2:
					renderer.SetDrawColor(200, 200, 200, 255)
					renderer.DrawPoint(int32(width_index) + int32(sprite_width), int32(height_index) + int32(sprite_height))
				case 3:
					renderer.SetDrawColor(255, 255, 255, 255)
					renderer.DrawPoint(int32(width_index) + int32(sprite_width), int32(height_index) + int32(sprite_height))
				}
			}
		}
	}

	renderer.Present()
	sdl.PollEvent()


	return 0
}



