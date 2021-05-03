package static

import (
	"embed"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

//go:embed fonts/*

var Fonts embed.FS

func InitializeFont(fileName string) (*ttf.Font, error) {
	data, err := Fonts.ReadFile("fonts/" + fileName)
	if err != nil {
		return nil, err
	}

	rwops, err := sdl.RWFromMem(data)
	if err != nil {
		return nil, err
	}

	font, err := ttf.OpenFontRW(rwops, 1, 18)
	if err != nil {
		return nil, err
	}

	return font, nil
}
