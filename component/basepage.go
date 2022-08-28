package component

type BasePage struct {
	BaseComponent

	Data interface{}
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

func (p *BasePage) SetData(data interface{}) {
	p.Data = data
}

func (p *BasePage) GetData() interface{} {
	return p.Data
}
