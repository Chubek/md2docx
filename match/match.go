package match

import (
	"fmt"
	"md2docx/patterns"
	"md2docx/util"
	"regexp"
)

type MatchFunction func(pattern *regexp.Regexp, text string) (x int)

type matcherType struct {
	Value   int
	Pattern string
	Matcher MatchFunction
}

func (matcherType matcherType) MatcherPattern(text string) int {
	rePatt := regexp.MustCompile(matcherType.Pattern)

	if matched := matcherType.Matcher(rePatt, text); matched == 134 {
		return matcherType.Value
	} else {
		return 100
	}

}

func matchBold(pattern *regexp.Regexp, text string) (x int) {
	twoUAP := regexp.MustCompile(patterns.TwoUA)

	if pattern.MatchString(text) {
		firstTwo := text[:2]
		lastTwo := text[len(text)-2:]

		if twoUAP.MatchString(firstTwo) && twoUAP.MatchString(lastTwo) {
			if firstTwo != util.Reverse(lastTwo) {
				return 121
			} else {
				return 132
			}
		}
	}

	return 100
}

func matchItalic(pattern *regexp.Regexp, text string) (x int) {
	oneUAP := regexp.MustCompile(patterns.OneUA)

	if pattern.MatchString(text) {
		firstTwo := text[:2]
		lastTwo := text[len(text)-2:]

		if oneUAP.MatchString(firstTwo) && oneUAP.MatchString(lastTwo) {
			if firstTwo != util.Reverse(lastTwo) {
				return 121
			} else {
				return 132
			}
		}
	}

	return 100
}

func matchBoldItalic(pattern *regexp.Regexp, text string) (x int) {
	threeUAP := regexp.MustCompile(patterns.ThreeUA)

	if pattern.MatchString(text) {
		firstTwo := text[:3]
		lastTwo := text[len(text)-3:]

		if threeUAP.MatchString(firstTwo) && threeUAP.MatchString(lastTwo) {
			if firstTwo != util.Reverse(lastTwo) {
				return 121
			} else {
				return 132
			}
		}
	}

	return 100
}

func matchNormal(pattern *regexp.Regexp, text string) (x int) {
	if pattern.MatchString(text) {
		return 132
	}

	return 100
}

func matchHeader(pattern *regexp.Regexp, text string) int {
	if pattern.MatchString(text) {
		return 132
	}

	return 100
}

var (
	headerOnematchr   = matcherType{1, fmt.Sprintf(patterns.HeaderAll, 1), matchHeader}
	headerTwomatchr   = matcherType{2, fmt.Sprintf(patterns.HeaderAll, 2), matchHeader}
	headerThreematchr = matcherType{3, fmt.Sprintf(patterns.HeaderAll, 3), matchHeader}
	headerFourmatchr  = matcherType{4, fmt.Sprintf(patterns.HeaderAll, 4), matchHeader}
	headerFivematchr  = matcherType{5, fmt.Sprintf(patterns.HeaderAll, 5), matchHeader}
	headerSixmatchr   = matcherType{6, fmt.Sprintf(patterns.HeaderAll, 6), matchHeader}
	boldmatchr        = matcherType{7, patterns.BoldItalicText, matchBold}
	italicmatchr      = matcherType{8, patterns.BoldItalicText, matchItalic}
	boldItalicmatchr  = matcherType{9, patterns.BoldItalicText, matchBoldItalic}
	linkmatchr        = matcherType{10, patterns.LinkText, matchNormal}
	listmatchr        = matcherType{11, patterns.ListText, matchNormal}
	emailUrlmatchr    = matcherType{12, patterns.ListText, matchNormal}
	imagematchr       = matcherType{13, patterns.ImageFile, matchNormal}
	bQmatchr          = matcherType{14, patterns.BlockQuote, matchNormal}
	horizLinematchr   = matcherType{15, patterns.LinkText, matchNormal}
	codeBlockmatchr   = matcherType{16, patterns.CodeBlock, matchNormal}
	inlineCodematchr  = matcherType{17, patterns.InlineCode, matchNormal}
)

func MatchAll(text string) int {
	allTypes := []matcherType{headerOnematchr,
		headerTwomatchr,
		headerThreematchr,
		headerFourmatchr,
		headerFivematchr,
		headerSixmatchr,
		boldItalicmatchr,
		italicmatchr,
		boldmatchr,
		linkmatchr,
		listmatchr,
		emailUrlmatchr,
		imagematchr,
		bQmatchr,
		horizLinematchr,
		codeBlockmatchr,
		inlineCodematchr}

	for _, pT := range allTypes {
		if pTRes := pT.MatcherPattern(text); pTRes != 100 {
			return pTRes
		}
	}

	return -1
}
