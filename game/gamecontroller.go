package game

import (
	"github.com/hculpan/gosdl/component"
	"github.com/veandco/go-sdl2/sdl"
)

type GameController struct {
	width  int
	height int

	game Game

	window *component.Window
}

func NewGameController(window *component.Window, game Game) GameController {
	return GameController{width: int(window.Width), height: int(window.Height), window: window, game: game}
}

func (g *GameController) Setup() {
}

func (g *GameController) DrawFrame() error {
	return g.window.DrawScreen()
}

func (g *GameController) Update() error {
	return nil
}

func (g *GameController) Cleanup() {
	g.window.CleanUp()
}

func (g *GameController) Run() error {
	g.Setup()
	if err := g.window.Setup(); err != nil {
		return err
	}
	defer g.Cleanup()
	if err := g.window.Show(); err != nil {
		return err
	}

	running := true
	for running {
		event := sdl.PollEvent()
		if event != nil {
			switch event := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				keycode := sdl.GetKeyFromScancode(event.Keysym.Scancode)
				if event.Type == sdl.KEYDOWN && keycode == sdl.K_ESCAPE {
					running = false
				}
			}
		}

		if err := g.game.Update(); err != nil {
			return err
		}

		if err := g.window.DrawScreen(); err != nil {
			return err
		}
	}

	return nil
}
