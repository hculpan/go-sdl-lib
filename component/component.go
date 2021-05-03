package component

// Component is the base interface for all
// components in this system
type Component interface {
	Position() (int, int)
	Size() (int, int)
}
