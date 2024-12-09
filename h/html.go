package h

func Div(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "div",
		Attrs:    attrs,
		Children: children,
	}
}

func Html(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:         "html",
		Attrs:       attrs,
		Children:    children,
		SelfClosing: false,
	}
}

func Head(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:         "head",
		Attrs:       attrs,
		Children:    children,
		SelfClosing: false,
	}
}

func Body(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:         "body",
		Attrs:       attrs,
		Children:    children,
		SelfClosing: false,
	}
}

func H1(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "h1",
		Attrs:    attrs,
		Children: children,
	}
}

func Base(attrs []*Attribute) *Element {
	return &Element{
		Tag:         "base",
		Attrs:       attrs,
		SelfClosing: true,
	}
}

func Link(
	attrs []*Attribute,
) *Element {
	return &Element{
		Tag:         "link",
		Attrs:       attrs,
		SelfClosing: true,
	}
}

func Meta(attrs []*Attribute) *Element {
	return &Element{
		Tag:         "meta",
		Attrs:       attrs,
		SelfClosing: true,
	}
}

func Style(
	attrs []*Attribute,
	css string,
) *Element {
	return &Element{
		Tag:      "style",
		Attrs:    attrs,
		Children: []Renderable{Text(css)},
	}
}

func Title(
	attrs []*Attribute,
	title string,
) *Element {
	return &Element{
		Tag:      "title",
		Attrs:    attrs,
		Children: []Renderable{Text(title)},
	}
}

func A(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "a",
		Attrs:    attrs,
		Children: children,
	}
}

func Address(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "address",
		Attrs:    attrs,
		Children: children,
	}
}

func Article(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "article",
		Attrs:    attrs,
		Children: children,
	}
}

func Aside(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "aside",
		Attrs:    attrs,
		Children: children,
	}
}

func Footer(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "footer",
		Attrs:    attrs,
		Children: children,
	}
}

func Header(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "header",
		Attrs:    attrs,
		Children: children,
	}
}

func H2(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "h2",
		Attrs:    attrs,
		Children: children,
	}
}

func H3(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "h3",
		Attrs:    attrs,
		Children: children,
	}
}

func H4(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "h4",
		Attrs:    attrs,
		Children: children,
	}
}

func H5(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "h5",
		Attrs:    attrs,
		Children: children,
	}
}

func H6(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "h6",
		Attrs:    attrs,
		Children: children,
	}
}

func HGroup(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "hgroup",
		Attrs:    attrs,
		Children: children,
	}
}

func Main(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "main",
		Attrs:    attrs,
		Children: children,
	}
}

func Nav(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "nav",
		Attrs:    attrs,
		Children: children,
	}
}

func Section(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "section",
		Attrs:    attrs,
		Children: children,
	}
}
func Search(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "search",
		Attrs:    attrs,
		Children: children,
	}
}

func Blockquote(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "blockquote",
		Attrs:    attrs,
		Children: children,
	}
}

func DD(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "dd",
		Attrs:    attrs,
		Children: children,
	}
}

func DL(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "dl",
		Attrs:    attrs,
		Children: children,
	}
}
func DT(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "dt",
		Attrs:    attrs,
		Children: children,
	}
}

func Figcaption(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "figcaption",
		Attrs:    attrs,
		Children: children,
	}
}

func Figure(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "figure",
		Attrs:    attrs,
		Children: children,
	}
}

func HR(
	attrs []*Attribute,
) *Element {
	return &Element{
		Tag:         "hr",
		Attrs:       attrs,
		SelfClosing: true,
	}
}

func Li(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "li",
		Attrs:    attrs,
		Children: children,
	}
}
func Menu(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "menu",
		Attrs:    attrs,
		Children: children,
	}
}

func OL(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "ol",
		Attrs:    attrs,
		Children: children,
	}
}
func P(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "p",
		Attrs:    attrs,
		Children: children,
	}
}
func Pre(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "pre",
		Attrs:    attrs,
		Children: children,
	}
}

func UL(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "ul",
		Attrs:    attrs,
		Children: children,
	}
}

func Abbr(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "abbr",
		Attrs:    attrs,
		Children: children,
	}
}
func B(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "b",
		Attrs:    attrs,
		Children: children,
	}
}
func Bdi(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "bdi",
		Attrs:    attrs,
		Children: children,
	}
}

func Bdo(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "bdo",
		Attrs:    attrs,
		Children: children,
	}
}

func Br(
	attrs []*Attribute,
) *Element {
	return &Element{
		Tag:         "br",
		Attrs:       attrs,
		SelfClosing: true,
	}
}

func Cite(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "cite",
		Attrs:    attrs,
		Children: children,
	}
}

func Code(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "code",
		Attrs:    attrs,
		Children: children,
	}
}

func Data(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "data",
		Attrs:    attrs,
		Children: children,
	}
}

func Dfn(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "dfn",
		Attrs:    attrs,
		Children: children,
	}
}

func Em(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "em",
		Attrs:    attrs,
		Children: children,
	}
}

func I(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "i",
		Attrs:    attrs,
		Children: children,
	}
}

func Kbd(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "kbd",
		Attrs:    attrs,
		Children: children,
	}
}

func Mark(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "mark",
		Attrs:    attrs,
		Children: children,
	}
}

func Q(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "q",
		Attrs:    attrs,
		Children: children,
	}
}

func Rb(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "rb",
		Attrs:    attrs,
		Children: children,
	}
}

func Rp(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "rp",
		Attrs:    attrs,
		Children: children,
	}
}

func Rt(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "rt",
		Attrs:    attrs,
		Children: children,
	}
}
func Ruby(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "ruby",
		Attrs:    attrs,
		Children: children,
	}
}
func S(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "s",
		Attrs:    attrs,
		Children: children,
	}
}

func Samp(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "samp",
		Attrs:    attrs,
		Children: children,
	}
}

func Small(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "small",
		Attrs:    attrs,
		Children: children,
	}
}

func Span(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "span",
		Attrs:    attrs,
		Children: children,
	}
}

func Strong(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "strong",
		Attrs:    attrs,
		Children: children,
	}
}

func Sub(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "sub",
		Attrs:    attrs,
		Children: children,
	}
}

func Time(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "time",
		Attrs:    attrs,
		Children: children,
	}
}

func U(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "u",
		Attrs:    attrs,
		Children: children,
	}
}

func Var(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "var",
		Attrs:    attrs,
		Children: children,
	}
}

func Wbr(
	attrs []*Attribute,
) *Element {
	return &Element{
		Tag:         "wbr",
		Attrs:       attrs,
		SelfClosing: true,
	}
}

func Area(
	attrs []*Attribute,
) *Element {
	return &Element{
		Tag:         "area",
		Attrs:       attrs,
		SelfClosing: true,
	}
}

func Audio(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "audio",
		Attrs:    attrs,
		Children: children,
	}
}

func Img(
	attrs []*Attribute,
) *Element {
	return &Element{
		Tag:         "img",
		Attrs:       attrs,
		SelfClosing: true,
	}
}

func Map(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "map",
		Attrs:    attrs,
		Children: children,
	}
}

func Track(
	attrs []*Attribute,
) *Element {
	return &Element{
		Tag:         "track",
		Attrs:       attrs,
		SelfClosing: true,
	}
}

func Video(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "video",
		Attrs:    attrs,
		Children: children,
	}
}

func Embed(
	attrs []*Attribute,
) *Element {
	return &Element{
		Tag:         "embed",
		Attrs:       attrs,
		SelfClosing: true,
	}
}

func Iframe(
	attrs []*Attribute,
) *Element {
	return &Element{
		Tag:         "iframe",
		Attrs:       attrs,
		SelfClosing: true,
	}
}

func Object(
	attrs []*Attribute,
) *Element {
	return &Element{
		Tag:         "object",
		Attrs:       attrs,
		SelfClosing: true,
	}
}

func Picture(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "picture",
		Attrs:    attrs,
		Children: children,
	}
}

func Portal(
	attrs []*Attribute,
) *Element {
	return &Element{
		Tag:   "portal",
		Attrs: attrs,
	}
}

func Source(
	attrs []*Attribute,
) *Element {
	return &Element{
		Tag:   "source",
		Attrs: attrs,
		Void:  true,
	}
}

func Svg(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "svg",
		Attrs:    attrs,
		Children: children,
	}
}

func Math(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "math",
		Attrs:    attrs,
		Children: children,
	}
}

func Canvas(
	attrs []*Attribute,
) *Element {
	return &Element{
		Tag:   "canvas",
		Attrs: attrs,
	}
}

func Noscript(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "noscript",
		Attrs:    attrs,
		Children: children,
	}
}

func Script(
	attrs []*Attribute,
	js string,
) *Element {
	return &Element{
		Tag:      "script",
		Attrs:    attrs,
		Children: []Renderable{Text(js)},
	}
}

func Del(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "del",
		Attrs:    attrs,
		Children: children,
	}
}

func Ins(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "ins",
		Attrs:    attrs,
		Children: children,
	}
}

func Caption(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "caption",
		Attrs:    attrs,
		Children: children,
	}
}

func Col(
	attrs []*Attribute,
) *Element {
	return &Element{
		Tag:         "col",
		Attrs:       attrs,
		SelfClosing: true,
	}
}

func Colgroup(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "colgroup",
		Attrs:    attrs,
		Children: children,
	}
}
func Table(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "table",
		Attrs:    attrs,
		Children: children,
	}
}

func Tbody(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "tbody",
		Attrs:    attrs,
		Children: children,
	}
}

func Td(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "td",
		Attrs:    attrs,
		Children: children,
	}
}

func Tfoot(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "tfoot",
		Attrs:    attrs,
		Children: children,
	}
}

func Th(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "th",
		Attrs:    attrs,
		Children: children,
	}
}

func Thead(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "thead",
		Attrs:    attrs,
		Children: children,
	}
}

func Tr(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "tr",
		Attrs:    attrs,
		Children: children,
	}
}

func Button(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "button",
		Attrs:    attrs,
		Children: children,
	}
}

func Datalist(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "datalist",
		Attrs:    attrs,
		Children: children,
	}
}

func Fieldset(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "fieldset",
		Attrs:    attrs,
		Children: children,
	}
}

func Form(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "form",
		Attrs:    attrs,
		Children: children,
	}
}

func Input(
	attrs []*Attribute,
) *Element {
	return &Element{
		Tag:         "input",
		Attrs:       attrs,
		SelfClosing: true,
	}
}

func Label(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "label",
		Attrs:    attrs,
		Children: children,
	}
}

func Legend(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "legend",
		Attrs:    attrs,
		Children: children,
	}
}

func Meter(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "meter",
		Attrs:    attrs,
		Children: children,
	}
}

func Optgroup(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "optgroup",
		Attrs:    attrs,
		Children: children,
	}
}

func Option(
	attrs []*Attribute,
	label string,
) *Element {
	return &Element{
		Tag:      "option",
		Attrs:    attrs,
		Children: []Renderable{Text(label)},
	}
}

func Output(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "output",
		Attrs:    attrs,
		Children: children,
	}
}

func Progress(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "progress",
		Attrs:    attrs,
		Children: children,
	}
}

func Select(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "select",
		Attrs:    attrs,
		Children: children,
	}
}

func Textarea(
	attrs []*Attribute,
	content string,
) *Element {
	return &Element{
		Tag:      "textarea",
		Attrs:    attrs,
		Children: []Renderable{Text(content)},
	}
}

func Details(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "details",
		Attrs:    attrs,
		Children: children,
	}
}

func Dialog(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "dialog",
		Attrs:    attrs,
		Children: children,
	}
}

func Summary(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "summary",
		Attrs:    attrs,
		Children: children,
	}
}

func Slot(
	attrs []*Attribute,
) *Element {
	return &Element{
		Tag:   "slot",
		Attrs: attrs,
	}
}

func Template(
	attrs []*Attribute,
	children ...Renderable,
) *Element {
	return &Element{
		Tag:      "template",
		Attrs:    attrs,
		Children: children,
	}
}
