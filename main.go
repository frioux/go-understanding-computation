package main

import "fmt"

type Expr interface {
   is_expr()
   asGo() string
}

type FARule struct { // {{{
   state int
   character rune
   next_state int
}

func (s FARule) does_apply_to(state int, character rune) bool {
   return s.state == state && s.character == character
}

func (s FARule) follow() int {
   return s.next_state
}

func (s FARule) String() string {
   return fmt.Sprintf(
      "#<FARule %s -- %c--> %s",
      s.state, s.character, s.next_state,
   )
}

// }}}

type DFARuleBook struct { // {{{
   rules []FARule
}

func (s DFARuleBook) next_state(state int, character rune) int {
   return s.rule_for(state, character).follow()
}

func (s DFARuleBook) rule_for(state int, character rune) FARule {
   for i := 0; i < len(s.rules); i++ {
      if s.rules[i].does_apply_to(state, character) {
         return s.rules[i]
      }
   }
   return FARule{1, '💀', 1}
}

// }}}

func main() {
   rulebook := DFARuleBook{
      []FARule{
         FARule{1, 'a', 2}, FARule{1, 'b', 1},
         FARule{2, 'a', 2}, FARule{2, 'b', 3},
         FARule{3, 'a', 3}, FARule{3, 'b', 3},
      },
   }
   fmt.Println(rulebook.next_state(1, 'a'))
   fmt.Println(rulebook.next_state(1, 'b'))
   fmt.Println(rulebook.next_state(2, 'b'))
}

// vim: foldmethod=marker
