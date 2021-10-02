package code

import (
	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

func AddInlineCode(text string, para document.Paragraph, _ *document.Document, _ int) int {
	run := para.AddRun()
	run.Properties().SetFontFamily("Courier New")
	run.Properties().SetColor(color.BlueViolet)
	run.AddText(text)
}

func AddCodeBlock(text string, para document.Paragraph, doc *document.Document, _ int) int {
	text = text[4 : len(text)-4]

	style := doc.Styles
	customStyle := style.AddStyle("CustomStyle1", wml.ST_StyleTypeParagraph, false)
	customStyle.SetName("Listing Style")
	customStyle.ParagraphProperties().SetSpacing(measurement.Inch*1, measurement.Inch*1)
	customStyle.ParagraphProperties().SetAlignment(wml.ST_JcBoth)
	customStyle.ParagraphProperties().SetFirstLineIndent(0)
	customStyle.ParagraphProperties().SetLineSpacing(2*measurement.Point, wml.ST_LineSpacingRuleAuto)
	
	run := para.AddRun()

	run.Properties().SetStyle("Listing Style")
	run.Properties().SetFontFamily("Courier New")
	run.Properties().SetKerning(2)
	run.Properties().SetColor(color.Blue)
	run.AddText(text)

	return 101

}
