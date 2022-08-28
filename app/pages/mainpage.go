package pages

import (
	"fmt"

	"github.com/hculpan/go-sdl-lib/component"
	"github.com/veandco/go-sdl2/sdl"
)

type MainPage struct {
	component.BasePage
}

func NewMainPage(name string, x, y, width, height int32) *MainPage {
	fmt.Println("drawing main page")
	p := MainPage{}
	p.Initialize()
	p.Name = "MainPage"
	p.SetPosition(0, 0)
	p.SetSize(width, height)

	p.AddChild(component.NewRectangleComponent(
		0, 0, width, height, sdl.Color{R: 0, G: 0, B: 0, A: 255},
	))

	return &p
}
