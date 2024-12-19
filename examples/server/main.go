package main

import (
	"io"
	"net/http"

	. "github.com/nilsleifeld/gotags/h"
)

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func page() HTML {
	todos := []string{"Buy milk", "Walk the dog", "Learn Go"}

	return Html().Lang("de").Child(
		Head().Child(Meta().Charset("utf-8")),
		Body().Child(
			Main().
				Child(If(true, H1(Text("Todo")))).
				Children(ForEach(todos, func(todo string) HTML {
					return Div().Child(Text(todo))
				})).
				Children(ForEach(todos, func(todo string) HTML {
					return Div().Child(Text(todo))
				})),
		),
	)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	html := Render(page())
	io.WriteString(w, html)
}
