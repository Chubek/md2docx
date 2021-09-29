package main

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	"github.com/unidoc/unioffice/schema/soo/wml"
)

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
	headerOne        = `(#{1}\s)(.*)`
	headeTwo         = `(#{2}\s)(.*)`
	headerThree      = `(#{3}\s)(.*)`
	headerFour       = `(#{4}\s)(.*)`
	headerFive       = `(#{5}\s)(.*)`
	headerSix        = `(#{6}\s)(.*)`
	boldItalicText   = `(\*|\_)+(\S+)(\*|\_)+`
	linkText         = `(\[.*\])(\((http)(?:s)?(\:\/\/).*\))`
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
	
)

func reverse(s string) string {
    rns := []rune(s) 
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

        rns[i], rns[j] = rns[j], rns[i]
    }
  
    return string(rns)
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
	twoUAP, err := regexp.Compile(twoUA)

	if err != nil {
		log.Fatal("Problem with compiling twoUA pattern!")
	}

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