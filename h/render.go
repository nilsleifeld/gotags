package h

import "strings"

func Render(e Renderable) string {
	builder := strings.Builder{}
	builder.WriteString("<!DOCTYPE html>")

	e.Render(&builder)

	return builder.String()
}
