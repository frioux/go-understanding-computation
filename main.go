package main

import (
	"fmt"
	"github.com/frioux/go-understanding-computation/lex"
)

func main() {
	lexer := lex.LexicalAnalyzer{"y = x * 7"}
	fmt.Println(lexer.Analyze())
}

// vim: foldmethod=marker
