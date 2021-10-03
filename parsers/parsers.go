package parsers

import (
	"md2docx/match"
	"md2docx/patterns"
	"regexp"
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

	for _, txtSplt := range textSplit {
		rowSplit := strings.Split(txtSplt[1:len(txtSplt)-1], "|")
		rowList = append(rowList, rowSplit)
	}

	return rowList

}

func ParseList(text string) []string {
	lstSplt := strings.Split(text, "\n")

	return lstSplt

}