package resources

import (
	"embed"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

//go:embed built-in-fonts/*

var builtInFonts embed.FS

type FontsManager struct {
	EmbeddedFonts embed.FS

	fonts map[string]*ttf.Font
}

var Fonts FontsManager

func FontsInit(embeddedFont embed.FS) error {
	if err := ttf.Init(); err != nil {
		return err
	}

	Fonts = FontsManager{EmbeddedFonts: embeddedFont}
	Fonts.fonts = make(map[string]*ttf.Font)

	return nil
}

func (f *FontsManager) RegisterFont(key, filename string, size int) error {
	// Load data: First check app's embedded fonts, then check
	// built-in fonts
	data, err := f.EmbeddedFonts.ReadFile(filename)
	if err != nil {
		data, err = builtInFonts.ReadFile(filename)
		if err != nil {
			return err
		}
	}

	rwops, err := sdl.RWFromMem(data)
	if err != nil {
		return err
	}

	font, err := ttf.OpenFontRW(rwops, 1, size)
	if err != nil {
		return err
	}

	f.fonts[key] = font

	return nil

}

func (f *FontsManager) GetFont(key string) (*ttf.Font, error) {
	if result, exists := f.fonts[key]; exists {
		return result, nil
	} else {
		return nil, fmt.Errorf("font '%s' not found", key)
	}
}

func (f *FontsManager) CreateTexture(msg string, c sdl.Color, fontKey string, renderer *sdl.Renderer) (*sdl.Texture, error) {
	font, err := f.GetFont(fontKey)
	if err != nil {
		return nil, err
	}

	surface, err := font.RenderUTF8Solid(msg, c)
	if err != nil {
		return nil, fmt.Errorf("unable to create texture surface: %v", err)
	}
	defer surface.Free()

	msgtext, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return nil, fmt.Errorf("unable to create texture from surface: %v", err)
	}

	return msgtext, nil
}
