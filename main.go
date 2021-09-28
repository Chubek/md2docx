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
	headerOne        = `(#{1})(.*)`
	headeTwo         = `(#{2})(.*)`
	headerThree      = `(#{3})(.*)`
	headerFour       = `(#{4})(.*)`
	headerFive       = `(#{5})(.*)`
	headerFix        = `(#{6})(.*)`
	boldItalicText   = `(\*|\_)+(\S+)(\*|\_)+`
	linkText         = `(\[.*\])(\((http)(?:s)?(\:\/\/).*\))`
	imageFile        = `(\!)(\[(?:.*)?\])(\(.*(\.(jpg|png|gif|tiff|bmp))(?:(\s\"|\')(\w|\W|\d)+(\"|\'))?\))`
	listText         = `((\W{1})(\s)(.*)(?:\n)?)+`
	underlineText    = `\_{1}.*\_{1}`
	numberedListText = `((\d+\.)(\s)(.*)(?:\n)?)+`
	blockQuote       = `((\>{1})(\s)(.*)(?:\n)?)+`
	inlineCode       = "(\\`{1})(.*)(\\`{1})"
	codeBlock        = "(\\`{3}\\n)(.*)(\\n\\`{3})"
	horizontalLine   = `(\=|\-|\*){3}`
	emailText        = `(\<{1})(\S+@\S+)(\>{1})`
	boldItalicLink   = `(\*|\_)+(\[.*\])(\((http)(?:s)?(\:\/\/).*\))(\*|\_)+`
	qText            = `(\"|\')(\w|\W|\d)+(\"|\')`
)
