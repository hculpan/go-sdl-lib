package pages

import (
	"github.com/hculpan/gosdl/component"
	"github.com/veandco/go-sdl2/sdl"
)

type MainPage struct {
	component.BasePage

	background sdl.Color
	foreground sdl.Color
}

func NewMainPage(name string, x, y, width, height int) *MainPage {
	p := MainPage{}
	p.Name = "MainPage"
	p.SetPosition(0, 0)
	p.SetSize(1024, 768)
	p.background = sdl.Color{R: 0, G: 0, B: 0, A: 0}
	p.foreground = sdl.Color{R: 255, G: 255, B: 255, A: 255}

	return &p
}

func (p MainPage) Draw(r *sdl.Renderer) error {
	r.SetDrawColor(p.background.R, p.background.G, p.background.B, p.background.A)
	r.Clear()

	r.Present()

	return nil
}
