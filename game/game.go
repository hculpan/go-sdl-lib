package game

type Game interface {
	// Called from game loop
	Update() error
}
