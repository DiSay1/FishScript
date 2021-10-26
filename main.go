package main

import (
	"fmt"

	"github.com/disay1/fishscript/src/lexer"
)

func main() {
	var lexems = make(map[int]lexer.LEXEM)
	lexems = lexer.CreateLexems("main.fis")

	fmt.Println("Result:", lexems[1].Value)
}
