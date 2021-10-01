package patterns

const (
	HeaderAll        = `(#{%d}\s)(.*)`
	BoldItalicText   = `(\*|\_)+(\S+)(\*|\_)+`
	LinkText         = `(\[.*\])(\((\w{3,5})(\:\/\/).*\))`
	ImageFile        = `(\!)(\[(?:.*)?\])(\(.*(\.(jpg|png|gif|tiff|bmp))(?:(\s\"|\')(\w|\W|\d)+(\"|\'))?\))`
	ListText         = `(^(\W{1})(\s)(.*)(?:$)?)+`
	NumberedListText = `(^(\d+\.)(\s)(.*)(?:$)?)+`
	BlockQuote       = `(^(\>{1})(\s)(.*)(?:$)?)+`
	InlineCode       = "(\\`{1})(.*)(\\`{1})"
	CodeBlock        = "(\\`{3}\\n+)(.*)(\\n+\\`{3})"
	HorizontalLine   = `(\=|\-|\*){3}`
	EmailUrlText     = `(\<{1})((\S+@\S+)|(\w{3,5}?:\/\/(?:www\.|(?!www))[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|www\.[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|https?:\/\/(?:www\.|(?!www))[a-zA-Z0-9]+\.[^\s]{2,}|www\.[a-zA-Z0-9]+\.[^\s]{2,}))(\>{1})`
	QText            = `(\"|\')(\w|\W|\S)+(\"|\')`
	TableText        = `(((\|)([a-zA-Z\d+\s#!@'"():;\\\/.\[\]\^<={$}>?(?!-))]+))+(\|))(?:\n)?((\|)(-+))+(\|)(\n)((\|)(\W+|\w+|\S+))+(\|$)`
	ThreeUA          = `(\_|\*){3}`
	TwoUA            = `(\_|\*){2}`
	OneUA            = `(\_|\*){1}`
	Url              = `(https?:\/\/(?:www\.|(?!www))[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|www\.[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|https?:\/\/(?:www\.|(?!www))[a-zA-Z0-9]+\.[^\s]{2,}|www\.[a-zA-Z0-9]+\.[^\s]{2,})`
	TableSep         = `((\|)(-+))+(\|)(\n)`
)
