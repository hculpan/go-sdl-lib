package components

import (
	"github.com/hculpan/go-sdl-lib/component"
	"github.com/veandco/go-sdl2/sdl"
)

type RectangleComponent struct {
	component.BaseComponent

	Background sdl.Color
}

func NewRectangleComponent(x, y, width, height int32, background sdl.Color) *RectangleComponent {
	result := &RectangleComponent{
		Background: background,
	}

	result.SetPosition(x, y)
	result.SetSize(int32(width), int32(height))

	return result
}

func (c *RectangleComponent) DrawComponent(r *sdl.Renderer) error {
	r.SetDrawColor(c.Background.R, c.Background.G, c.Background.B, c.Background.A)
	rect := sdl.Rect{X: c.X, Y: c.Y, W: c.Width, H: c.Height}
	r.FillRect(&rect)

	return nil
}

func (c *RectangleComponent) Draw(r *sdl.Renderer) error {
	if err := component.DrawParentAndChildren(r, c); err != nil {
		return err
	}

	return nil
}
