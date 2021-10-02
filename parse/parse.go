package parse

import (
	"md2docx/patterns"
	"md2docx/util"
	"regexp"
	"strings"
)

func ParseBold(pattern *regexp.Regexp, text string) (x int) {
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

func ParseItalic(pattern *regexp.Regexp, text string) (x int) {
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

func ParseBoldItalic(pattern *regexp.Regexp, text string) (x int) {
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

func ParseNormal(pattern *regexp.Regexp, text string) (x int) {
	if pattern.MatchString(text) {
		return 132
	}

	return 100
}

func ParseHeader(pattern *regexp.Regexp, text string) int {
	if pattern.MatchString(text) {
		return strings.Count(text[:6], "#")
	}

	return 100
}
