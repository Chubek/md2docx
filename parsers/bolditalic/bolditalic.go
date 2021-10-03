package bolditalic

import (
	"github.com/unidoc/unioffice/document"
)

func ParseBoldText(text string, para document.Paragraph, _ *document.Document, _ int) int {
	run := para.AddRun()
	run.Properties().SetBold(true)
	text = text[2 : len(text)-2]
	run.AddText(text)

	return 101

}

func ParseItalic(text string, para document.Paragraph, _ *document.Document, _ int) int {
	run := para.AddRun()
	run.Properties().SetItalic(true)
	text = text[2 : len(text)-2]
	run.AddText(text)

	return 101

}

func ParseBoldItalic(text string, para document.Paragraph, _ *document.Document, _ int) int {
	run := para.AddRun()
	run.Properties().SetItalic(true)
	run.Properties().SetBold(true)
	text = text[3 : len(text)-3]
	run.AddText(text)

	return 101

}
