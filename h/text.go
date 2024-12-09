package h

import "strings"

type text struct {
	value string
}

func (t *text) Render(builder *strings.Builder) {
	builder.WriteString(t.value)
}

func Text(value string) Renderable {
	return &text{value: value}
}
