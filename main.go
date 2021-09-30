package main

import (
	"fmt"
	"log"
	"os"
	"io"
	"path"
	"regexp"
	"net/http"
	"strings"
	"github.com/joho/godotenv"
	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
	"github.com/unidoc/unioffice/common"

)

func DownloadFile(filepath string, url string) error {
	err_env := godotenv.Load()
	if err_env != nil {
		log.Fatal("Error loading .env file")
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	imgSavePath := path.Join(os.Getenv(`FILES_PATH`), filepath)

	out, err := os.Create(imgSavePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
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
	emailText        = `(\<{1})(\S+@\S+)(\>{1})`
	qText            = `(\"|\')(\w|\W|\S)+(\"|\')`
	tableText        = `(((\|)([a-zA-Z\d+\s#!@'"():;\\\/.\[\]\^<={$}>?(?!-))]+))+(\|))(?:\n)?((\|)(-+))+(\|)(\n)((\|)(\W+|\w+|\S+))+(\|$)`
	threeUA	=	`(\_|\*){3}`
	twoUA	=	`(\_|\*){2}`
	oneUA	=	`(\_|\*){1}`
	url 	=	`(https?:\/\/(?:www\.|(?!www))[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|www\.[a-zA-Z0-9][a-zA-Z0-9-]+[a-zA-Z0-9]\.[^\s]{2,}|https?:\/\/(?:www\.|(?!www))[a-zA-Z0-9]+\.[^\s]{2,}|www\.[a-zA-Z0-9]+\.[^\s]{2,})`
	
)

func reverse(s string) string {
    rns := []rune(s) 
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

        rns[i], rns[j] = rns[j], rns[i]
    }
  
    return string(rns)
}
  
func addHeader(text string, para document.Paragraph, level int) {
	run := para.AddRun()
	style := fmt.Sprintf("Heading%d", level)
	para.SetStyle(style)
	text = text[2:]
	run.AddText(text)
}

func parseHeader(pattern *regexp.Regexp, text string) (x int) {
	if pattern.MatchString(text) {
		return 132
	}
	
	return 100
}

func parseLink(pattern *regexp.Regexp, text string) (x int) {
	if pattern.MatchString(text) {
		return 132
	}
	
	return 100
}

func addLink(text string, para document.Paragraph) {
	txtSplit := strings.Split(text, "](")
	linkText := txtSplit[0][1:]
	pattUnnTxt := regexp.MustCompile(`(\"|\')(\w|\W|\S)+(\"|\')`)
	hrefTooltip := pattUnnTxt.FindString(txtSplit[1])
	linkUrl := strings.Trim(pattUnnTxt.ReplaceAllString(txtSplit[1][:len(txtSplit[1]) - 1], ""), " ")

	hl := para.AddHyperLink()
	hl.SetTarget(linkUrl)
	run := hl.AddRun()
	run.Properties().SetStyle("Hyperlink")
	run.AddText(linkText)
	hl.SetToolTip(hrefTooltip)


}

func addImageLink(text string, para document.Paragraph) {
	firstIndex := strings.Index(text, "](")
	lastIndex := strings.LastIndex(text, "](")
	
	linkText = text[3:firstIndex]

	imgPathAndAlt := text[firstIndex:lastIndex - 1]

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

	var imgName string

	if patternUrl.MatchString(imgPathOrUrl) {
		imgName = imgPathOrUrl[strings.LastIndex(imgPathOrUrl, "/"):]

		DownloadFile(imgName, imgPathOrUrl)

	}

	hl := para.AddHyperLink()
	hl.SetTarget(linkUrl)
	run := hl.AddRun()
	run.Properties().SetStyle("Hyperlink")
	run.AddText(linkText)
	hl.SetToolTip(hrefTooltip)


}

func addBoldText(text string, para document.Paragraph) {
	run := para.AddRun()
	run.Properties().SetBold(true)
	text = text[2:len(text) - 2]
	run.AddText(text)

}

func addItalic(text string, para document.Paragraph) {
	run := para.AddRun()
	run.Properties().SetItalic(true)
	text = text[2:len(text) - 2]
	run.AddText(text)

}


func addItalicItalic(text string, para document.Paragraph) {
	run := para.AddRun()
	run.Properties().SetItalic(true)
	run.Properties().SetBold(true)
	text = text[3:len(text) - 3]
	run.AddText(text)

}

func parseBold(pattern *regexp.Regexp, text string) (x int) {
	twoUAP := regexp.MustCompile(twoUA)


	if pattern.MatchString(text) {
		firstTwo := text[:2]
		lastTwo := text[len(text) - 2:]

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
		lastTwo := text[len(text) - 2:]

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
		firstTwo := text[:2]
		lastTwo := text[len(text) - 2:]

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