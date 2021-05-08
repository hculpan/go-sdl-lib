package component

import (
	"github.com/veandco/go-sdl2/sdl"
)

type BaseComponent struct {
	X      int32
	Y      int32
	Width  int32
	Height int32

	Children []Component
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

func (c *BaseComponent) DrawComponent(r *sdl.Renderer) error {
	return nil
}

func callDrawComponent(c Component, r *sdl.Renderer) error {
	return c.DrawComponent(r)
}

func (c *BaseComponent) Draw(r *sdl.Renderer) error {
	if err := callDrawComponent(c, r); err != nil {
		return err
	}

	for _, child := range c.Children {
		if err := callDrawComponent(child, r); err != nil {
			return err
		}
	}

	return nil
}

func (c *BaseComponent) AddChild(comp Component) {
	c.Children = append(c.Children, comp)
}

func (c *BaseComponent) RemoveChild(index int) {
	c.Children = append(c.Children[:index], c.Children[index+1:]...)
}
