package component

import (
	"github.com/creasty/defaults"
	"github.com/hculpan/go-sdl-lib/resources"
	"github.com/veandco/go-sdl2/sdl"
)

type LabelConfig struct {
	FontName string    `default:"TruenoLight"`
	FontFile string    `default:"built-in-fonts/TruenoLight.otf"`
	FontSize int       `default:"24"`
	Justify  int       `default:"0" `
	Color    sdl.Color `default:"{\"R\": 50, \"G\": 255, \"B\": 50, \"A\": 255}"`
}

type LabelComponent struct {
	BaseComponent

	textFunc      LabelTextFunc
	fontKey       string
	justification int
	color         sdl.Color
}

const (
	JUSTIFY_LEFT = iota
	JUSTIFY_CENTER
	JUSTIFY_RIGHT
)

type LabelTextFunc func() string

func NewLabelComponent(x, y, width, height int32, fontSize int, labelTextFunc LabelTextFunc) *LabelComponent {
	labelConfig := NewLabelConfig()
	labelConfig.FontSize = fontSize
	return NewLabelComponentWithConfig(x, y, width, height, labelTextFunc, labelConfig)
	/*
	   result.SetPosition(x, y)
	   result.SetSize(width, height)
	   result.textFunc = labelTextFunc
	   result.fontKey = resources.Fonts.GetKey("TruenoLight", fontSize)
	   result.fontSize = fontSize
	   result.justification = JUSTIFY_LEFT

	   resources.Fonts.RegisterFont(result.fontKey, "built-in-fonts/TruenoLight.otf", fontSize)

	   return result
	*/
}

func NewLabelComponentWithConfig(x, y, width, height int32, labelTextFunc LabelTextFunc, labelConfig LabelConfig) *LabelComponent {
	result := basicLabelComponent(x, y, width, height, labelTextFunc)
	fontKey, err := registerFont(labelConfig.FontName, labelConfig.FontSize, labelConfig.FontFile)
	if err != nil {
		panic(err)
	}
	result.fontKey = fontKey
	result.justification = labelConfig.Justify
	result.color = labelConfig.Color

	return result
}

func NewLabelConfig() LabelConfig {
	result := &LabelConfig{}
	if err := defaults.Set(result); err != nil {
		panic(err)
	}
	return *result
}

func basicLabelComponent(x, y, width, height int32, labelTextFunc LabelTextFunc) *LabelComponent {
	result := &LabelComponent{}

	result.SetPosition(x, y)
	result.SetSize(width, height)
	result.textFunc = labelTextFunc

	return result
}

func registerFont(fontName string, fontSize int, fontFile string) (string, error) {
	fontKey := resources.Fonts.GetKey(fontName, fontSize)
	if err := resources.Fonts.RegisterFont(fontKey, fontFile, fontSize); err != nil {
		return fontKey, err
	}

	return fontKey, nil
}

func (c *LabelComponent) DrawComponent(r *sdl.Renderer) error {
	var msg string = "no text"
	if c.textFunc != nil {
		msg = c.textFunc()
	}

	text, err := resources.Fonts.CreateTexture(msg, c.color, c.fontKey, r)
	if err != nil {
		return err
	}
	defer text.Destroy()
	_, _, w, h, err := text.Query()
	if err != nil {
		return err
	}

	var destRect sdl.Rect
	switch c.justification {
	case JUSTIFY_LEFT:
		destRect = sdl.Rect{X: c.X + 5, Y: c.Y, W: int32(w), H: int32(h)}
	case JUSTIFY_CENTER:
		destRect = sdl.Rect{X: (c.Width / 2) - (w / 2), Y: c.Y, W: int32(w), H: int32(h)}
	case JUSTIFY_RIGHT:
		destRect = sdl.Rect{X: c.Width - (w + 5), Y: c.Y, W: int32(w), H: int32(h)}
	}
	r.Copy(text, &sdl.Rect{X: 0, Y: 0, W: w, H: h}, &destRect)

	return nil
}

func (c *LabelComponent) Draw(r *sdl.Renderer) error {
	if err := DrawParentAndChildren(r, c); err != nil {
		return err
	}

	return nil
}
