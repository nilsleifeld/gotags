package h

import "strings"

type Renderable interface {
	Render(builder *strings.Builder)
}

type Element struct {
	Tag         string
	Attrs       []*Attribute
	Children    []Renderable
	SelfClosing bool
	Void        bool
}

func NewSelfClosingElement(tag string, attrs []*Attribute) *Element {
	return &Element{
		Tag:         tag,
		Attrs:       attrs,
		SelfClosing: true,
	}
}

func (e *Element) Render(builder *strings.Builder) {
	builder.WriteString("<")
	builder.WriteString(e.Tag)

	for _, attr := range e.Attrs {
		builder.WriteString(" ")
		attr.Render(builder)
	}
	if e.SelfClosing {
		builder.WriteString(" />")
		return
	}

	builder.WriteString(">")

	if e.Void {
		return
	}

	for _, child := range e.Children {
		child.Render(builder)
	}

	builder.WriteString("</")
	builder.WriteString(e.Tag)
	builder.WriteString(">")
}
