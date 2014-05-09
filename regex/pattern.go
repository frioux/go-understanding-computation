package regex

import (
   a "github.com/frioux/go-understanding-computation/automata"
)

type Pattern interface {
   ToNFADesign() a.NFADesign
   Precedence() int
   String() string
}

func Bracket(s Pattern, precedence int) string {
   if s.Precedence() < precedence {
      return "(" + s.String() + ")"
   } else {
      return s.String()
   }
}

func Matches(s Pattern, str string) bool {
   return s.ToNFADesign().DoesAccept(str)
}
