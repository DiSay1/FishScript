package lexer

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func CreateLexems(filePath string) map[int]LEXEM {
	lexems := make(map[int]LEXEM)

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("Err:", err)
	}

	re := regexp.MustCompile(`\s\s+`)
	sourceCode := re.ReplaceAllString(string(data), " ")
	str := strings.Split(sourceCode, ";")

	for q := 0; q < len(str); q++ {
		str[q] = strings.Trim(str[q], " ")
	}

	for q := 0; q < len(str); q++ {
		var currentString = str[q]
		var words []string

		matched, err := regexp.MatchString(`use.`, currentString)
		if err != nil {
			fmt.Println("Err:", err)
			break
		}
		if matched {
			words = strings.Split(currentString, " ")
		}

		matched, err = regexp.MatchString(`say.`, currentString)
		if err != nil {
			fmt.Println("Err:", err)
			break
		}
		if matched {
			words = strings.Split(currentString, ":")
		}

		for i := 0; i < len(words); i++ {
			words[i] = strings.Trim(words[i], " ")
		}

		e := 1
		for p := 0; p < len(words); p++ {
			if words[p] == "say" {
				res := strings.Trim(words[p+1], "\"")
				lexems[e] = LEXEM{
					Key:   "say",
					Value: res,
				}
			}
		}
	}
	return lexems
}
