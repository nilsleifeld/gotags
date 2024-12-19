package h

type Tag struct {
	Name        string
	Attrs       []*Attribute
	Children    []HTML
	SelfClosing bool
}

func (t *Tag) Render(c *HTMLContext) {
	c.Builder.WriteString("<")
	c.Builder.WriteString(t.Name)

	class := ""
	for _, attr := range t.Attrs {
		if attr == nil {
			continue
		}
		if attr.Name == "class" {
			class += attr.Value
			class += " "
			continue
		}
		c.Builder.WriteString(" ")
		attr.Render(c)
	}

	if class != "" {
		c.Builder.WriteString(" class=\"")
		c.Builder.WriteString(class)
		c.Builder.WriteString("\" ")
	}

	if t.SelfClosing {
		c.Builder.WriteString(" />")
		return
	}

	c.Builder.WriteString(">")

	for _, child := range t.Children {
		if child != nil {
			child.Render(c)
		}
	}

	c.Builder.WriteString("</")
	c.Builder.WriteString(t.Name)
	c.Builder.WriteString(">")
}
