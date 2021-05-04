package component

import "github.com/veandco/go-sdl2/sdl"

// Component is the base interface for all
// components in this system
type Component interface {
	Position() (int32, int32)
	Size() (int32, int32)

	Draw(r *sdl.Renderer) error
}
