package component

import (
	"fmt"
	"sync"
)

type Page interface {
	Component

	GetName() string
	PageLoad() error
	PageUnload() error
	SetData(data interface{})
	GetData() interface{}

	Quit() bool
}

var pages = make(map[string]Page)

var ActivePage Page
var ActivePageWaitGroup sync.WaitGroup

func RegisterPage(p Page) {
	pages[p.GetName()] = p
	if ActivePage == nil {
		ActivePage = p
	}
}

func GetPage(name string) Page {
	return pages[name]
}

func SwitchPageWithData(newPage string, data interface{}) {
	SwitchPage(newPage)
	ActivePage.SetData(data)
}

func SwitchPage(newPage string) {
	ActivePageWaitGroup.Wait()
	ActivePageWaitGroup.Add(1)
	defer ActivePageWaitGroup.Done()

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
