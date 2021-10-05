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

type AdderContainer struct {
	Doc  *document.Document
	Para document.Paragraph
	Run  document.Run
}

func (adderContainer AdderContainer) AddBlockQuote(normalTexts []string) int {
	style := adderContainer.Doc.Styles
	customStyle := style.AddStyle("CustomStyle1", wml.ST_StyleTypeParagraph, false)
	customStyle.SetName("BQ Style")
	customStyle.ParagraphProperties().SetSpacing(measurement.Inch*1, measurement.Inch*1)
	customStyle.ParagraphProperties().SetAlignment(wml.ST_JcBoth)
	customStyle.ParagraphProperties().SetFirstLineIndent(8)
	customStyle.ParagraphProperties().SetLineSpacing(4*measurement.Point, wml.ST_LineSpacingRuleAuto)

	adderContainer.Para.SetStyle("BQ Style")

	Run := adderContainer.Para.AddRun()
	Run.Properties().SetFontFamily("Trebuchet MS")

	for _, txt := range normalTexts {
		Run.AddText(txt)
	}

	return 101
}

func (adderContainer AdderContainer) AddHorizLine(normalTexts []string) int {
	Run := adderContainer.Para.AddRun()
	Run.AddText(normalTexts[0])

	return 101
}

func (adderContainer AdderContainer) AddTable(tableTexts [][]string) int {
	table := adderContainer.Doc.AddTable()
	table.Properties().SetWidthPercent(100)
	borders := table.Properties().Borders()
	borders.SetAll(wml.ST_BorderSingle, color.Auto, 2*measurement.Point)

	for _, rowText := range tableTexts {
		row := table.AddRow()

		for _, cellText := range rowText {
			cell := row.AddCell()
			ParaCell := cell.AddParagraph()
			ParaCell.AddRun().AddText(cellText)
		}
	}

	return 101
}

func (adderContainer AdderContainer) AddList(normalTexts []string) int {
	for _, txt := range normalTexts {
		Run := adderContainer.Para.AddRun()
		Run.AddText(txt)
	}

	return 101

}

func (adderContainer AdderContainer) AddImage(normalTexts []string) int {
	imgPath := normalTexts[0]
	imgHint := normalTexts[1]

	img, err := common.ImageFromFile(imgPath)
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}

	imgRef, err := adderContainer.Doc.AddImage(img)
	if err != nil {
		log.Fatalf("unable to Parse image to Document: %s", err)
	}

	anchored, err := adderContainer.Para.AddRun().AddDrawingAnchored(imgRef)
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

func (adderContainer AdderContainer) AddEmailUrl(normalTexts []string) int {
	emailOrUrl := normalTexts[0]
	protocol := normalTexts[1]
	explainText := normalTexts[2]

	hl := adderContainer.Para.AddHyperLink()
	hl.SetTarget(fmt.Sprintf("%s:%s", protocol, emailOrUrl))
	Run := hl.AddRun()
	style := "Hyper Link"
	Run.Properties().SetStyle(style)
	Run.AddText(fmt.Sprintf("%s%s", emailOrUrl, explainText))
	hl.SetToolTip(explainText)

	return 101
}

func (adderContainer AdderContainer) AddLinkImage(normalTexts []string) int {
	imgPath := normalTexts[1]
	toolTip := normalTexts[2]
	linkUrl := normalTexts[3]

	img, err := common.ImageFromFile(imgPath)
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}

	imgRef, err := adderContainer.Doc.AddImage(img)
	if err != nil {
		log.Fatalf("unable to Parse image to Document: %s", err)
	}

	hl := adderContainer.Para.AddHyperLink()
	hl.SetTarget(linkUrl)
	Run := hl.AddRun()
	Run.Properties().SetStyle("Hyperlink")
	Run.AddDrawingAnchored(imgRef)
	hl.SetToolTip(toolTip)

	return 101
}

func (adderContainer AdderContainer) AddLink(text []string) int {
	if text[0] == "ImageLink" {
		return adderContainer.AddLinkImage(text)
	}

	linkText := text[1]
	hrefTooltip := text[2]
	linkUrl := text[3]

	hl := adderContainer.Para.AddHyperLink()
	hl.SetTarget(linkUrl)
	Run := hl.AddRun()
	Run.Properties().SetStyle("Hyperlink")
	Run.AddText(linkText)
	hl.SetToolTip(hrefTooltip)

	return 101

}

func (adderContainer AdderContainer) AddHeader(texts []string) int {
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

		hl := adderContainer.Para.AddHyperLink()
		adderContainer.Para.SetStyle(style)
		hl.SetTarget(linkUrl)
		Run := hl.AddRun()
		Run.Properties().SetStyle("Hyperlink")
		Run.AddText(linkText)
		hl.SetToolTip(hrefTooltip)

		return 101
	}

	adderContainer.Para.SetStyle(style)
	Run := adderContainer.Para.AddRun()
	Run.AddText(text)

	return 101
}

func (adderContainer AdderContainer) AddInlineCode(text []string) int {
	adderContainer.Run.Properties().SetFontFamily("Courier New")
	adderContainer.Run.Properties().SetColor(color.BlueViolet)

	for _, txt := range text {
		adderContainer.Run.AddText(txt)
	}

	return 101
}

func (adderContainer AdderContainer) AddBold(text []string) int {
	adderContainer.Run.Properties().SetBold(true)

	for _, txt := range text {
		adderContainer.Run.AddText(txt)
	}

	return 101
}

func (adderContainer AdderContainer) AddItalic(text []string) int {
	adderContainer.Run.Properties().SetItalic(true)

	for _, txt := range text {
		adderContainer.Run.AddText(txt)
	}

	return 101
}

func (adderContainer AdderContainer) AddBoldItalic(text []string) int {
	adderContainer.Run.Properties().SetBold(true)
	adderContainer.Run.Properties().SetItalic(true)

	for _, txt := range text {
		adderContainer.Run.AddText(txt)
	}

	return 101
}

func (adderContainer AdderContainer) AddCodeBlock(texts []string) int {
	text := texts[0]

	style := adderContainer.Doc.Styles
	customStyle := style.AddStyle("CustomStyle1", wml.ST_StyleTypeParagraph, false)
	customStyle.SetName("Listing Style")
	customStyle.ParagraphProperties().SetSpacing(measurement.Inch*1, measurement.Inch*1)
	customStyle.ParagraphProperties().SetAlignment(wml.ST_JcBoth)
	customStyle.ParagraphProperties().SetFirstLineIndent(0)
	customStyle.ParagraphProperties().SetLineSpacing(2*measurement.Point, wml.ST_LineSpacingRuleAuto)

	Run := adderContainer.Para.AddRun()

	Run.Properties().SetStyle("Listing Style")
	Run.Properties().SetFontFamily("Courier New")
	Run.Properties().SetKerning(2)
	Run.Properties().SetColor(color.Blue)
	Run.AddText(text)

	return 101

}


func (adderContainer AdderContainer) AddPlain(texts []string) int {
	adderContainer.Run.AddText(texts[0])

	return 101
}