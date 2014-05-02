package main

import "fmt"

type Pattern interface {
   bracket(int) string
}

type Empty struct { // {{{
}

func (s Empty) precedence() int {
   return 3
}

func (s Empty) bracket(outer_precedence int) string {
   if s.precedence() < outer_precedence {
      return "(" + s.String() + ")"
   } else {
      return s.String()
   }
}

func (s Empty) String() string {
   return ""
}

// }}}

type Literal struct { // {{{
   character byte
}

func (s Literal) precedence() int {
   return 3
}

func (s Literal) bracket(outer_precedence int) string {
   if s.precedence() < outer_precedence {
      return "(" + s.String() + ")"
   } else {
      return s.String()
   }
}

func (s Literal) String() string {
   return fmt.Sprintf("%c", s.character)
}

// }}}

type Concatenate struct { // {{{
   first Pattern
   second Pattern
}

func (s Concatenate) precedence() int {
   return 1
}

func (s Concatenate) bracket(outer_precedence int) string {
   if s.precedence() < outer_precedence {
      return "(" + s.String() + ")"
   } else {
      return s.String()
   }
}

func (s Concatenate) String() string {
   return s.first.bracket(s.precedence()) +
      s.second.bracket(s.precedence())
}

// }}}

type Choose struct { // {{{
   first Pattern
   second Pattern
}

func (s Choose) precedence() int {
   return 0
}

func (s Choose) bracket(outer_precedence int) string {
   if s.precedence() < outer_precedence {
      return "(" + s.String() + ")"
   } else {
      return s.String()
   }
}

func (s Choose) String() string {
   return s.first.bracket(s.precedence()) + "|" +
      s.second.bracket(s.precedence())
}

// }}}

type Repeat struct { // {{{
   pattern Pattern
}

func (s Repeat) precedence() int {
   return 2
}

func (s Repeat) bracket(outer_precedence int) string {
   if s.precedence() < outer_precedence {
      return "(" + s.String() + ")"
   } else {
      return s.String()
   }
}

func (s Repeat) String() string {
   return s.pattern.bracket(s.precedence()) + "*"
}

// }}}

func main() {
   fmt.Println(Repeat{
      Choose{
         Concatenate{Literal{'a'}, Literal{'b'}},
         Literal{'a'},
      },
   })
}

// vim: foldmethod=marker
