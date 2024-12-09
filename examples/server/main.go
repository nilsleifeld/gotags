package main

import (
	"io"
	"net/http"

	"github.com/nilsleifeld/gotags/h"
)

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
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
			),
		),
	)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	html := h.Render(page())
	io.WriteString(w, html)
}
