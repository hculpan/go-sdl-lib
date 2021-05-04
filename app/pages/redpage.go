package pages

import (
	"github.com/hculpan/gosdl/component"
	"github.com/veandco/go-sdl2/sdl"
)

type RedPage struct {
	component.BasePage

	background sdl.Color
}

func NewRedPage(name string, x, y, width, height int) *RedPage {
	p := RedPage{}
	p.Name = "RedPage"
	p.SetPosition(0, 0)
	p.SetSize(1024, 768)
	p.background = sdl.Color{R: 255, G: 0, B: 0, A: 255}

	return &p
}

func (p RedPage) Draw(r *sdl.Renderer) error {
	r.SetDrawColor(p.background.R, p.background.G, p.background.B, p.background.A)
	r.Clear()

	r.Present()

	return nil
}
