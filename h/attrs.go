package h

import "strings"

type Attribute struct {
	Name  string
	Value string
}

func Attr(name, value string) *Attribute {
	return &Attribute{
		Name:  name,
		Value: value,
	}
}

func (a *Attribute) Render(builder *strings.Builder) {
	builder.WriteString(a.Name)
	if a.Value != "" {
		builder.WriteString("=\"")
		builder.WriteString(a.Value)
		builder.WriteString("\"")
	}
}

// Attrs is a helper function to create a slice of attributes.
func Attrs(attrs ...*Attribute) []*Attribute {
	return attrs
}

func Id(value string) *Attribute {
	return &Attribute{
		Name:  "id",
		Value: value,
	}
}

func Class(value string) *Attribute {
	return &Attribute{
		Name:  "class",
		Value: value,
	}
}

func Disabled() *Attribute {
	return &Attribute{
		Name:  "disabled",
		Value: "",
	}
}
