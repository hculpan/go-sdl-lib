package component

type BasePage struct {
	BaseComponent

	Name string
}

func (p BasePage) GetName() string {
	return p.Name
}

func (p *BasePage) PageLoad() error {
	return nil
}

func (p *BasePage) PageUnload() error {
	return nil
}

func (p *BasePage) Quit() bool {
	return true
}
