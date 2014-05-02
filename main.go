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

type States []int // {{{

func (s States) is_subset_of(other States) bool {
   self_set := make(map[int]bool)
   other_set := make(map[int]bool)
   for i := 0; i < len(s); i++ {
      self_set[s[i]] = true
   }
   for i := 0; i < len(other); i++ {
      other_set[other[i]] = true
   }
   for k := range self_set {
      _, ok := other_set[k]
      if !ok {
         return false
      }
   }
   return true
}

func (s States) union(other States) States {
   set := make(map[int]bool)
   for i := 0; i < len(s); i++ {
      set[s[i]] = true
   }
   for i := 0; i < len(other); i++ {
      set[other[i]] = true
   }
   ret := States{}
   for k := range set {
      ret = append(ret, k)
   }
   return ret
}

// }}}

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

func (s NFARuleBook) follow_free_moves(states States) States {
   more_states := s.next_states(states, 0)

   if more_states.is_subset_of(states) {
      return states
   } else {
      return s.follow_free_moves(states.union(more_states))
   }
}

// }}}

type NFA struct { // {{{
   current_states States
   accept_states States
   rulebook NFARuleBook
}

func (s NFA) CurrentStates() States {
   return s.rulebook.follow_free_moves(s.current_states)
}

func (s NFA) is_accepting() bool {
   for i := 0; i < len(s.CurrentStates()); i++ {
      for j := 0; j < len(s.accept_states); j++ {
         if s.current_states[i] == s.accept_states[j] {
            return true
         }
      }
   }

   return false
}

func (s *NFA) read_character(character byte) {
   s.current_states = s.rulebook.next_states(s.CurrentStates(), character)
}

func (s *NFA) read_string(str string) {
   for i := 0; i < len(str); i++ {
      s.read_character(str[i])
   }
}

// }}}

type NFADesign struct { // {{{
   start_state int
   accept_states States
   rulebook NFARuleBook
}

func (s NFADesign) does_accept(str string) bool {
   nfa := s.to_nfa()
   nfa.read_string(str)
   return nfa.is_accepting()
}

func (s NFADesign) to_nfa() NFA {
   return NFA{States{s.start_state}, s.accept_states, s.rulebook}
}

// }}}

func main() {
   rulebook := NFARuleBook{
      []FARule{
         FARule{1,  0 , 2}, FARule{1, 0, 4},
         FARule{2, 'a', 3},
         FARule{3, 'a', 2},
         FARule{4, 'a', 5},
         FARule{5, 'a', 6},
         FARule{6, 'a', 4},
      },
   }
   nfa_design := NFADesign{1, States{2,4}, rulebook}
   fmt.Println(nfa_design.does_accept("aa"))
   fmt.Println(nfa_design.does_accept("aaa"))
   fmt.Println(nfa_design.does_accept("aaaaa"))
   fmt.Println(nfa_design.does_accept("aaaaaa"))
}

// vim: foldmethod=marker
