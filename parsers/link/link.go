package link

import (
	"fmt"
	"log"
	"md2docx/patterns"
	"md2docx/util"
	"regexp"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/document"
)

func ParseEmailUrl(text string, para document.Paragraph, doc *document.Document, _ int) int {
	firstIndex := strings.Index(text, "<")
	lastIndex := strings.Index(text, ">")
	text = text[firstIndex : len(text)-lastIndex]

	qTextPatt := regexp.MustCompile(patterns.QText)
	urlPattern := regexp2.MustCompile(patterns.Url, 0)

	explainText := qTextPatt.FindAllString(text, 1)[0]

	emailOrUrl := strings.Trim(strings.Replace(text, explainText, "", 1), " ")

	var protocol string

	if isEmail, _ := urlPattern.MatchString(emailOrUrl); isEmail {
		protocol = strings.Split(emailOrUrl, "://")[0] + "://"
		emailOrUrl = strings.Split(emailOrUrl, "://")[1]
	} else {
		protocol = "mailto"
	}

	if len(explainText) >= 1 {
		explainText = fmt.Sprintf(" :%s", explainText)
	}

	hl := para.AddHyperLink()
	hl.SetTarget(fmt.Sprintf("%s:%s", protocol, emailOrUrl))
	run := hl.AddRun()
	style := "Hyper Link"
	run.Properties().SetStyle(style)
	run.AddText(fmt.Sprintf("%s%s", emailOrUrl, explainText))
	hl.SetToolTip(explainText)

	return 101
}

func ParseLink(text string, para document.Paragraph, doc *document.Document, _ int) int {
	if strings.Count(text, "](") >= 1 {
		return ParseImageLink(text, para, doc, 0)
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

func ParseImageLink(text string, para document.Paragraph, doc *document.Document, _ int) int {
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

	patternUrl := regexp2.MustCompile(patterns.Url, 0)

	var toolTip string

	if len(imgTooltip) >= 1 {
		toolTip = imgTooltip
	} else {
		toolTip = linkText
	}

	var imgPath string
	var errDownload error

	if isUrl, _ := patternUrl.MatchString(imgPathOrUrl); isUrl {
		imgName := imgPathOrUrl[strings.LastIndex(imgPathOrUrl, "/"):]

		imgPath, errDownload = util.DownloadFile(imgName, imgPathOrUrl)

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

	imgRef, err := doc.ParseImage(img)
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
