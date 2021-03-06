package component

import (
	"fmt"
)

type Page interface {
	Component

	GetName() string
	PageLoad() error
	PageUnload() error

	Quit() bool
}

var pages = make(map[string]Page)

var ActivePage Page

func RegisterPage(p Page) {
	pages[p.GetName()] = p
	if ActivePage == nil {
		ActivePage = p
	}
}

func GetPage(name string) Page {
	return pages[name]
}

func SwitchPage(newPage string) {
	if ActivePage != nil {
		if err := ActivePage.PageUnload(); err != nil {
			panic(err)
		}
	}

	p := GetPage(newPage)
	if p == nil {
		panic(fmt.Sprintf("No page defined for '%s'", newPage))
	} else if err := p.PageLoad(); err != nil {
		panic(err)
	}

	ActivePage = p
}
