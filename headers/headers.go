package headers

import (
	"fmt"

	"github.com/unidoc/unioffice/document"
)

func AddHeader(text string, para document.Paragraph, _ *document.Document, level int) int {
	
	style := fmt.Sprintf("Heading%d", level)
	para.SetStyle(style)
	text = text[2:]

	ParseAndAddText(text, para, false)

	return 101
}
