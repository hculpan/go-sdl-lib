package pages

import "github.com/veandco/go-sdl2/sdl"

type MainPage struct {
	x      int
	y      int
	width  int
	height int
	name   string

	background sdl.Color
	foreground sdl.Color
}

func NewMainPage(name string, x, y, width, height int) MainPage {
	p := MainPage{
		name:   "MainPage",
		x:      0,
		y:      0,
		width:  1024,
		height: 768,
	}
	p.background = sdl.Color{R: 0, G: 0, B: 0, A: 0}
	p.foreground = sdl.Color{R: 255, G: 255, B: 255, A: 255}

	return p
}

func (p MainPage) Position() (int, int) {
	return p.x, p.y
}

func (p MainPage) Size() (int, int) {
	return p.width, p.height
}

func (p MainPage) Name() string {
	return p.name
}

func (p MainPage) DrawPage(r *sdl.Renderer) error {
	r.SetDrawColor(p.background.R, p.background.G, p.background.B, p.background.A)
	r.Clear()

	r.Present()

	return nil
}
