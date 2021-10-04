package parsers

import (
	"fmt"
	"log"
	"md2docx/match"
	"md2docx/patterns"
	"md2docx/util"
	"regexp"
	"github.com/dlclark/regexp2"
	"strings"
)

func ParseBlockQuote(text string, _ int) []string {
	ret := make([]string, 1)

	if match.MatchAll(text) == 14 {
		ret[0] = text[2:]
	}

	return ret
}

func ParseHorizontalLine(text string, _ int) []string {
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

func ParseImageLink(text string) []string{
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