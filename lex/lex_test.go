package lex

import "testing"

func TestLexicalAnalyzer(t *testing.T) {
	lexer := LexicalAnalyzer{"y = x * 7"}
	testSyntax(t, lexer.Analyze(), Tokens{VarName, Equals, VarName, Mult, Number})

	lexer = LexicalAnalyzer{"while (x < 5) { x = x * 3 }"}
	testSyntax(t, lexer.Analyze(), Tokens{
		While, LParen, VarName, LessThan, Number, RParen,
		LBrace, VarName, Equals, VarName, Mult, Number, RBrace,
	})

	lexer = LexicalAnalyzer{"if (x < 10) { y = true; x = 0 } else { do-nothing }"}
	testSyntax(t, lexer.Analyze(), Tokens{
		If, LParen, VarName, LessThan, Number, RParen, LBrace,
			VarName, Equals, Boolean, Semic,
			VarName, Equals, Number,
	   RBrace, Else, LBrace,
			DoNothing,
		RBrace,
	})
}

func testSyntax(t *testing.T, got, expected Tokens) {
	gotLen, expectedLen := len(got), len(expected)
	if gotLen != expectedLen {
		t.Errorf("got", len(got), "tokens and expected", len(expected), "tokens")
		return
	}
	for i, v := range got {
		if v != expected[i] {
			t.Errorf("got["+string(i)+"]:", v, "expected["+string(i)+"]:", expected[i])
			return
		}
	}
}
