package h

func Html(children ...HTML) *TagBuilder {
	return NewTagBuilder("html", children...)
}

func Head(children ...HTML) *TagBuilder {
	return NewTagBuilder("head", children...)
}

func Body(children ...HTML) *TagBuilder {
	return NewTagBuilder("body", children...)
}

func Main(children ...HTML) *TagBuilder {
	return NewTagBuilder("main", children...)
}

func Input() *TagBuilder {
	return NewTagBuilder("main")
}

func Portal() *TagBuilder {
	return NewTagBuilder("main")
}

func Meta() *TagBuilder {
	return NewTagBuilder("meta").
		SetSelfClosing()
}

func Button(children ...HTML) *TagBuilder {
	return NewTagBuilder("button", children...)
}

func Div(children ...HTML) *TagBuilder {
	return NewTagBuilder("div", children...)
}

func Base() *TagBuilder {
	return NewTagBuilder("base").
		SetSelfClosing()
}

func Link() *TagBuilder {
	return NewTagBuilder("link").
		SetSelfClosing()
}

func Style(children ...HTML) *TagBuilder {
	return NewTagBuilder("style", children...)
}

func Title(children ...HTML) *TagBuilder {
	return NewTagBuilder("title", children...)
}

func A(children ...HTML) *TagBuilder {
	return NewTagBuilder("a", children...)
}

func Address(children ...HTML) *TagBuilder {
	return NewTagBuilder("address", children...)
}

func Article(children ...HTML) *TagBuilder {
	return NewTagBuilder("article", children...)
}

func Aside(children ...HTML) *TagBuilder {
	return NewTagBuilder("aside", children...)
}

func Footer(children ...HTML) *TagBuilder {
	return NewTagBuilder("footer", children...)
}

func Header(children ...HTML) *TagBuilder {
	return NewTagBuilder("header", children...)
}

func H1(children ...HTML) *TagBuilder {
	return NewTagBuilder("h1", children...)
}

func H2(children ...HTML) *TagBuilder {
	return NewTagBuilder("h2", children...)
}

func H3(children ...HTML) *TagBuilder {
	return NewTagBuilder("h3", children...)
}

func H4(children ...HTML) *TagBuilder {
	return NewTagBuilder("h4", children...)
}

func H5(children ...HTML) *TagBuilder {
	return NewTagBuilder("h5", children...)
}

func H6(children ...HTML) *TagBuilder {
	return NewTagBuilder("h6", children...)
}

func HGroup(children ...HTML) *TagBuilder {
	return NewTagBuilder("hgroup", children...)
}

func Nav(children ...HTML) *TagBuilder {
	return NewTagBuilder("nav", children...)
}

func Section(children ...HTML) *TagBuilder {
	return NewTagBuilder("section", children...)
}

func Search(children ...HTML) *TagBuilder {
	return NewTagBuilder("search", children...)
}

func Blockquote(children ...HTML) *TagBuilder {
	return NewTagBuilder("blockquote", children...)
}

func Dd(children ...HTML) *TagBuilder {
	return NewTagBuilder("dd", children...)
}

func Dl(children ...HTML) *TagBuilder {
	return NewTagBuilder("dl", children...)
}

func Dt(children ...HTML) *TagBuilder {
	return NewTagBuilder("dt", children...)
}

func Figcaption(children ...HTML) *TagBuilder {
	return NewTagBuilder("figcaption", children...)
}

func Figure(children ...HTML) *TagBuilder {
	return NewTagBuilder("figure", children...)
}

func Hr() *TagBuilder {
	return NewTagBuilder("hr").
		SetSelfClosing()
}

func Li(children ...HTML) *TagBuilder {
	return NewTagBuilder("li", children...)
}

func Menu(children ...HTML) *TagBuilder {
	return NewTagBuilder("menu", children...)
}

func OL(children ...HTML) *TagBuilder {
	return NewTagBuilder("ol", children...)
}

func P(children ...HTML) *TagBuilder {
	return NewTagBuilder("p", children...)
}

func Pre(children ...HTML) *TagBuilder {
	return NewTagBuilder("pre", children...)
}

func Ul(children ...HTML) *TagBuilder {
	return NewTagBuilder("ul", children...)
}

func Abbr(children ...HTML) *TagBuilder {
	return NewTagBuilder("abbr", children...)
}

func B(children ...HTML) *TagBuilder {
	return NewTagBuilder("b", children...)
}

func Bdi(children ...HTML) *TagBuilder {
	return NewTagBuilder("bdi", children...)
}

func Bdo(children ...HTML) *TagBuilder {
	return NewTagBuilder("bdo", children...)
}

func Br() *TagBuilder {
	return NewTagBuilder("br").
		SetSelfClosing()
}

func Cite(children ...HTML) *TagBuilder {
	return NewTagBuilder("cite", children...)
}

func Code(children ...HTML) *TagBuilder {
	return NewTagBuilder("code", children...)
}

func Data(children ...HTML) *TagBuilder {
	return NewTagBuilder("data", children...)
}

func Dfn(children ...HTML) *TagBuilder {
	return NewTagBuilder("dfn", children...)
}

func Em(children ...HTML) *TagBuilder {
	return NewTagBuilder("em", children...)
}

func I(children ...HTML) *TagBuilder {
	return NewTagBuilder("i", children...)
}

func Kbd(children ...HTML) *TagBuilder {
	return NewTagBuilder("kbd", children...)
}

func Mark(children ...HTML) *TagBuilder {
	return NewTagBuilder("mark", children...)
}

func Q(children ...HTML) *TagBuilder {
	return NewTagBuilder("q", children...)
}

func Rb(children ...HTML) *TagBuilder {
	return NewTagBuilder("rb", children...)
}

func Rp(children ...HTML) *TagBuilder {
	return NewTagBuilder("rp", children...)
}

func Rt(children ...HTML) *TagBuilder {
	return NewTagBuilder("rt", children...)
}

func Ruby(children ...HTML) *TagBuilder {
	return NewTagBuilder("ruby", children...)
}

func S(children ...HTML) *TagBuilder {
	return NewTagBuilder("s", children...)
}

func Samp(children ...HTML) *TagBuilder {
	return NewTagBuilder("samp", children...)
}

func Small(children ...HTML) *TagBuilder {
	return NewTagBuilder("small", children...)
}

func Span(children ...HTML) *TagBuilder {
	return NewTagBuilder("span", children...)
}

func Strong(children ...HTML) *TagBuilder {
	return NewTagBuilder("strong", children...)
}

func Sub(children ...HTML) *TagBuilder {
	return NewTagBuilder("sub", children...)
}

func Time(children ...HTML) *TagBuilder {
	return NewTagBuilder("time", children...)
}

func U(children ...HTML) *TagBuilder {
	return NewTagBuilder("u", children...)
}

func Var(children ...HTML) *TagBuilder {
	return NewTagBuilder("var", children...)
}

func Wbr() *TagBuilder {
	return NewTagBuilder("wbr").
		SetSelfClosing()
}

func Area() *TagBuilder {
	return NewTagBuilder("area").
		SetSelfClosing()
}

func Audio(children ...HTML) *TagBuilder {
	return NewTagBuilder("audio", children...)
}

func Img() *TagBuilder {
	return NewTagBuilder("img").
		SetSelfClosing()
}

func Map(children ...HTML) *TagBuilder {
	return NewTagBuilder("map", children...)
}

func Track() *TagBuilder {
	return NewTagBuilder("track").
		SetSelfClosing()
}

func Video(children ...HTML) *TagBuilder {
	return NewTagBuilder("video", children...)
}

func Embed() *TagBuilder {
	return NewTagBuilder("embed").
		SetSelfClosing()
}

func Iframe() *TagBuilder {
	return NewTagBuilder("iframe").
		SetSelfClosing()
}

func Object() *TagBuilder {
	return NewTagBuilder("object").
		SetSelfClosing()
}

func Picture(children ...HTML) *TagBuilder {
	return NewTagBuilder("picture", children...)
}

func Source() *TagBuilder {
	return NewTagBuilder("source").
		SetSelfClosing()
}

func Svg(children ...HTML) *TagBuilder {
	return NewTagBuilder("svg", children...)
}

func Math(children ...HTML) *TagBuilder {
	return NewTagBuilder("math", children...)
}

func Canvas() *TagBuilder {
	return NewTagBuilder("canvas")
}

func Noscript(children ...HTML) *TagBuilder {
	return NewTagBuilder("noscript", children...)
}

func Script(children ...HTML) *TagBuilder {
	return NewTagBuilder("script", children...)
}

func Del(children ...HTML) *TagBuilder {
	return NewTagBuilder("del", children...)
}

func Ins(children ...HTML) *TagBuilder {
	return NewTagBuilder("ins", children...)
}

func Caption(children ...HTML) *TagBuilder {
	return NewTagBuilder("caption", children...)
}

func Col() *TagBuilder {
	return NewTagBuilder("col").
		SetSelfClosing()
}

func Colgroup(children ...HTML) *TagBuilder {
	return NewTagBuilder("colgroup", children...)
}

func Table(children ...HTML) *TagBuilder {
	return NewTagBuilder("table", children...)
}

func Tbody(children ...HTML) *TagBuilder {
	return NewTagBuilder("tbody", children...)
}

func Td(children ...HTML) *TagBuilder {
	return NewTagBuilder("td", children...)
}

func Tfoot(children ...HTML) *TagBuilder {
	return NewTagBuilder("tfoot", children...)
}

func Th(children ...HTML) *TagBuilder {
	return NewTagBuilder("th", children...)
}

func Thead(children ...HTML) *TagBuilder {
	return NewTagBuilder("thead", children...)
}

func Tr(children ...HTML) *TagBuilder {
	return NewTagBuilder("tr", children...)
}

func Datalist(children ...HTML) *TagBuilder {
	return NewTagBuilder("datalist", children...)
}

func Fieldset(children ...HTML) *TagBuilder {
	return NewTagBuilder("fieldset", children...)
}

func Form(children ...HTML) *TagBuilder {
	return NewTagBuilder("form", children...)
}

func Label(children ...HTML) *TagBuilder {
	return NewTagBuilder("label", children...)
}

func Legend(children ...HTML) *TagBuilder {
	return NewTagBuilder("legend", children...)
}

func Meter(children ...HTML) *TagBuilder {
	return NewTagBuilder("meter", children...)
}

func Optgroup(children ...HTML) *TagBuilder {
	return NewTagBuilder("optgroup", children...)
}

func Option(label string) *TagBuilder {
	return NewTagBuilder("option", Text(label))
}

func Output(children ...HTML) *TagBuilder {
	return NewTagBuilder("output", children...)
}

func Progress(children ...HTML) *TagBuilder {
	return NewTagBuilder("progress", children...)
}

func Select(children ...HTML) *TagBuilder {
	return NewTagBuilder("select", children...)
}

func Textarea(content string) *TagBuilder {
	return NewTagBuilder("textarea", Text(content))
}

func Details(children ...HTML) *TagBuilder {
	return NewTagBuilder("details", children...)
}

func Dialog(children ...HTML) *TagBuilder {
	return NewTagBuilder("dialog", children...)
}

func Summary(children ...HTML) *TagBuilder {
	return NewTagBuilder("summary", children...)
}

func Slot() *TagBuilder {
	return NewTagBuilder("slot")
}

func Template(children ...HTML) *TagBuilder {
	return NewTagBuilder("template", children...)
}
