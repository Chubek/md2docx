package process

import (
	"os"
	"log"
	"fmt"
	"md2docx/preprocess"
	"md2docx/parsers"
	"md2docx/adders"
	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
	"github.com/joho/godotenv"
	"md2docx/util"

)

func init() {
	errGTE := godotenv.Load()
  	if errGTE != nil {
    	log.Fatal("Error loading .env file")
  	}

	err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	if err != nil {
		log.Fatal(err)
	}
}

func ProcessAndWrite(filepath string) {
	log.Println("Writing process started. Creating document...")
	doc := document.New()
	defer doc.Close()
	log.Println("Doc created")

	log.Println("Reading file and tokenizing...")
	sentsAndParas, lenPara, lenSent := preprocess.ReadAndSplit(filepath)
	log.Println(fmt.Sprintf("Parsed into %d paragraphs with a total of %d sentences", lenPara, lenSent))

	log.Println("Parsing and adding sentences to document...")

	for _, para := range sentsAndParas {
		paraDoc := document.New().AddParagraph()

		for _, sent := range para {
			runDoc := paraDoc.AddRun()

			adderContainer := adders.AdderContainer{Doc: doc, Para: paraDoc, Run: runDoc}

			if arrParsed := parsers.ParseBlockQuote(sent); !util.IsEmpty(arrParsed) {
				adderContainer.AddBlockQuote(arrParsed)
			}
		}


	}


}