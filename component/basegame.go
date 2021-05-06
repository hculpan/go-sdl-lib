package component

type BaseGame struct {
	GameWidth  int32
	GameHeight int32
}

// Initialize should be called by any descending structs
func (g *BaseGame) Initialize(gameWidth int32, gameHeight int32) {
	g.GameWidth = gameWidth
	g.GameHeight = gameHeight
}

func (g *BaseGame) LoadPages() []Page {
	return []Page{}
}

func (g *BaseGame) RegisterPages(pages []Page) error {
	for _, page := range pages {
		RegisterPage(page)
	}

	return nil
}
