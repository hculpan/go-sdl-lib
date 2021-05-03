package pages

import "github.com/veandco/go-sdl2/sdl"

type RedPage struct {
	x      int
	y      int
	width  int
	height int
	name   string

	background sdl.Color
}

func NewRedPage(name string, x, y, width, height int) RedPage {
	p := RedPage{
		name:   "RedPage",
		x:      0,
		y:      0,
		width:  1024,
		height: 768,
	}
	p.background = sdl.Color{R: 255, G: 0, B: 0, A: 255}

	return p
}

func (p RedPage) Position() (int, int) {
	return p.x, p.y
}

func (p RedPage) Size() (int, int) {
	return p.width, p.height
}

func (p RedPage) Name() string {
	return p.name
}

func (p RedPage) DrawPage(r *sdl.Renderer) error {
	r.SetDrawColor(p.background.R, p.background.G, p.background.B, p.background.A)
	r.Clear()

	r.Present()

	return nil
}
