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

// DrawWithChildren is a method that should be called by all
// classes that implement the Component interface
// It will call the Draw() method on the children first, and then
// call the DrawComponent() on the current component.
func (c *BaseComponent) DrawWithChildren(r *sdl.Renderer, f func(*sdl.Renderer) error) error {
	if err := f(r); err != nil {
		return err
	}

	for _, child := range c.Children {
		if err := child.Draw(r); err != nil {
			return err
		}
	}

	return nil
}

func (c *BaseComponent) DrawComponent(r *sdl.Renderer) error {
	return nil
}

func (c *BaseComponent) Draw(r *sdl.Renderer) error {
	if err := c.DrawWithChildren(r, c.DrawComponent); err != nil {
		return err
	}

	return nil
}

func (c *BaseComponent) AddChild(comp Component) {
	c.Children = append(c.Children, comp)
}

func (c *BaseComponent) RemoveChild(index int) {
	c.Children = append(c.Children[:index], c.Children[index+1:]...)
}
