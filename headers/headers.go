package headers

import (
	"fmt"
	"md2docx/patterns"
	"regexp"
	"strings"
	"github.com/dlclark/regexp2"
	"github.com/unidoc/unioffice/document"
)


func AddHeader(text string, para document.Paragraph, level int) int {
	run := para.AddRun()
	style := fmt.Sprintf("Heading%d", level)
	para.SetStyle(style)
	text = text[2:]

	linkPattern := regexp2.MustCompile(patterns.Url,0)

	if isMatch, _ := linkPattern.MatchString(text); isMatch {
		return addLinkHeader(text, para, level)
	}

	run.AddText(text)

	return 101
}

func addLinkHeader(text string, para document.Paragraph, level int) int {
	txtSplit := strings.Split(text, "](")
	linkText := txtSplit[0][1:]
	pattUnnTxt := regexp.MustCompile(`(\"|\')(\w|\W|\S)+(\"|\')`)
	hrefTooltip := pattUnnTxt.FindString(txtSplit[1])
	linkUrl := strings.Trim(pattUnnTxt.ReplaceAllString(txtSplit[1][:len(txtSplit[1])-1], ""), " ")

	hl := para.AddHyperLink()
	hl.SetTarget(linkUrl)
	run := hl.AddRun()
	style := fmt.Sprintf("Heading%d", level)
	run.Properties().SetStyle(style)
	run.AddText(linkText)
	hl.SetToolTip(hrefTooltip)

	return 101
}