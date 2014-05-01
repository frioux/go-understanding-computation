package main

import "fmt"

type FARule struct { // {{{
   state int
   character byte
   next_state int
}

func (s FARule) does_apply_to(state int, character byte) bool {
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

type States []int

type NFARuleBook struct { // {{{
   rules []FARule
}

func (s NFARuleBook) next_states(states States, character byte) States {
   set := make(map[int]bool)
   for x := 0; x < len(states); x++ {
      inner_states := s.follow_rules_for(states[x], character)
      for y := 0; y < len(inner_states); y++ {
         set[inner_states[y]] = true
      }
   }
   ret := States{}
   for k := range set {
      ret = append(ret, k)
   }
   return ret
}

func (s NFARuleBook) follow_rules_for(state int, character byte) States {
   states := s.rules_for(state, character)
   ret := States{}
   for x := 0; x < len(states); x++ {
      ret = append(ret, states[x].follow())
   }
   return ret
}

func (s NFARuleBook) rules_for(state int, character byte) []FARule {
   ret := []FARule{}
   for x := 0; x < len(s.rules); x++ {
      if s.rules[x].does_apply_to(state, character) {
         ret = append(ret, s.rules[x])
      }
   }
   return ret
}

// }}}


type NFA struct { // {{{
   current_states States
   accept_states States
   rulebook NFARuleBook
}

func (s NFA) is_accepting() bool {
   for i := 0; i < len(s.current_states); i++ {
      for j := 0; j < len(s.accept_states); j++ {
         if s.current_states[i] == s.accept_states[j] {
            return true
         }
      }
   }

   return false
}

// }}}

func main() {
   rulebook := NFARuleBook{
      []FARule{
         FARule{1, 'a', 1}, FARule{1, 'b', 1}, FARule{1, 'b', 2},
         FARule{2, 'a', 3}, FARule{2, 'b', 3},
         FARule{3, 'a', 4}, FARule{3, 'b', 4},
      },
   }
   fmt.Println(NFA{States{1}, States{4}, rulebook}.is_accepting())
   fmt.Println(NFA{States{1,2,4}, States{4}, rulebook}.is_accepting())
}

// vim: foldmethod=marker
