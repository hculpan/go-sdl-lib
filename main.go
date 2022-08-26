package main

import (
	"fmt"

	"github.com/hculpan/go-sdl-lib/app"
	"github.com/hculpan/go-sdl-lib/component"
	"github.com/hculpan/go-sdl-lib/game"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	gameWidth  int32 = 800
	gameHeight int32 = 600
)

func main() {
	// TODO: Set to the desired default window background
	windowBackground := sdl.Color{R: 255, G: 255, B: 255, A: 255}

	if err := component.SetupSDL(); err != nil {
		fmt.Println(err)
		return
	}

	w := component.NewWindow(gameWidth, gameHeight, "GoSDL", windowBackground)
	game := game.NewGameController(
		w,
		GetGame(w.Width, w.Height))
	if err := game.Run(); err != nil {
		fmt.Println(err.Error())
	}
}

// TODO: Must update this with specific setup for game
func GetGame(w, h int32) game.Game {
	return app.NewMyGame(w, h)
}
