package game

import (
	"math"

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
	var start, end uint64
	for running {
		start = sdl.GetPerformanceCounter()
		event := sdl.PollEvent()
		if event != nil {
			switch event := event.(type) {
			case *sdl.QuitEvent:
				running = !component.ActivePage.Quit()
			case *sdl.KeyboardEvent:
				keycode := sdl.GetKeyFromScancode(event.Keysym.Scancode)
				if event.Type == sdl.KEYDOWN {
					if keycode == sdl.K_ESCAPE {
						running = !component.ActivePage.Quit()
					} else {
						component.ActivePage.KeyEvent(event)
					}
				}
			case *sdl.MouseButtonEvent:
				component.ActivePage.BaseMouseButtonEvent(event)
			}
		}

		if err := g.Game.Update(); err != nil {
			return err
		}

		if err := g.Window.DrawScreen(); err != nil {
			return err
		}

		end = sdl.GetPerformanceCounter()
		elapsed := (float64(end-start) / float64(sdl.GetPerformanceFrequency()))

		elapsedMS := elapsed * 1000.0
		delay := uint32(math.Floor(16.666 - elapsedMS))

		// Cap to 60 FPS
		if delay > 0 && delay < 17 {
			sdl.Delay(delay)
		}
	}

	return nil
}
