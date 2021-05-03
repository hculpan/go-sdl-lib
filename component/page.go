package component

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Page interface {
	Component
	Name() string

	DrawPage(renderer *sdl.Renderer) error
}

var ActivePage Page

var pages = make(map[string]Page)

func RegisterPage(p Page) {
	pages[p.Name()] = p
	if ActivePage == nil {
		ActivePage = p
	}
}

func GetPage(name string) Page {
	return pages[name]
}

func SwitchPage(newPage string) {
	p := GetPage(newPage)
	if p == nil {
		panic(fmt.Sprintf("No page defined for '%s'", newPage))
	}
}
