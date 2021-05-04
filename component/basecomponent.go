package component

import "github.com/veandco/go-sdl2/sdl"

type BaseComponent struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
}

func (c BaseComponent) Position() (int32, int32) {
	return c.X, c.Y
}

func (c *BaseComponent) SetPosition(x, y int32) {
	c.X = x
	c.Y = y
}

func (c BaseComponent) Size() (int32, int32) {
	return c.Width, c.Height
}

func (c *BaseComponent) SetSize(width, height int32) {
	c.Width = width
	c.Height = height
}

func (c *BaseComponent) Draw(r *sdl.Renderer) error {
	return nil
}
