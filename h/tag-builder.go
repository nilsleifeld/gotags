package h

import "unsafe"

type TagBuilder[T any] struct {
	Tag *Tag
}

type DefaultTagBuilder struct {
	TagBuilder[DefaultTagBuilder]
}

func NewDefaultTagBuilder(
	name string,
	children ...HTML,
) *DefaultTagBuilder {
	return &DefaultTagBuilder{
		TagBuilder: NewTagBuilder[DefaultTagBuilder](
			name,
			children...,
		),
	}
}

func NewTagBuilder[T any](
	name string,
	children ...HTML,
) TagBuilder[T] {
	return TagBuilder[T]{
		Tag: &Tag{
			Name:     name,
			Children: children,
		},
	}
}

func (t *TagBuilder[T]) Build() *Tag {
	return t.Tag
}

func (t *TagBuilder[T]) SetSelfClosing() *T {
	t.Tag.SelfClosing = true
	return t.AsT()
}

func (t *TagBuilder[T]) Render(c *HTMLContext) {
	t.Tag.Render(c)
}

func (t *TagBuilder[T]) Child(children ...HTML) *T {
	t.Tag.Children = append(t.Tag.Children, children...)
	return t.AsT()
}

func (t *TagBuilder[T]) Children(children []HTML) *T {
	t.Tag.Children = append(t.Tag.Children, children...)
	return t.AsT()
}

func (t *TagBuilder[T]) If(condition bool, fun func(*T)) *T {
	o := t.AsT()
	if condition {
		fun(o)
	}
	return o
}

func (t *TagBuilder[T]) AsT() *T {
	return (*T)(unsafe.Pointer(t))
}

func (t *TagBuilder[T]) IfElse(condition bool, fun func(*T),
	elseFun func(*T),
) *T {
	o := t.AsT()
	if condition {
		fun(o)
	} else {
		elseFun(o)
	}
	return o
}

func (t *TagBuilder[T]) ForEach(values []string, fun func(*T, string),
) *T {
	o := t.AsT()
	for _, value := range values {
		fun(o, value)
	}
	return o
}

func (t *TagBuilder[T]) With(fun ...func(*T)) *T {
	o := t.AsT()
	for _, f := range fun {
		f(o)
	}
	return o
}
