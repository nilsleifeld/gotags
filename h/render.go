package h

import (
	"net/http"
	"strings"
)

func Render(e HTML) string {
	c := &HTMLContext{
		Builder: strings.Builder{},
	}
	c.Builder.WriteString("<!DOCTYPE html>")

	e.Render(c)

	return c.Builder.String()
}

func Response(w http.ResponseWriter, e HTML) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(Render(e)))
}
