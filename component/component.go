package component

import "github.com/veandco/go-sdl2/sdl"

// Component is the base interface for all
// components in this system
type Component interface {
	Position() (int32, int32)
	Size() (int32, int32)

	AddChild(Component)
	RemoveChild(int)
	GetChildren() []Component

	Draw(r *sdl.Renderer) error
	DrawComponent(r *sdl.Renderer) error

	KeyEvent(event *sdl.KeyboardEvent) bool
	MouseButtonEvent(event *sdl.MouseButtonEvent) bool
	BaseMouseButtonEvent(event *sdl.MouseButtonEvent) bool

	IsPointInComponent(x, y int32) bool
}

type OnActionFunc func()
