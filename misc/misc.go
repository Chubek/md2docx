package misc

import (
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

func AddBlockQuote(text string, para document.Paragraph, doc *document.Document, _ int) int {
	style := doc.Styles
	customStyle := style.AddStyle("CustomStyle1", wml.ST_StyleTypeParagraph, false)
	customStyle.SetName("BQ Style")
	customStyle.ParagraphProperties().SetSpacing(measurement.Inch*1, measurement.Inch*1)
	customStyle.ParagraphProperties().SetAlignment(wml.ST_JcBoth)
	customStyle.ParagraphProperties().SetFirstLineIndent(8)
	customStyle.ParagraphProperties().SetLineSpacing(4*measurement.Point, wml.ST_LineSpacingRuleAuto)

	para.SetStyle("BQ Style")

	run := para.AddRun()
	run.Properties().SetFontFamily("Trebuchet MS")
	ParseAndAddText(text, para, false)

	return 101
}

func AddHorizontalLine(_ string, para document.Paragraph, _ *document.Document, _ int) int {
	run := para.AddRun()
	run.AddText("---\n")

	return 101

}
