package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

func DownloadFile(filepath string, url string) (string, error) {
	err_env := godotenv.Load()
	if err_env != nil {
		log.Fatal("Error loading .env file")
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	imgSavePath := path.Join(os.Getenv(`FILES_PATH`), filepath)

	out, err := os.Create(imgSavePath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return imgSavePath, err
}

func init() {
	err_env := godotenv.Load()
	if err_env != nil {
		log.Fatal("Error loading .env file")
	}

	err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	if err != nil {
		panic(err)
	}
}

const (
	headerAll        = `(#{%d}\s)(.*)`
	boldItalicText   = `(\*|\_)+(\S+)(\*|\_)+`
	linkText         = `(\[.*\])(\((\w{3,5})(\:\/\/).*\))`
	imageFile        = `(\!)(\[(?:.*)?\])(\(.*(\.(jpg|png|gif|tiff|bmp))(?:(\s\"|\')(\w|\W|\d)+(\"|\'))?\))`
	listText         = `(^(\W{1})(\s)(.*)(?:$)?)+`
	numberedListText = `(^(\d+\.)(\s)(.*)(?:$)?)+`
	blockQuote       = `(^(\>{1})(\s)(.*)(?:$)?)+`
	inlineCode       = "(\\`{1})(.*)(\\`{1})"
	codeBlock        = "(\\`{3}\\n+)(.*)(\\n+\\`{3})"
	horizontalLine   = `(\=|\-|\*){3}`
	emailUrlText     = `(\<{1})((\S+@\S+)|(\w{3,5}?:\/\/(?:www\.|(?!www))[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|www\.[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|https?:\/\/(?:www\.|(?!www))[a-zA-Z0-9]+\.[^\s]{2,}|www\.[a-zA-Z0-9]+\.[^\s]{2,}))(\>{1})`
	qText            = `(\"|\')(\w|\W|\S)+(\"|\')`
	tableText        = `(((\|)([a-zA-Z\d+\s#!@'"():;\\\/.\[\]\^<={$}>?(?!-))]+))+(\|))(?:\n)?((\|)(-+))+(\|)(\n)((\|)(\W+|\w+|\S+))+(\|$)`
	threeUA          = `(\_|\*){3}`
	twoUA            = `(\_|\*){2}`
	oneUA            = `(\_|\*){1}`
	url              = `(https?:\/\/(?:www\.|(?!www))[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|www\.[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|https?:\/\/(?:www\.|(?!www))[a-zA-Z0-9]+\.[^\s]{2,}|www\.[a-zA-Z0-9]+\.[^\s]{2,})`
	tableSep		 =	`((\|)(-+))+(\|)(\n)`
)

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

func addTable(text string, doc *document.Document) {
	tableSepPatt := regexp.MustCompile(tableSep)
	textSplit := tableSepPatt.Split(text, -1)
	var rowList [][]string

	for _, txtSplt := range textSplit {
		rowSplit := strings.Split(txtSplt[1:len(txtSplt) - 1], "|")
		rowList = append(rowList, rowSplit)
	}

	table := doc.AddTable()
	table.Properties().SetWidthPercent(100)
	borders := table.Properties().Borders()
	borders.SetAll(wml.ST_BorderSingle, color.Auto, 2*measurement.Point)

	rowHeader := table.AddRow()

	for _, cellText := range rowList[0]{
		cell := rowHeader.AddCell()
		run := cell.AddParagraph().AddRun()
		run.Properties().SetBold(true)
		run.AddText(cellText)
	}

	for _, rowText := range rowList[1:] {
		row := table.AddRow()

		for _, cellText := range rowText {
			cell := row.AddCell()
			run := cell.AddParagraph().AddRun()
			run.AddText(cellText)
		}
	}
	


}

func addList(text string, para document.Paragraph) {
	run := para.AddRun()
	run.AddText(text)
}

func addInlineCode(text string, para document.Paragraph) {
	run := para.AddRun()
	run.Properties().SetFontFamily("Courier New")
	run.Properties().SetColor(color.BlueViolet)
	run.AddText(text)
}

func addBlockQuote(text string, para document.Paragraph, doc *document.Document) {
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
	run.AddText(text)
}

func addEmailUrl(text string, para document.Paragraph) {
	firstIndex := strings.Index(text, "<")
	lastIndex := strings.Index(text, ">")
	text = text[firstIndex : len(text)-lastIndex]

	qTextPatt := regexp.MustCompile(qText)
	urlPattern := regexp.MustCompile(url)

	explainText := qTextPatt.FindAllString(text, 1)[0]

	emailOrUrl := strings.Trim(strings.Replace(text, explainText, "", 1), " ")

	var protocol string

	if urlPattern.MatchString(emailOrUrl) {
		protocol = strings.Split(emailOrUrl, "://")[0] + "://"
		emailOrUrl = strings.Split(emailOrUrl, "://")[1]
	} else {
		protocol = "mailto"
	}

	if len(explainText) >= 1 {
		explainText = fmt.Sprintf(" :%", explainText)
	}

	hl := para.AddHyperLink()
	hl.SetTarget(fmt.Sprintf("%s:%s", protocol, emailOrUrl))
	run := hl.AddRun()
	style := fmt.Sprintf("Hyper Link")
	run.Properties().SetStyle(style)
	run.AddText(fmt.Sprint("%s%s", emailOrUrl, explainText))
	hl.SetToolTip(explainText)
}

func parseEmailUrl(patternUrlEmail *regexp.Regexp, text string) int {
	if patternUrlEmail.MatchString(text) {
		return 124
	}

	return 100
}

func addCodeBlock(text string, para document.Paragraph, doc *document.Document) {
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

}

func addHorizontalLine(para document.Paragraph) int {
	run := para.AddRun()
	run.AddText("---\n")

	return 101

}

func addHeader(text string, para document.Paragraph, level int) int {
	run := para.AddRun()
	style := fmt.Sprintf("Heading%d", level)
	para.SetStyle(style)
	text = text[2:]

	linkPattern := regexp.MustCompile(url)

	if linkPattern.MatchString(text) {
		return addLinkHeader(text, para, level)
	}

	run.AddText(text)

	return 101
}

func addLinkHeader(text string, para document.Paragraph, level int) int {
	txtSplit := strings.Split(text, "](")
	linkText := txtSplit[0][1:]
	pattUnnTxt := regexp.MustCompile(`(\"|\')(\w|\W|\S)+(\"|\')`)
	hrefTooltip := pattUnnTxt.FindString(txtSplit[1])
	linkUrl := strings.Trim(pattUnnTxt.ReplaceAllString(txtSplit[1][:len(txtSplit[1])-1], ""), " ")

	hl := para.AddHyperLink()
	hl.SetTarget(linkUrl)
	run := hl.AddRun()
	style := fmt.Sprintf("Heading%d", level)
	run.Properties().SetStyle(style)
	run.AddText(linkText)
	hl.SetToolTip(hrefTooltip)

	return 101
}

func parseNormal(pattern *regexp.Regexp, text string) (x int) {
	if pattern.MatchString(text) {
		return 132
	}

	return 100
}

func addLink(text string, para document.Paragraph, doc *document.Document) int {
	if strings.Count(text, "](") >= 1 {
		return addImageLink(text, para, doc)
	}

	txtSplit := strings.Split(text, "](")
	linkText := txtSplit[0][1:]
	pattUnnTxt := regexp.MustCompile(`(\"|\')(\w|\W|\S)+(\"|\')`)
	hrefTooltip := pattUnnTxt.FindString(txtSplit[1])
	linkUrl := strings.Trim(pattUnnTxt.ReplaceAllString(txtSplit[1][:len(txtSplit[1])-1], ""), " ")

	hl := para.AddHyperLink()
	hl.SetTarget(linkUrl)
	run := hl.AddRun()
	run.Properties().SetStyle("Hyperlink")
	run.AddText(linkText)
	hl.SetToolTip(hrefTooltip)

	return 101

}

func addImageLink(text string, para document.Paragraph, doc *document.Document) int {
	firstIndex := strings.Index(text, "](")
	lastIndex := strings.LastIndex(text, "](")

	linkText := text[3:firstIndex]

	imgPathAndAlt := text[firstIndex : lastIndex-1]

	pattUnnTxt := regexp.MustCompile(`(\"|\')(\w|\W|\S)+(\"|\')`)
	imgTooltip := pattUnnTxt.FindString(imgPathAndAlt)

	singleQuoteIndex := strings.Index(imgPathAndAlt, "'")
	doubleQuoteIndex := strings.Index(imgPathAndAlt, "\"")

	var quoteIndex int

	if singleQuoteIndex == -1 {
		quoteIndex = doubleQuoteIndex
	} else {
		quoteIndex = singleQuoteIndex
	}

	imgPathOrUrl := strings.Trim(pattUnnTxt.ReplaceAllString(imgPathAndAlt[:quoteIndex], ""), " ")

	lastParanIndex := strings.LastIndex(text, ")")

	linkUrl := strings.Trim(text[lastIndex:lastParanIndex], " ")

	patternUrl := regexp.MustCompile(url)

	var toolTip string

	if len(imgTooltip) >= 1 {
		toolTip = imgTooltip
	} else {
		toolTip = linkText
	}

	var imgPath string
	var errDownload error

	if patternUrl.MatchString(imgPathOrUrl) {
		imgName := imgPathOrUrl[strings.LastIndex(imgPathOrUrl, "/"):]

		imgPath, errDownload = DownloadFile(imgName, imgPathOrUrl)

		if errDownload != nil {
			log.Fatal(fmt.Sprintf("Error downloading image %s", errDownload))
		}

	} else {
		imgPath = imgPathOrUrl
	}

	img, err := common.ImageFromFile(imgPath)
	if err != nil {
		log.Fatalf("unable to create image: %s", err)
	}

	imgRef, err := doc.AddImage(img)
	if err != nil {
		log.Fatalf("unable to add image to document: %s", err)
	}

	hl := para.AddHyperLink()
	hl.SetTarget(linkUrl)
	run := hl.AddRun()
	run.Properties().SetStyle("Hyperlink")
	run.AddDrawingAnchored(imgRef)
	hl.SetToolTip(toolTip)

	return 101

}

func addBoldText(text string, para document.Paragraph) int {
	run := para.AddRun()
	run.Properties().SetBold(true)
	text = text[2 : len(text)-2]
	run.AddText(text)

	return 101

}

func addItalic(text string, para document.Paragraph) int {
	run := para.AddRun()
	run.Properties().SetItalic(true)
	text = text[2 : len(text)-2]
	run.AddText(text)

	return 101

}

func addItalicItalic(text string, para document.Paragraph) int {
	run := para.AddRun()
	run.Properties().SetItalic(true)
	run.Properties().SetBold(true)
	text = text[3 : len(text)-3]
	run.AddText(text)

	return 101

}

func parseBold(pattern *regexp.Regexp, text string) (x int) {
	twoUAP := regexp.MustCompile(twoUA)

	if pattern.MatchString(text) {
		firstTwo := text[:2]
		lastTwo := text[len(text)-2:]

		if twoUAP.MatchString(firstTwo) && twoUAP.MatchString(lastTwo) {
			if firstTwo != reverse(lastTwo) {
				return 121
			} else {
				return 132
			}
		}
	}

	return 100
}

func parseItalic(pattern *regexp.Regexp, text string) (x int) {
	oneUAP, err := regexp.Compile(oneUA)

	if err != nil {
		log.Fatal("Problem with compiling oneUA pattern!")
	}

	if pattern.MatchString(text) {
		firstTwo := text[:2]
		lastTwo := text[len(text)-2:]

		if oneUAP.MatchString(firstTwo) && oneUAP.MatchString(lastTwo) {
			if firstTwo != reverse(lastTwo) {
				return 121
			} else {
				return 132
			}
		}
	}

	return 100
}

func parseBoldItalic(pattern *regexp.Regexp, text string) (x int) {
	threeUAP, err := regexp.Compile(threeUA)

	if err != nil {
		log.Fatal("Problem with compiling threeUA pattern!")
	}

	if pattern.MatchString(text) {
		firstTwo := text[:3]
		lastTwo := text[len(text)-3:]

		if threeUAP.MatchString(firstTwo) && threeUAP.MatchString(lastTwo) {
			if firstTwo != reverse(lastTwo) {
				return 121
			} else {
				return 132
			}
		}
	}

	return 100
}

func main() {
	doc := document.New()
	defer doc.Close()

	para := doc.AddParagraph()

	pattern, err := regexp.Compile(boldItalicText)

}
