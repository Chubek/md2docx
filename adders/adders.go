package adders

import (
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
	"github.com/unidoc/unioffice/color"
)


func AddBlockQuote(texts []string, doc *document.Document) int {
	style := doc.Styles
	customStyle := style.AddStyle("CustomStyle1", wml.ST_StyleTypeParagraph, false)
	customStyle.SetName("BQ Style")
	customStyle.ParagraphProperties().SetSpacing(measurement.Inch*1, measurement.Inch*1)
	customStyle.ParagraphProperties().SetAlignment(wml.ST_JcBoth)
	customStyle.ParagraphProperties().SetFirstLineIndent(8)
	customStyle.ParagraphProperties().SetLineSpacing(4*measurement.Point, wml.ST_LineSpacingRuleAuto)
	
	para := doc.AddParagraph()

	para.SetStyle("BQ Style")

	run := para.AddRun()
	run.Properties().SetFontFamily("Trebuchet MS")
	
	for _, txt := range texts {
		run.AddText(txt)
	}

	return 101
}

func AddTable(texts [][]string, doc *document.Document) int {
	table := doc.AddTable()
	table.Properties().SetWidthPercent(100)
	borders := table.Properties().Borders()
	borders.SetAll(wml.ST_BorderSingle, color.Auto, 2*measurement.Point)

	for _, rowText := range  texts{
		row := table.AddRow()

		for _, cellText := range rowText {
			cell := row.AddCell()
			paraCell := cell.AddParagraph()
			paraCell.AddRun().AddText(cellText)
		}
	}

	return 101
}


func AddList(texts []string, doc *document.Document) int {
	para := doc.AddParagraph()

	for _, txt := range texts {
		run := para.AddRun()
		run.AddText(txt)
	}

	return 101

}