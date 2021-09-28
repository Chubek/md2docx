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
	headerOne    = `(#{1})(.*)`
	headeTwo     = `(#{2})(.*)`
	headerThree  = `(#{3})(.*)`
	headerFour   = `(#{4})(.*)`
	headerFive   = `(#{5})(.*)`
	headerFix    = `(#{6})(.*)`
	asteriskText = `\*.*\*`
	linkText     = `(\[.*\])(\((http)(?:s)?(\:\/\/).*\))`
	imageFile    = `(\[(?:.*)?\])(\(.*(\.(jpg|png|gif|tiff|bmp))\))`
	listText     = `((\*{1})(\s)(.*)(?:\n)?)+`
)
