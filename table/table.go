package table

import (
	"fmt"
	"md2docx/patterns"
	"regexp"
	"strings"
	"github.com/dlclark/regexp2"
	"github.com/unidoc/unioffice/document"
)



func AddTable(text string, _ document.Paragraph, doc *document.Document, _ int) int {
	tableSepPatt := regexp.MustCompile(patterns.TableSep)
	textSplit := tableSepPatt.Split(text, -1)
	var rowList [][]string

	for _, txtSplt := range textSplit {
		rowSplit := strings.Split(txtSplt[1:len(txtSplt)-1], "|")
		rowList = append(rowList, rowSplit)
	}

	table := doc.AddTable()
	table.Properties().SetWidthPercent(100)
	borders := table.Properties().Borders()
	borders.SetAll(wml.ST_BorderSingle, color.Auto, 2*measurement.Point)

	rowHeader := table.AddRow()

	for _, cellText := range rowList[0] {
		cell := rowHeader.AddCell()
		paraCell := cell.AddParagraph()
		ParseAndAddText(cellText, paraCell, true)
	}

	for _, rowText := range rowList[1:] {
		row := table.AddRow()

		for _, cellText := range rowText {
			cell := row.AddCell()
			paraCell := cell.AddParagraph()
			ParseAndAddText(cellText, paraCell, false)
		}
	}

	return 101
