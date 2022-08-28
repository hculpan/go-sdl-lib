package component

import (
	"github.com/hculpan/go-sdl-lib/resources"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	buttonUp = iota
	buttonDown
)

type ButtonComponent struct {
	BaseComponent

	Text     string
	OnAction OnActionFunc

	normalTexture   *sdl.Texture
	pressedTexture  *sdl.Texture
	backgroundColor sdl.Color
	buttonColor     sdl.Color
	textColor       sdl.Color

	buttonState int
}

func NewButtonComponent(x, y, w, h int32, text string, backgroundColor sdl.Color, buttonColor sdl.Color, textColor sdl.Color, onAction OnActionFunc) *ButtonComponent {
	result := ButtonComponent{buttonState: buttonUp}
	result.Initialize()
	result.SetPosition(x, y)
	result.SetSize(w, h)
	result.backgroundColor = backgroundColor
	result.buttonColor = buttonColor
	result.textColor = textColor
	result.Text = text
	result.OnAction = onAction

	resources.Fonts.RegisterFont("Trueno-16", "built-in-fonts/Trueno.otf", 32)

	return &result
}

func (c *ButtonComponent) initializePressedTexture(r *sdl.Renderer) error {
	var err error
	c.pressedTexture, err = resources.CreateBlankTexture(r, c.Width, c.Height)
	if err != nil {
		return err
	}

	target := r.GetRenderTarget()

	r.SetRenderTarget(c.pressedTexture)
	r.SetDrawColor(c.backgroundColor.R, c.backgroundColor.G, c.backgroundColor.B, c.backgroundColor.A)
	r.FillRect(&sdl.Rect{X: 0, Y: 0, W: c.Width, H: c.Height})
	gfx.RoundedBoxRGBA(r, 2, 2, c.Width-2, c.Height-2, 4, c.textColor.R, c.textColor.G, c.textColor.B, c.textColor.A)

	text, err := resources.Fonts.CreateTexture(c.Text, sdl.Color{R: c.buttonColor.R, G: c.buttonColor.G, B: c.buttonColor.B, A: c.buttonColor.A}, "Trueno-16", r)
	if err != nil {
		return err
	}
	defer text.Destroy()

	_, _, tw, th, err := text.Query()
	if err != nil {
		return err
	}
	r.Copy(text,
		&sdl.Rect{X: 0, Y: 0, W: tw, H: th},
		&sdl.Rect{X: (c.Width / 2) - (tw / 2), Y: (c.Height / 2) - (th / 2), W: tw, H: th},
	)

	r.SetRenderTarget(target)

	return nil
}

func (c *ButtonComponent) initializeNormalTexture(r *sdl.Renderer) error {
	var err error
	c.normalTexture, err = resources.CreateBlankTexture(r, c.Width, c.Height)
	if err != nil {
		return err
	}

	target := r.GetRenderTarget()

	r.SetRenderTarget(c.normalTexture)
	r.SetDrawColor(c.backgroundColor.R, c.backgroundColor.G, c.backgroundColor.B, c.backgroundColor.A)
	r.FillRect(&sdl.Rect{X: 0, Y: 0, W: c.Width, H: c.Height})
	gfx.RoundedBoxRGBA(r, 2, 2, c.Width-2, c.Height-2, 4, c.buttonColor.R, c.buttonColor.G, c.buttonColor.B, c.buttonColor.A)

	text, err := resources.Fonts.CreateTexture(c.Text, sdl.Color{R: c.textColor.R, G: c.textColor.G, B: c.textColor.B, A: c.textColor.A}, "Trueno-16", r)
	if err != nil {
		return err
	}
	defer text.Destroy()
	_, _, tw, th, err := text.Query()
	if err != nil {
		return err
	}
	r.Copy(text,
		&sdl.Rect{X: 0, Y: 0, W: tw, H: th},
		&sdl.Rect{X: (c.Width / 2) - (tw / 2), Y: (c.Height / 2) - (th / 2), W: tw, H: th},
	)

	r.SetRenderTarget(target)

	return nil
}

func (c *ButtonComponent) DrawComponent(r *sdl.Renderer) error {
	if c.normalTexture == nil {
		c.initializeNormalTexture(r)
		c.initializePressedTexture(r)
	}

	if c.buttonState == buttonUp {
		r.Copy(c.normalTexture,
			&sdl.Rect{X: 0, Y: 0, W: c.Width, H: c.Height},
			&sdl.Rect{X: c.X, Y: c.Y, W: c.Width, H: c.Height},
		)
	} else {
		r.Copy(c.pressedTexture,
			&sdl.Rect{X: 0, Y: 0, W: c.Width, H: c.Height},
			&sdl.Rect{X: c.X, Y: c.Y, W: c.Width, H: c.Height},
		)
	}

	return nil
}

func (c *ButtonComponent) Draw(r *sdl.Renderer) error {
	if err := DrawParentAndChildren(r, c); err != nil {
		return err
	}

	return nil
}

func (c *ButtonComponent) KeyEvent(event *sdl.KeyboardEvent) bool {
	return PassKeyEventToChildren(event, c.Children)
}

func (c *ButtonComponent) MouseButtonEvent(event *sdl.MouseButtonEvent) bool {
	if event.Type == sdl.MOUSEBUTTONDOWN {
		c.buttonState = buttonDown
	} else {
		c.buttonState = buttonUp
		if c.OnAction != nil {
			c.OnAction()
		}
	}

	return true
}
