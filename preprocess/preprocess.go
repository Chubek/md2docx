package preprocess

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func tokenizeSentence(input string) []string {
	re := regexp.MustCompile(`\.|\?|\!`)
	sep := "|"

	indexes := re.FindAllStringIndex(input, -1)

	move := 0
	for _, v := range indexes {
		p1 := v[0] + move
		p2 := v[1] + move
		input = input[:p1] + sep + input[p1:p2] + sep + input[p2:]
		move += 2
	}

	result := strings.Split(input, sep)

	return result
}

func splitPara(input string) []string {
	patt := regexp.MustCompile(`\n{2}`)

	return patt.Split(input, -1)
}

func ReadAndSplit(mdpath string) ([][]string, int, int) {
	file, err := os.Open(mdpath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	var finalRet [][]string
	var lenPara int
	var lenSent int

	splitPar := splitPara(string(b))

	for _, para := range splitPar {
		para = strings.Trim(para, " ")
		para = strings.Trim(para, "\n")

		splitSent := tokenizeSentence(para)

		var innerSents []string

		for _, sent := range splitSent {
			sent = strings.Trim(sent, " ")
			sent = strings.Trim(sent, "\n")

			innerSents = append(innerSents, sent)

			lenSent += 1
		}

		finalRet = append(finalRet, innerSents)
		lenPara += 1
	}

	return finalRet, lenPara, lenSent
}
