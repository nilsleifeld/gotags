package h

func (t *TagBuilder[T]) Attr(name, value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute(name, value))
	return t.AsT()
}

func (t *TagBuilder[T]) Lang(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("lang", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Id(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("id", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Class(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("class", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Style(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("style", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Src(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("src", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Href(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("href", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Charset(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("charset", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Name(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("name", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Value(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("value", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Type(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("type", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Placeholder(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("placeholder", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Disabled() *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("disabled", ""))
	return t.AsT()
}

func (t *TagBuilder[T]) Selected() *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("selected", ""))
	return t.AsT()
}

func (t *TagBuilder[T]) Required() *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("required", ""))
	return t.AsT()
}

func (t *TagBuilder[T]) Pattern(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("pattern", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Min(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("min", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Max(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("max", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Step(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("step", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Target(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("target", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Rel(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("rel", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Alt(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("alt", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Title(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("title", value))
	return t.AsT()
}

func (t *TagBuilder[T]) For(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("for", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Method(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("method", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Action(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("action", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Accept(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("accept", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Autocomplete(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("autocomplete", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Autoplay() *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("autoplay", ""))
	return t.AsT()
}

func (t *TagBuilder[T]) Controls() *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("controls", ""))
	return t.AsT()
}

func (t *TagBuilder[T]) Loop() *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("loop", ""))
	return t.AsT()
}

func (t *TagBuilder[T]) Media(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("media", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Defer() *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("defer", ""))
	return t.AsT()
}

func (t *TagBuilder[T]) Rowspan(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("rowspan", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Colspan(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("colspan", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Headers(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("headers", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Scope(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("scope", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Async() *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("async", ""))
	return t.AsT()
}

func (t *TagBuilder[T]) Integrity(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("integrity", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Crossorigin(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("crossorigin", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Download() *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("download", ""))
	return t.AsT()
}

func (t *TagBuilder[T]) Datalist(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("datalist", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Enctype(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("enctype", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Novalidate() *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("novalidate", ""))
	return t.AsT()
}

func (t *TagBuilder[T]) Formaction(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("formaction", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Formmethod(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("formmethod", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Formnovalidate() *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("formnovalidate", ""))
	return t.AsT()
}

func (t *TagBuilder[T]) Formtarget(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("formtarget", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Ping(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("ping", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Inputmode(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("inputmode", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Content(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("content", value))
	return t.AsT()
}

func (t *TagBuilder[T]) Slot(value string) *T {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("slot", value))
	return t.AsT()
}
