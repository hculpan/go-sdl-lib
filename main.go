package main

import (
	"fmt"

	"github.com/hculpan/gosdl/component"
	"github.com/hculpan/gosdl/game"
	"github.com/hculpan/gosdl/pages"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	gameeenWidth  = 1600
	gameeenHeight = 1024
)

func main() {
	setup()

	game := game.NewGame(gameeenWidth, gameeenHeight, "GoSDL")
	if err := game.Show(); err != nil {
		fmt.Println(err.Error())
		return
	}
	defer game.CleanUp()
	defer cleanup()

	for {
		event := sdl.PollEvent()
		if event != nil {
			switch event := event.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.KeyboardEvent:
				keycode := sdl.GetKeyFromScancode(event.Keysym.Scancode)
				if event.Type == sdl.KEYDOWN && keycode == sdl.K_ESCAPE {
					return
				}
			}
		}
		game.DrawScreen()
	}
}

func setup() {
	component.RegisterPage(pages.NewMainPage("MainPage", 0, 0, gameeenWidth, gameeenHeight))
}

func cleanup() {

}
