package parsers

import (
	"fmt"
	"log"
	"md2docx/match"
	"md2docx/patterns"
	"md2docx/util"
	"regexp"
	"strings"

	"github.com/dlclark/regexp2"
)



func ParseBlockQuote(text string) []string {
	ret := make([]string, 1)

	if match.MatchAll(text) == 14 {
		ret[0] = text[2:]
	}

	return ret
}

func ParseHorizontalLine(text string) []string {
	ret := make([]string, 1)

	if match.MatchAll(text) == 15 {
		ret[0] = "---"
	}

	return ret

}

func ParseTable(text string) [][]string {
	tableSepPatt := regexp.MustCompile(patterns.TableSep)
	textSplit := tableSepPatt.Split(text, -1)
	var rowList [][]string

	if match.MatchAll(text) != 18 {
		return rowList
	}

	for _, txtSplt := range textSplit {
		rowSplit := strings.Split(txtSplt[1:len(txtSplt)-1], "|")
		rowList = append(rowList, rowSplit)
	}

	return rowList

}

func ParseList(text string) []string {
	if match.MatchAll(text) != 11 {
		return []string{}
	}

	lstSplt := strings.Split(text, "\n")

	return lstSplt

}

func ParseImage(text string) []string {
	if match.MatchAll(text) != 13 {
		return []string{}
	}

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

	patternUrl := regexp2.MustCompile(patterns.Url, 0)

	var imgPath string
	var errDownload error

	var imgHint string

	if len(imgText) >= 1 {
		imgHint = imgText
	} else {
		imgHint = imgTooltip
	}

	if urlMatch, _ := patternUrl.MatchString(imgPathOrUrl); urlMatch {
		imgName := imgPathOrUrl[strings.LastIndex(imgPathOrUrl, "/"):]

		imgPath, errDownload = util.DownloadFile(imgName, imgPathOrUrl)

		if errDownload != nil {
			log.Fatal(fmt.Sprintf("Error downloading image %s", errDownload))
		}

	} else {
		imgPath = imgPathOrUrl
	}

	return []string{imgPath, imgHint}
}

func ParseEmailUrl(text string) []string {
	if match.MatchAll(text) != 12 {
		return []string{}
	}

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

	return []string{emailOrUrl, protocol, explainText}

}

func ParseImageLink(text string) []string {
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

	return []string{"ImageLink", imgPath, toolTip, linkUrl}

}

func ParseLink(text string) []string {
	if match.MatchAll(text) != 10 && match.MatchAll(text) != 13 {
		return []string{}
	}

	if strings.Count(text, "](") > 1 {
		return ParseImageLink(text)
	}

	txtSplit := strings.Split(text, "](")
	linkText := txtSplit[0][1:]
	pattUnnTxt := regexp.MustCompile(`(\"|\')(\w|\W|\S)+(\"|\')`)
	hrefTooltip := pattUnnTxt.FindString(txtSplit[1])
	linkUrl := strings.Trim(pattUnnTxt.ReplaceAllString(txtSplit[1][:len(txtSplit[1])-1], ""), " ")

	return []string{"Normalink", linkText, hrefTooltip, linkUrl}
}

func ParseHeaderOne(text string) []string {
	if match.MatchAll(text) == 1 {
		return []string{"1", text}
	}

	return []string{}
}

func ParseHeaderTwo(text string) []string {
	if match.MatchAll(text) == 2 {
		return []string{"2", text}
	}

	return []string{}
}

func ParseHeaderThree(text string) []string {
	if match.MatchAll(text) == 3 {
		return []string{"3", text}
	}

	return []string{}
}

func ParseHeaderFour(text string) []string {
	if match.MatchAll(text) == 4 {
		return []string{"4", text}
	}

	return []string{}
}

func ParseHeaderFive(text string) []string {
	if match.MatchAll(text) == 5 {
		return []string{"5", text}
	}

	return []string{}
}

func ParseHeaderSix(text string) []string {
	if match.MatchAll(text) == 6 {
		return []string{"6", text}
	}

	return []string{}
}

func ParseCodeBlock(text string) []string {
	if match.MatchAll(text) == 16 {
		return []string{text[4 : len(text)-4]}
	}

	return []string{}
}

func ParseInlineCode(text string) []string {
	if match.MatchAll(text) == 17 {
		pattIC := regexp.MustCompile("(\\`{1})(.*)(\\`{1})")

		inlines := pattIC.FindAllString(text, -1)

		return inlines
	}

	return []string{}
}

func ParseBold(text string) []string {
	if match.MatchAll(text) == 7 {
		pattBold := regexp.MustCompile(`(\_|\*){2}(\w|\W|\d)+(\_|\*){2}`)

		bolds := pattBold.FindAllString(text, -1)

		boldsFin := make([]string, len(bolds))

		for i, bB := range bolds {
			boldsFin[i] = bB[2 : len(bB)-2]
		}

		return boldsFin
	}

	return []string{}
}

func ParseItalic(text string) []string {
	if match.MatchAll(text) == 8 {
		pattItalic := regexp.MustCompile(`(\_|\*){1}(\w|\W|\d)+(\_|\*){1}`)

		italics := pattItalic.FindAllString(text, -1)

		italicsFin := make([]string, len(italics))

		for i, bB := range italics {
			italicsFin[i] = bB[1 : len(bB)-1]
		}

		return italicsFin
	}

	return []string{}
}

func ParseBoldItalic(text string) []string {
	if match.MatchAll(text) == 8 {
		pattBoldItalic := regexp.MustCompile(`(\_|\*){3}(\w|\W|\d)+(\_|\*){3}`)

		boldItalics := pattBoldItalic.FindAllString(text, -1)

		boldItalicsFin := make([]string, len(boldItalics))

		for i, bB := range boldItalics {
			boldItalicsFin[i] = bB[3 : len(bB)-3]
		}

		return boldItalicsFin
	}

	return []string{}
}

