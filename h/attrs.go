package h

func (t *TagBuilder) Attr(name, value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute(name, value))
	return t
}

func (t *TagBuilder) Lang(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("lang", value))
	return t
}

func (t *TagBuilder) Id(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("id", value))
	return t
}

func (t *TagBuilder) Class(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("class", value))
	return t
}

func (t *TagBuilder) Style(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("style", value))
	return t
}

func (t *TagBuilder) Src(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("src", value))
	return t
}

func (t *TagBuilder) Href(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("href", value))
	return t
}

func (t *TagBuilder) Charset(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("charset", value))
	return t
}

func (t *TagBuilder) Name(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("name", value))
	return t
}

func (t *TagBuilder) Value(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("value", value))
	return t
}

func (t *TagBuilder) Type(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("type", value))
	return t
}

func (t *TagBuilder) Placeholder(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("placeholder", value))
	return t
}

func (t *TagBuilder) Disabled() *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("disabled", ""))
	return t
}

func (t *TagBuilder) Selected() *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("selected", ""))
	return t
}

func (t *TagBuilder) Required() *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("required", ""))
	return t
}

func (t *TagBuilder) Pattern(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("pattern", value))
	return t
}

func (t *TagBuilder) Min(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("min", value))
	return t
}

func (t *TagBuilder) Max(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("max", value))
	return t
}

func (t *TagBuilder) Step(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("step", value))
	return t
}

func (t *TagBuilder) Target(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("target", value))
	return t
}

func (t *TagBuilder) Rel(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("rel", value))
	return t
}

func (t *TagBuilder) Alt(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("alt", value))
	return t
}

func (t *TagBuilder) Title(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("title", value))
	return t
}

func (t *TagBuilder) For(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("for", value))
	return t
}

func (t *TagBuilder) Method(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("method", value))
	return t
}

func (t *TagBuilder) Action(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("action", value))
	return t
}

func (t *TagBuilder) Accept(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("accept", value))
	return t
}

func (t *TagBuilder) Autocomplete(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("autocomplete", value))
	return t
}

func (t *TagBuilder) Autoplay() *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("autoplay", ""))
	return t
}

func (t *TagBuilder) Controls() *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("controls", ""))
	return t
}

func (t *TagBuilder) Loop() *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("loop", ""))
	return t
}

func (t *TagBuilder) Media(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("media", value))
	return t
}

func (t *TagBuilder) Defer() *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("defer", ""))
	return t
}

func (t *TagBuilder) Rowspan(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("rowspan", value))
	return t
}

func (t *TagBuilder) Colspan(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("colspan", value))
	return t
}

func (t *TagBuilder) Headers(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("headers", value))
	return t
}

func (t *TagBuilder) Scope(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("scope", value))
	return t
}

func (t *TagBuilder) Async() *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("async", ""))
	return t
}

func (t *TagBuilder) Integrity(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("integrity", value))
	return t
}

func (t *TagBuilder) Crossorigin(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("crossorigin", value))
	return t
}

func (t *TagBuilder) Download() *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("download", ""))
	return t
}

func (t *TagBuilder) Datalist(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("datalist", value))
	return t
}

func (t *TagBuilder) Enctype(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("enctype", value))
	return t
}

func (t *TagBuilder) Novalidate() *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("novalidate", ""))
	return t
}

func (t *TagBuilder) Formaction(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("formaction", value))
	return t
}

func (t *TagBuilder) Formmethod(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("formmethod", value))
	return t
}

func (t *TagBuilder) Formnovalidate() *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("formnovalidate", ""))
	return t
}

func (t *TagBuilder) Formtarget(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("formtarget", value))
	return t
}

func (t *TagBuilder) Ping(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("ping", value))
	return t
}

func (t *TagBuilder) Inputmode(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("inputmode", value))
	return t
}

func (t *TagBuilder) Content(value string) *TagBuilder {
	t.Tag.Attrs = append(t.Tag.Attrs, NewAttribute("content", value))
	return t
}
