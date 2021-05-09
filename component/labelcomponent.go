package component

import (
	"github.com/hculpan/go-sdl-lib/resources"
	"github.com/veandco/go-sdl2/sdl"
)

type LabelComponent struct {
	BaseComponent

	textFunc LabelTextFunc
	fontKey  string
	fontSize int
}

type LabelTextFunc func() string

func NewLabelComponent(x, y, width, height int32, fontSize int, labelTextFunc LabelTextFunc) *LabelComponent {
	result := &LabelComponent{}

	result.SetPosition(x, y)
	result.SetSize(width, height)
	result.textFunc = labelTextFunc
	result.fontKey = resources.Fonts.GetKey("TruenoLight", fontSize)
	result.fontSize = fontSize

	resources.Fonts.RegisterFont(result.fontKey, "built-in-fonts/TruenoLight.otf", fontSize)

	return result
}

func (c *LabelComponent) DrawComponent(r *sdl.Renderer) error {
	var msg string = "no text"
	if c.textFunc != nil {
		msg = c.textFunc()
	}
	text, err := resources.Fonts.CreateTexture(msg, sdl.Color{R: 50, G: 255, B: 50, A: 255}, "TruenoLight-24", r)
	if err != nil {
		return err
	}
	defer text.Destroy()
	_, _, w, h, err := text.Query()
	if err != nil {
		return err
	}
	r.Copy(text, &sdl.Rect{X: 0, Y: 0, W: w, H: h}, &sdl.Rect{X: c.X + 5, Y: c.Y, W: int32(w), H: int32(h)})

	return nil
}

func (c *LabelComponent) Draw(r *sdl.Renderer) error {
	if err := DrawParentAndChildren(r, c); err != nil {
		return err
	}

	return nil
}
