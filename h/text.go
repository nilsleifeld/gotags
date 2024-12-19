package h

type text struct {
	value string
}

func (t *text) Render(c *HTMLContext) {
	c.Builder.WriteString(t.value)
}

func Text(value string) HTML {
	return &text{value: value}
}
