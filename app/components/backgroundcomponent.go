package components

import (
	"github.com/hculpan/gosdl/component"
	"github.com/veandco/go-sdl2/sdl"
)

type BackgroundComponent struct {
	component.BaseComponent

	Background sdl.Color
}

func NewBackgroundComponent(width, height int, background sdl.Color) *BackgroundComponent {
	result := &BackgroundComponent{
		Background: background,
	}

	result.SetPosition(0, 0)
	result.SetSize(int32(width), int32(height))

	return result
}

func (c *BackgroundComponent) Draw(r *sdl.Renderer) error {
	r.SetDrawColor(c.Background.R, c.Background.G, c.Background.B, c.Background.A)
	r.Clear()

	r.Present()

	return nil
}
