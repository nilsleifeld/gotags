package main

import (
	"fmt"

	"github.com/nilsleifeld/gotags/h"
)

func main() {
	fmt.Println(h.Render(page()))
}

func page() *h.Element {
	return h.Html(nil,
		h.Head(nil),
		h.Body(nil,
			h.Main(
				h.Attrs(
					h.Id("app"),
					h.Class("container"),
				),
				h.H1(nil, h.Text("Hello, world!")),
				h.Input(h.Attrs(h.Id("input"), h.Disabled())),
				h.Portal(nil),
			),
		),
	)
}
