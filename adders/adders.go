package adders

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

func AddBlockQuote(doc *document.Document, para document.Paragraph) int {
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

	return 101
}

func AddHorizLine(texts []string, para document.Paragraph) int {
	run := para.AddRun()
	run.AddText(texts[0])

	return 101
}

func AddTable(texts [][]string, doc *document.Document) int {
	table := doc.AddTable()
	table.Properties().SetWidthPercent(100)
	borders := table.Properties().Borders()
	borders.SetAll(wml.ST_BorderSingle, color.Auto, 2*measurement.Point)

	for _, rowText := range texts {
		row := table.AddRow()

		for _, cellText := range rowText {
			cell := row.AddCell()
			paraCell := cell.AddParagraph()
			paraCell.AddRun().AddText(cellText)
		}
	}

	return 101
}

func AddList(texts []string, para document.Paragraph) int {
	for _, txt := range texts {
		run := para.AddRun()
		run.AddText(txt)
	}

	return 101

}

func ParseImage(text []string, doc *document.Document, para document.Paragraph) int {
	imgPath := text[0]
	imgHint := text[1]

	img, err := common.ImageFromFile(imgPath)
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}

	imgRef, err := doc.AddImage(img)
	if err != nil {
		log.Fatalf("unable to Parse image to document: %s", err)
	}

	anchored, err := para.AddRun().AddDrawingAnchored(imgRef)
	if err != nil {
		log.Fatalf("unable to Parse anchored image: %s", err)
	}
	anchored.SetName(imgHint)
	anchored.SetSize(2*measurement.Inch, 2*measurement.Inch)
	anchored.SetOrigin(wml.WdST_RelFromHPage, wml.WdST_RelFromVTopMargin)
	anchored.SetHAlignment(wml.WdST_AlignHCenter)
	anchored.SetYOffset(3 * measurement.Inch)
	anchored.SetTextWrapSquare(wml.WdST_WrapTextBothSides)

	return 101

}

func AddEmailUrl(text []string, doc *document.Document, para document.Paragraph) int {
	emailOrUrl := text[0]
	protocol := text[1]
	explainText := text[2]

	hl := para.AddHyperLink()
	hl.SetTarget(fmt.Sprintf("%s:%s", protocol, emailOrUrl))
	run := hl.AddRun()
	style := "Hyper Link"
	run.Properties().SetStyle(style)
	run.AddText(fmt.Sprintf("%s%s", emailOrUrl, explainText))
	hl.SetToolTip(explainText)

	return 101
}

func AddLinkImage(text []string, doc *document.Document, para document.Paragraph) int {
	imgPath := text[1]
	toolTip := text[2]
	linkUrl := text[3]

	img, err := common.ImageFromFile(imgPath)
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}

	imgRef, err := doc.AddImage(img)
	if err != nil {
		log.Fatalf("unable to Parse image to document: %s", err)
	}

	hl := para.AddHyperLink()
	hl.SetTarget(linkUrl)
	run := hl.AddRun()
	run.Properties().SetStyle("Hyperlink")
	run.AddDrawingAnchored(imgRef)
	hl.SetToolTip(toolTip)

	return 101
}

func AddLink(text []string, doc *document.Document, para document.Paragraph) int {
	if text[0] == "ImageLink" {
		return AddLinkImage(text, doc, para)
	}

	linkText := text[1]
	hrefTooltip := text[2]
	linkUrl := text[3]

	hl := para.AddHyperLink()
	hl.SetTarget(linkUrl)
	run := hl.AddRun()
	run.Properties().SetStyle("Hyperlink")
	run.AddText(linkText)
	hl.SetToolTip(hrefTooltip)

	return 101

}

func AddHeader(texts []string, para document.Paragraph) int {
	level := texts[0]
	style := fmt.Sprintf("Heading%s", level)

	text := texts[1][2:]

	urlPattern := regexp.MustCompile(`(\[.*\])(\((\w{3,5})(\:\/\/).*\))`)

	if urlPattern.MatchString(text) {
		txtSplit := strings.Split(text, "](")
		linkText := txtSplit[0][1:]
		pattUnnTxt := regexp.MustCompile(`(\"|\')(\w|\W|\S)+(\"|\')`)
		hrefTooltip := pattUnnTxt.FindString(txtSplit[1])
		linkUrl := strings.Trim(pattUnnTxt.ReplaceAllString(txtSplit[1][:len(txtSplit[1])-1], ""), " ")

		hl := para.AddHyperLink()
		para.SetStyle(style)
		hl.SetTarget(linkUrl)
		run := hl.AddRun()
		run.Properties().SetStyle("Hyperlink")
		run.AddText(linkText)
		hl.SetToolTip(hrefTooltip)

		return 101
	}

	para.SetStyle(style)
	run := para.AddRun()
	run.AddText(text)

	return 101
}

func AddInlineCode(text []string, run document.Run) int {
	run.Properties().SetFontFamily("Courier New")
	run.Properties().SetColor(color.BlueViolet)

	for _, txt := range text {
		run.AddText(txt)
	}

	return 101
}

func AddBold(text []string, run document.Run) int {
	run.Properties().SetBold(true)

	for _, txt := range text {
		run.AddText(txt)
	}

	return 101
}

func AddItalic(text []string, run document.Run) int {
	run.Properties().SetItalic(true)

	for _, txt := range text {
		run.AddText(txt)
	}

	return 101
}

func AddBoldItalic(text []string, run document.Run) int {
	run.Properties().SetBold(true)
	run.Properties().SetItalic(true)

	for _, txt := range text {
		run.AddText(txt)
	}

	return 101
}

func AddCodeBlock(texts []string, doc *document.Document, para document.Paragraph) int {
	text := texts[0]

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
