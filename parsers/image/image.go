package image

import (
	"fmt"
	"log"
	"md2docx/patterns"
	"md2docx/util"
	"regexp"
	"strings"

	"github.com/unidoc/unioffice/schema/soo/wml"

	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
)

func ParseImage(text string, para document.Paragraph, doc *document.Document, _ int) int {
	firstIndex := strings.Index(text, "](")
	lastIndex := strings.LastIndex(text, ")")

	imgText := text[2:firstIndex]

	imgPathAndAlt := text[firstIndex : lastIndex-1]

	pattUnnTxt := regexp.MustCompile(patterns.UnText)
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

	patternUrl := regexp.MustCompile(patterns.Url)

	var imgPath string
	var errDownload error

	var imgHint string

	if len(imgText) >= 1 {
		imgHint = imgText
	} else {
		imgHint = imgTooltip
	}

	if patternUrl.MatchString(imgPathOrUrl) {
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