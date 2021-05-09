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
	_, exists := f.fonts[key]
	if exists {
		return nil
	}

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

	// This is a kludge.  I find I have fewer issues if I
	// preload a texture of the font before actually using it.
	// No idea why this is needed, but whatever.
	if err := f.preloadFont(key); err != nil {
		panic(err)
	}

	return nil
}

func (f *FontsManager) GetKey(fontname string, size int) string {
	return fmt.Sprintf("%s-%d", fontname, size)
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

func (f *FontsManager) preloadFont(key string) error {
	window, err := sdl.CreateWindow(
		"fonts",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		100,
		100,
		sdl.WINDOW_OPENGL|sdl.WINDOW_HIDDEN,
	)
	if err != nil {
		fmt.Println("Initializing window for FontManager: ", err)
		return err
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Initializing renderer: ", err)
		return err
	}
	defer renderer.Destroy()

	f.CreateTexture("ABCDEFGGUHJKNBIOQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789`~!@#$%^&*()_+|=\\[]{};':\",.<>/?", sdl.Color{R: 50, G: 255, B: 50, A: 255}, key, renderer)

	return nil
}
