package main

import (
	"fmt"
	. "github.com/nilsleifeld/gotags/h"
)

func main() {
	fmt.Println(Render(button()))
}

func button() HTML {
	isPrimary := true
	return Button_().
		If(isPrimary, func(t *buttonBuilder) {
			t.Variant(ButtonVariantPrimary)
		}).
		Child(Text("Hello World"))
}

type buttonVariant string

const (
	ButtonVariantDefault buttonVariant = "default"
	ButtonVariantPrimary buttonVariant = "primary"
	ButtonVariantSuccess buttonVariant = "success"
)

func Button_(
	children ...HTML,
) *buttonBuilder {
	return newButtonBuilder(children...)
}

type buttonBuilder struct {
	TagBuilder[buttonBuilder]
}

func newButtonBuilder(
	children ...HTML,
) *buttonBuilder {
	return &buttonBuilder{
		TagBuilder: NewTagBuilder[buttonBuilder](
			"sl-button",
			children...,
		),
	}
}

func (b *buttonBuilder) Variant(variant buttonVariant) *buttonBuilder {
	b.Attr("variant", string(variant))
	return b
}
