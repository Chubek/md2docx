package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/unidoc/unioffice/common/license"
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
)
