package game

import (
	"github.com/hculpan/go-sdl-lib/component"
	"github.com/veandco/go-sdl2/sdl"
)

type GameController struct {
	Game Game

	Window *component.Window
}

func NewGameController(window *component.Window, game Game) GameController {
	return GameController{Window: window, Game: game}
}

func (g *GameController) Setup() {
}

func (g *GameController) DrawFrame() error {
	return g.Window.DrawScreen()
}

func (g *GameController) Update() error {
	return nil
}

func (g *GameController) Cleanup() {
	g.Window.CleanUp()
}

func (g *GameController) Run() error {
	g.Setup()
	defer g.Cleanup()

	if err := g.Window.Show(); err != nil {
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

		if err := g.Game.Update(); err != nil {
			return err
		}

		if err := g.Window.DrawScreen(); err != nil {
			return err
		}
	}

	return nil
}
