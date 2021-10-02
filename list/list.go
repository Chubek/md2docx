package list

import (
	"strings"

	"github.com/unidoc/unioffice/document"
)

func AddList(text string, para document.Paragraph, _ *document.Document, _ int) (x int) {
	lstSplt := strings.Split(text, "\n")

	for _, line := range lstSplt {
		ParseAndAddText(line, para, false)
	}

	return 101

}
