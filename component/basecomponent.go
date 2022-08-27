package component

import (
	"github.com/veandco/go-sdl2/sdl"
)

type BaseComponent struct {
	X        int32
	Y        int32
	Width    int32
	Height   int32
	Children []Component
}

var Scaling float32 = 1.0

/**************************************
* Boiletplate functions
*
* All components should include these
***************************************/

func (c *BaseComponent) KeyEvent(event *sdl.KeyboardEvent) bool {
	return PassKeyEventToChildren(event, c.Children)
}

func (c *BaseComponent) Draw(r *sdl.Renderer) error {
	return DrawParentAndChildren(r, c)
}

func (c *BaseComponent) BaseMouseButtonEvent(event *sdl.MouseButtonEvent) bool {
	event.X = int32(float32(event.X) * Scaling)
	event.Y = int32(float32(event.Y) * Scaling)
	return c.MouseButtonEvent(event)
}

func (c *BaseComponent) MouseButtonEvent(event *sdl.MouseButtonEvent) bool {
	if c.IsPointInComponent(event.X, event.Y) {
		return PassMouseButtonEventToChildren(event, c.Children)
	}

	return false
}

/**************************************
* End boilerplate functions
***************************************/

func PassKeyEventToChildren(event *sdl.KeyboardEvent, children []Component) bool {
	for _, child := range children {
		if child.KeyEvent(event) {
			return true
		}
	}

	return false
}

func DrawParentAndChildren(r *sdl.Renderer, c Component) error {
	if err := c.DrawComponent(r); err != nil {
		return err
	}

	for _, child := range c.GetChildren() {
		if err := child.Draw(r); err != nil {
			return err
		}
	}

	return nil
}

func PassMouseButtonEventToChildren(event *sdl.MouseButtonEvent, children []Component) bool {
	for _, child := range children {
		if child.IsPointInComponent(event.X, event.Y) && child.MouseButtonEvent(event) {
			return true
		}
	}

	return false
}

func (c BaseComponent) IsPointInComponent(x, y int32) bool {
	return x >= c.X && x <= c.Width+c.X && y >= c.Y && y <= c.Height+c.Y
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

func (c BaseComponent) GetChildren() []Component {
	return c.Children
}

func (c *BaseComponent) AddChild(comp Component) {
	c.Children = append(c.Children, comp)
}

func (c *BaseComponent) RemoveChild(index int) {
	c.Children = append(c.Children[:index], c.Children[index+1:]...)
}
