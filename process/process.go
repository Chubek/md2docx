package process

import (
	"fmt"
	"log"
	"md2docx/adders"
	"md2docx/parsers"
	"md2docx/preprocess"
	"md2docx/util"
	"os"

	"github.com/joho/godotenv"
	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
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

func ProcessAndWrite(filepath, savepath string) {
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

			if arrParsed := parsers.ParsePlain(sent); !util.IsEmpty(arrParsed) {
				if res := adderContainer.AddPlain(arrParsed); res == 101 {
					log.Println("Plain text added")
				}
			}

			if arrParsed := parsers.ParseBlockQuote(sent); !util.IsEmpty(arrParsed) {
				if res := adderContainer.AddBlockQuote(arrParsed); res == 101 {
					log.Println("BlockQuote added")
				}
			}

			if arrParsed := parsers.ParseHeader(sent); !util.IsEmpty(arrParsed) {
				if res := adderContainer.AddHeader(arrParsed); res == 101 {
					log.Println("Header added")
				}
			}

			if arrParsed := parsers.ParseTable(sent); len(arrParsed) > 0 {
				if res := adderContainer.AddTable(arrParsed); res == 101 {
					log.Println("Table added")
				}
			}

			if arrParsed := parsers.ParseList(sent); !util.IsEmpty(arrParsed) {
				if res := adderContainer.AddList(arrParsed); res == 101 {
					log.Println("List added")
				}
			}

			if arrParsed := parsers.ParseInlineCode(sent); !util.IsEmpty(arrParsed) {
				if res := adderContainer.AddInlineCode(arrParsed); res == 101 {
					log.Println("InlineCode added")
				}
			}

			if arrParsed := parsers.ParseCodeBlock(sent); !util.IsEmpty(arrParsed) {
				if res := adderContainer.AddCodeBlock(arrParsed); res == 101 {
					log.Println("CodeBlock added")
				}
			}

			if arrParsed := parsers.ParseLink(sent); !util.IsEmpty(arrParsed) {
				if res := adderContainer.AddLink(arrParsed); res == 101 {
					log.Println("Link added")
				}
			}

			if arrParsed := parsers.ParseEmailUrl(sent); !util.IsEmpty(arrParsed) {
				if res := adderContainer.AddEmailUrl(arrParsed); res == 101 {
					log.Println("EmailUrl added")
				}
			}

			if arrParsed := parsers.ParseEmailUrl(sent); !util.IsEmpty(arrParsed) {
				if res := adderContainer.AddEmailUrl(arrParsed); res == 101 {
					log.Println("EmailUrl added")
				}
			}

			if arrParsed := parsers.ParseImage(sent); !util.IsEmpty(arrParsed) {
				if res := adderContainer.AddImage(arrParsed); res == 101 {
					log.Println("Image added")
				}
			}

			if arrParsed := parsers.ParseBold(sent); !util.IsEmpty(arrParsed) {
				if res := adderContainer.AddBold(arrParsed); res == 101 {
					log.Println("Bold text added")
				}
			}

			if arrParsed := parsers.ParseItalic(sent); !util.IsEmpty(arrParsed) {
				if res := adderContainer.AddItalic(arrParsed); res == 101 {
					log.Println("Italic text added")
				}
			}

			if arrParsed := parsers.ParseBoldItalic(sent); !util.IsEmpty(arrParsed) {
				if res := adderContainer.AddBoldItalic(arrParsed); res == 101 {
					log.Println("Bold  Italic text added")
				}
			}

			if arrParsed := parsers.ParseHorizontalLine(sent); !util.IsEmpty(arrParsed) {
				if res := adderContainer.AddHorizLine(arrParsed); res == 101 {
					log.Println("Horizontal line text added")
				}
			}

		}

	}

	finSP := util.SavePathParser(savepath)

	log.Println(fmt.Sprintf("Parsing and adding done... saving file to %s", finSP))

	doc.SaveToFile(finSP)

	log.Println(fmt.Sprintf("File successfully saved to %s", finSP))

}
