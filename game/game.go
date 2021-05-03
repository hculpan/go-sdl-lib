package game

import (
	"fmt"

	"github.com/hculpan/gosdl/component"
	"github.com/hculpan/gosdl/static"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Game represents the main object to display text output
type Game struct {
	width  int32
	height int32
	title  string

	window   *sdl.Window
	renderer *sdl.Renderer
	font     *ttf.Font
}

// NewGame creates a new Game object
func NewGame(width int, height int, title string) *Game {
	return &Game{width: int32(width), height: int32(height), title: title}
}

// CleanUp cleans up all the resources creates by the Game object
func (s *Game) CleanUp() {
	s.font.Close()
	s.renderer.Destroy()
	s.window.Destroy()
}

func (s *Game) initializeFonts() error {
	font, err := static.InitializeFont("OldComputerManualMonospaced-KmlZ.ttf")
	if err != nil {
		return err
	}
	s.font = font

	return nil
}

// Show initializes the main Game and shows it
func (s *Game) Show() error {
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

	window, err := sdl.CreateWindow(
		s.title,
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		int32(s.width),
		int32(s.height),
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
func (s *Game) DrawScreen() {
	component.ActivePage.DrawPage(s.renderer)
}
