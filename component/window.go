package component

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Window represents the main object to display text output
type Window struct {
	Width  int32
	Height int32
	Title  string

	window   *sdl.Window
	renderer *sdl.Renderer

	Background sdl.Color
}

// NewWindow creates a new Window object
func NewWindow(width, height int32, title string, background sdl.Color) *Window {
	return &Window{Width: int32(width), Height: int32(height), Title: title, Background: background}
}

func SetupSDL() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Initializing SDL: ", err)
		return err
	}

	return nil
}

// CleanUp cleans up all the resources creates by the Window object
func (s *Window) CleanUp() {
	s.renderer.Destroy()
	s.window.Destroy()
}

// Show initializes the main Window and shows it
func (s *Window) Show() error {
	window, err := sdl.CreateWindow(
		s.Title,
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		int32(s.Width),
		int32(s.Height),
		sdl.WINDOW_OPENGL,
	)
	if err != nil {
		fmt.Println("Initializing window: ", err)
		return err
	}
	s.window = window

	renderer, err := sdl.CreateRenderer(s.window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Initializing renderer: ", err)
		return err
	}
	s.renderer = renderer

	return nil
}

// DrawScreen draws the active page
func (s *Window) DrawScreen() error {
	s.renderer.SetDrawColor(s.Background.R, s.Background.G, s.Background.B, s.Background.A)
	s.renderer.Clear()
	if err := ActivePage.Draw(s.renderer); err != nil {
		panic(err)
	}
	s.renderer.Present()
	return nil
}
