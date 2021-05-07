package app

import (
	"github.com/hculpan/go-sdl-lib/app/pages"
	"github.com/hculpan/go-sdl-lib/component"
)

type MyGame struct {
	component.BaseGame

	cycle       int
	currentPage int
}

const (
	BLACK_PAGE = iota
	RED_PAGE
)

func NewMyGame(gameWidth, gameHeight int32) *MyGame {
	result := &MyGame{}

	result.Initialize(gameWidth, gameHeight)
	result.RegisterPages(result.LoadPages())

	return result
}

func (g *MyGame) LoadPages() []component.Page {
	return []component.Page{
		pages.NewMainPage("MainPage", 0, 0, g.GameWidth, g.GameHeight),
		pages.NewRedPage("RedPage", 0, 0, g.GameWidth, g.GameHeight),
	}
}

func (g *MyGame) Update() error {
	g.cycle += 1
	if g.cycle%5000 == 0 {
		switch g.currentPage {
		case BLACK_PAGE:
			component.SwitchPage("RedPage")
			g.currentPage = RED_PAGE
		case RED_PAGE:
			component.SwitchPage("MainPage")
			g.currentPage = BLACK_PAGE
		}
	}
	return nil
}
