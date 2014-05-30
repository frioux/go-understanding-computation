package lex

import (
	"regexp"
	"strings"
)

type Token int
type Tokens []Token

const (
	If Token = iota
	Else
	While
	DoNothing
	LParen
	RParen
	LBrace
	RBrace
	Semic
	Equals
	Add
	Mult
	LessThan
	Number
	Boolean
	VarName
)

func (s Token) String() string {
	mapping := []string{
		"If",
		"Else",
		"While",
		"DoNothing",
		"LParen",
		"RParen",
		"LBrace",
		"RBrace",
		"Semic",
		"Equals",
		"Add",
		"Mult",
		"LessThan",
		"Number",
		"Boolean",
		"VarName",
	}
	return mapping[s]
}

type rule struct {
	token   Token
	pattern string
}

var grammar []rule

func init() {
	grammar = []rule{
		{If, "if"},
		{Else, "else"},
		{While, "while"},
		{DoNothing, "do-nothing"},
		{LParen, "\\("},
		{RParen, "\\)"},
		{LBrace, "\\{"},
		{RBrace, "\\}"},
		{Semic, ";"},
		{Equals, "="},
		{Add, "\\+"},
		{Mult, "\\*"},
		{LessThan, "<"},
		{Number, "[0-9]+"},
		{Boolean, "true|false"},
		{VarName, "[a-z]+"},
	}
}

type LexicalAnalyzer struct {
	Str string
}

func (s *LexicalAnalyzer) Analyze() Tokens {
	ret := Tokens{}

	for s.hasMoreTokens() {
		ret = append(ret, s.nextToken())
	}

	return ret
}

func (s LexicalAnalyzer) hasMoreTokens() bool {
	return len(s.Str) > 0
}

func (s *LexicalAnalyzer) nextToken() Token {
	rule, match := s.ruleMatching(s.Str)
	s.Str = stringAfter(match)
	return rule.token
}

func (s LexicalAnalyzer) ruleMatching(str string) (rule, string) {
	matches := [][][]byte{}
	rules := []rule{}

	for i := 0; i < len(grammar); i++ {
		match := s.matchAtBeginning(grammar[i].pattern, str)
		if match != nil {
			matches = append(matches, match)
			rules = append(rules, grammar[i])
		}
	}

	return s.ruleWithLongestMatch(rules, matches)
}

func (s LexicalAnalyzer) matchAtBeginning(pattern, str string) [][]byte {
	re := regexp.MustCompile("\\A(" + pattern + ")(.*)")
	match := re.FindSubmatch(([]byte)(str))

	if match != nil {
		match = match[1:]
	}

	return match
}

func (s LexicalAnalyzer) ruleWithLongestMatch(rules []rule, matches [][][]byte) (rule, string) {
	rule := rules[0]
	matchRet := matches[0][1]
	maxLen := len(matches[0][0])
	for i := 0; i < len(matches); i++ {
		newLen := len(matches[i][0])
		if maxLen < newLen {
			rule = rules[i]
			maxLen = newLen
			matchRet = matches[i][1]
		}
	}
	return rule, string(matchRet)
}

func stringAfter(str string) string {
	return strings.TrimLeft(str, " \t\n")
}
