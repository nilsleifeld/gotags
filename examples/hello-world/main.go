package main

import (
	"fmt"

	. "github.com/nilsleifeld/gotags/h"
)

func main() {
	fmt.Println(Render(page()))
}

func page() HTML {
	return Html().Lang("de").Child(
		Head().Child(Meta().Charset("utf-8")),
		Body().Child(
			Main().Child(Text("Hello World")),
		),
	)
}
