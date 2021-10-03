package headers

import (
	"fmt"

	"github.com/unidoc/unioffice/document"
)

func ParseHeader(text string, para document.Paragraph, doc *document.Document, level int) int {

	style := fmt.Sprintf("Heading%d", level)
	para.SetStyle(style)
	text = text[2:]

	ParseAndAddText(text, para, doc, false)

	return 101
}
