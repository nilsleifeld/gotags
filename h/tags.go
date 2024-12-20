package h

func Html(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("html", children...)
}

func Head(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("head", children...)
}

func Body(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("body", children...)
}

func Main(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("main", children...)
}

func Input() *DefaultTagBuilder {
	return NewDefaultTagBuilder("main")
}

func Portal() *DefaultTagBuilder {
	return NewDefaultTagBuilder("main")
}

func Meta() *DefaultTagBuilder {
	return NewDefaultTagBuilder("meta").
		SetSelfClosing()
}

func Button(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("button", children...)
}

func Div(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("div", children...)
}

func Base() *DefaultTagBuilder {
	return NewDefaultTagBuilder("base").
		SetSelfClosing()
}

func Link() *DefaultTagBuilder {
	return NewDefaultTagBuilder("link").
		SetSelfClosing()
}

func Style(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("style", children...)
}

func Title(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("title", children...)
}

func A(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("a", children...)
}

func Address(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("address", children...)
}

func Article(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("article", children...)
}

func Aside(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("aside", children...)
}

func Footer(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("footer", children...)
}

func Header(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("header", children...)
}

func H1(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("h1", children...)
}

func H2(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("h2", children...)
}

func H3(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("h3", children...)
}

func H4(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("h4", children...)
}

func H5(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("h5", children...)
}

func H6(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("h6", children...)
}

func HGroup(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("hgroup", children...)
}

func Nav(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("nav", children...)
}

func Section(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("section", children...)
}

func Search(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("search", children...)
}

func Blockquote(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("blockquote", children...)
}

func Dd(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("dd", children...)
}

func Dl(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("dl", children...)
}

func Dt(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("dt", children...)
}

func Figcaption(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("figcaption", children...)
}

func Figure(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("figure", children...)
}

func Hr() *DefaultTagBuilder {
	return NewDefaultTagBuilder("hr").
		SetSelfClosing()
}

func Li(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("li", children...)
}

func Menu(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("menu", children...)
}

func OL(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("ol", children...)
}

func P(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("p", children...)
}

func Pre(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("pre", children...)
}

func Ul(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("ul", children...)
}

func Abbr(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("abbr", children...)
}

func B(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("b", children...)
}

func Bdi(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("bdi", children...)
}

func Bdo(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("bdo", children...)
}

func Br() *DefaultTagBuilder {
	return NewDefaultTagBuilder("br").
		SetSelfClosing()
}

func Cite(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("cite", children...)
}

func Code(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("code", children...)
}

func Data(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("data", children...)
}

func Dfn(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("dfn", children...)
}

func Em(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("em", children...)
}

func I(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("i", children...)
}

func Kbd(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("kbd", children...)
}

func Mark(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("mark", children...)
}

func Q(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("q", children...)
}

func Rb(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("rb", children...)
}

func Rp(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("rp", children...)
}

func Rt(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("rt", children...)
}

func Ruby(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("ruby", children...)
}

func S(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("s", children...)
}

func Samp(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("samp", children...)
}

func Small(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("small", children...)
}

func Span(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("span", children...)
}

func Strong(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("strong", children...)
}

func Sub(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("sub", children...)
}

func Time(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("time", children...)
}

func U(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("u", children...)
}

func Var(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("var", children...)
}

func Wbr() *DefaultTagBuilder {
	return NewDefaultTagBuilder("wbr").
		SetSelfClosing()
}

func Area() *DefaultTagBuilder {
	return NewDefaultTagBuilder("area").
		SetSelfClosing()
}

func Audio(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("audio", children...)
}

func Img() *DefaultTagBuilder {
	return NewDefaultTagBuilder("img").
		SetSelfClosing()
}

func Map(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("map", children...)
}

func Track() *DefaultTagBuilder {
	return NewDefaultTagBuilder("track").
		SetSelfClosing()
}

func Video(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("video", children...)
}

func Embed() *DefaultTagBuilder {
	return NewDefaultTagBuilder("embed").
		SetSelfClosing()
}

func Iframe() *DefaultTagBuilder {
	return NewDefaultTagBuilder("iframe").
		SetSelfClosing()
}

func Object() *DefaultTagBuilder {
	return NewDefaultTagBuilder("object").
		SetSelfClosing()
}

func Picture(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("picture", children...)
}

func Source() *DefaultTagBuilder {
	return NewDefaultTagBuilder("source").
		SetSelfClosing()
}

func Svg(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("svg", children...)
}

func Math(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("math", children...)
}

func Canvas() *DefaultTagBuilder {
	return NewDefaultTagBuilder("canvas")
}

func Noscript(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("noscript", children...)
}

func Script(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("script", children...)
}

func Del(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("del", children...)
}

func Ins(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("ins", children...)
}

func Caption(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("caption", children...)
}

func Col() *DefaultTagBuilder {
	return NewDefaultTagBuilder("col").
		SetSelfClosing()
}

func Colgroup(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("colgroup", children...)
}

func Table(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("table", children...)
}

func Tbody(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("tbody", children...)
}

func Td(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("td", children...)
}

func Tfoot(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("tfoot", children...)
}

func Th(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("th", children...)
}

func Thead(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("thead", children...)
}

func Tr(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("tr", children...)
}

func Datalist(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("datalist", children...)
}

func Fieldset(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("fieldset", children...)
}

func Form(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("form", children...)
}

func Label(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("label", children...)
}

func Legend(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("legend", children...)
}

func Meter(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("meter", children...)
}

func Optgroup(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("optgroup", children...)
}

func Option(label string) *DefaultTagBuilder {
	return NewDefaultTagBuilder("option", Text(label))
}

func Output(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("output", children...)
}

func Progress(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("progress", children...)
}

func Select(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("select", children...)
}

func Textarea(content string) *DefaultTagBuilder {
	return NewDefaultTagBuilder("textarea", Text(content))
}

func Details(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("details", children...)
}

func Dialog(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("dialog", children...)
}

func Summary(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("summary", children...)
}

func Slot() *DefaultTagBuilder {
	return NewDefaultTagBuilder("slot")
}

func Template(children ...HTML) *DefaultTagBuilder {
	return NewDefaultTagBuilder("template", children...)
}
