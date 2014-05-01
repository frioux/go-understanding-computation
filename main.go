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

type DFA struct { // {{{
   current_state int
   accept_states []int
   rulebook DFARuleBook
}

func (s DFA) is_accepting() bool {
   for x := 0; x < len(s.accept_states); x++ {
      if s.current_state == s.accept_states[x] {
         return true
      }
   }
   return false
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
   fmt.Println(DFA{1, []int {1, 3}, rulebook}.is_accepting())
   fmt.Println(DFA{1, []int {3}, rulebook}.is_accepting())
}

// vim: foldmethod=marker
