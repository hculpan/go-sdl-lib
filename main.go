package main

import (
	"fmt"

	"github.com/hculpan/gosdl/app"
	"github.com/hculpan/gosdl/app/pages"
	"github.com/hculpan/gosdl/component"
	"github.com/hculpan/gosdl/game"
)

const (
	gameWidth  = 1600
	gameHeight = 1024
)

func main() {
	game := game.NewGameController(component.NewWindow(gameWidth, gameHeight, "GoSDL"), Setup())
	Setup()
	if err := game.Run(); err != nil {
		fmt.Println(err.Error())
	}
}

// TODO: Must update this with specific setup for game
func Setup() game.Game {
	component.RegisterPage(pages.NewMainPage("MainPage", 0, 0, int(gameWidth), int(gameHeight)))
	component.RegisterPage(pages.NewRedPage("RedPage", 0, 0, int(gameWidth), int(gameHeight)))

	return &app.MyGame{}
}
