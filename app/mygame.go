package app

import (
	"github.com/hculpan/gosdl/component"
)

type MyGame struct {
	cycle       int
	currentPage int
}

const (
	BLACK_PAGE = iota
	RED_PAGE
)

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
