package h

type Attribute struct {
	Name  string
	Value string
}

func NewAttribute(name, value string) *Attribute {
	return &Attribute{
		Name:  name,
		Value: value,
	}
}

func (a *Attribute) Render(c *HTMLContext) {
	c.Builder.WriteString(a.Name)
	if a.Value != "" {
		c.Builder.WriteString("=\"")
		c.Builder.WriteString(a.Value)
		c.Builder.WriteString("\"")
	}
}
