package lexer

import (
	"fmt"
	"strings"
)

func Start(str []string) {
	for q := 0; q < len(str); q++ {
		var currentString = str[q]

		if currentString != "" {
			var words = strings.Split(currentString, ":")
			if words[0] == "say" {
				args := strings.ReplaceAll(words[1], "\"", "")
				fmt.Println(args)
			}
		}
	}
}
