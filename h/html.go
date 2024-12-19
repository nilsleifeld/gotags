package h

import "strings"

type HTMLContext struct {
	Builder strings.Builder
}

type HTML interface {
	Render(c *HTMLContext)
}
