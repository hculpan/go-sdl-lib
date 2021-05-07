package component

import (
	"fmt"

	"github.com/hculpan/go-sdl-lib/static"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Window represents the main object to display text output
type Window struct {
	Width  int32
	Height int32
	Title  string

	window   *sdl.Window
	renderer *sdl.Renderer
	font     *ttf.Font

	Background sdl.Color
}

// NewWindow creates a new Window object
func NewWindow(width, height int32, title string, background sdl.Color) *Window {
	return &Window{Width: int32(width), Height: int32(height), Title: title, Background: background}
}

func (s *Window) Setup() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("Initializing SDL: ", err)
		return err
	}

	if err := ttf.Init(); err != nil {
		fmt.Println("Initializing SDL: ", err)
		return err
	}

	if err := s.initializeFonts(); err != nil {
		fmt.Println("Initializing fonts: ", err)
		return err
	}

	return nil
}

// CleanUp cleans up all the resources creates by the Window object
func (s *Window) CleanUp() {
	s.font.Close()
	s.renderer.Destroy()
	s.window.Destroy()
}

func (s *Window) initializeFonts() error {
	font, err := static.InitializeFont("OldComputerManualMonospaced-KmlZ.ttf")
	if err != nil {
		return err
	}
	s.font = font

	return nil
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
