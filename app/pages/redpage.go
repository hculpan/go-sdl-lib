package pages

import (
	"github.com/hculpan/gosdl/app/components"
	"github.com/hculpan/gosdl/component"
	"github.com/veandco/go-sdl2/sdl"
)

type RedPage struct {
	component.BasePage
}

func NewRedPage(name string, x, y, width, height int32) *RedPage {
	p := RedPage{}
	p.Name = "RedPage"
	p.SetPosition(0, 0)
	p.SetSize(width, height)
	p.AddChild(components.NewRectangleComponent(
		0, 0, width, height/2,
		sdl.Color{R: 255, G: 0, B: 0, A: 255},
	))
	p.AddChild(components.NewRectangleComponent(
		0, height/2+1, width, height/2,
		sdl.Color{R: 0, G: 255, B: 0, A: 255},
	))

	return &p
}
