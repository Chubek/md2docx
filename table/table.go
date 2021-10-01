package table

import (
	"fmt"
	"md2docx/patterns"
	"regexp"
	"strings"
	"github.com/dlclark/regexp2"
	"github.com/unidoc/unioffice/document"
)



func AddTable(text string, doc *document.Document) {
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
		run := cell.AddParagraph().AddRun()
		run.Properties().SetBold(true)
		ParseAndAddText(cellText, run)
	}

	for _, rowText := range rowList[1:] {
		row := table.AddRow()

		for _, cellText := range rowText {
			cell := row.AddCell()
			run := cell.AddParagraph().AddRun()
			ParseAndAddText(cellText, run)
		}
	}
