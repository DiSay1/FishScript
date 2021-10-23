package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/disay1/fishscript/src/lexer"
)

func main() {
	var sourceCode string

	data, err := ioutil.ReadFile("main.fis")
	if err != nil {
		log.Print("Err:", err)
	}

	re := regexp.MustCompile(`\s\s+`)
	sourceCode = re.ReplaceAllString(string(data), " ")
	str := strings.Split(sourceCode, ";")

	for q := 0; q < len(str); q++ {
		str[q] = strings.Trim(str[q], " ")
	}

	lexer.Start(str)
}
