package h

type TagBuilder struct {
	Tag *Tag
}

func NewTagBuilder(
	name string,
	children ...HTML,
) *TagBuilder {
	return &TagBuilder{
		Tag: &Tag{
			Name:     name,
			Children: children,
		},
	}
}

func (t *TagBuilder) Build() *Tag {
	return t.Tag
}

func (t *TagBuilder) SetSelfClosing() *TagBuilder {
	t.Tag.SelfClosing = true
	return t
}

func (t *TagBuilder) Render(c *HTMLContext) {
	t.Tag.Render(c)
}

func (t *TagBuilder) Child(children ...HTML) *TagBuilder {
	t.Tag.Children = append(t.Tag.Children, children...)
	return t
}

func (t *TagBuilder) Children(children []HTML) *TagBuilder {
	t.Tag.Children = append(t.Tag.Children, children...)
	return t
}

func (t *TagBuilder) If(condition bool, fun func(*TagBuilder),
) *TagBuilder {
	if condition {
		fun(t)
	}
	return t
}

func (t *TagBuilder) IfElse(condition bool, fun func(*TagBuilder),
	elseFun func(*TagBuilder),
) *TagBuilder {
	if condition {
		fun(t)
	} else {
		elseFun(t)
	}
	return t
}

func (t *TagBuilder) ForEach(values []string, fun func(*TagBuilder, string),
) *TagBuilder {
	for _, value := range values {
		fun(t, value)
	}
	return t
}

func (t *TagBuilder) With(fun ...func(*TagBuilder)) *TagBuilder {
	for _, f := range fun {
		f(t)
	}
	return t
}
