package main

import "fmt"

var unique_int int = 0

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
   curr := s.CurrentStates()
   for i := 0; i < len(curr); i++ {
      for j := 0; j < len(s.accept_states); j++ {
         if curr[i] == s.accept_states[j] {
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
   nfa := s.to_nfa_default()
   nfa.read_string(str)
   return nfa.is_accepting()
}

func (s NFADesign) to_nfa_default() NFA {
   return s.to_nfa(States{s.start_state})
}

func (s NFADesign) to_nfa(start States) NFA {
   return NFA{start, s.accept_states, s.rulebook}
}

// }}}

type NFASimulation struct { // {{{
   nfa_design NFADesign
}

func (s NFASimulation) next_state(states States, character byte) States {
   nfa := s.nfa_design.to_nfa(states)
   nfa.read_character(character)
   return nfa.CurrentStates()
}

// }}}

type Pattern interface { // {{{
   to_nfa_design() NFADesign
   precedence() int
   String() string
}

func bracket(s Pattern, precedence int) string {
   if s.precedence() < precedence {
      return "(" + s.String() + ")"
   } else {
      return s.String()
   }
}

func matches(s Pattern, str string) bool {
   return s.to_nfa_design().does_accept(str)
}

// }}}

type Empty struct { // {{{
}

func (s Empty) precedence() int {
   return 3
}

func (s Empty) String() string {
   return ""
}

func (s Empty) to_nfa_design() NFADesign {
   var start_state int = unique_int
   unique_int++
   accept_states := []int{start_state}
   rulebook := NFARuleBook{}

   return NFADesign{start_state, accept_states, rulebook}
}

// }}}

type Literal struct { // {{{
   character byte
}

func (s Literal) precedence() int {
   return 3
}

func (s Literal) String() string {
   return fmt.Sprintf("%c", s.character)
}

func (s Literal) to_nfa_design() NFADesign {
   var start_state int = unique_int
   unique_int++
   accept_states := unique_int
   unique_int++
   rulebook := NFARuleBook{
      []FARule{FARule{start_state, s.character, accept_states},
   }}

   return NFADesign{start_state, States{accept_states}, rulebook}
}

// }}}

type Concatenate struct { // {{{
   first Pattern
   second Pattern
}

func (s Concatenate) precedence() int {
   return 1
}

func (s Concatenate) String() string {
   return bracket(s.first, s.precedence()) +
      bracket(s.second, s.precedence())
}

func (s Concatenate) to_nfa_design() NFADesign {
   first_nfa := s.first.to_nfa_design()
   second_nfa := s.second.to_nfa_design()

   start_state := first_nfa.start_state
   accept_states := second_nfa.accept_states
   rules := first_nfa.rulebook.rules
   for i := 0; i < len(second_nfa.rulebook.rules); i++ {
      rules = append(rules, second_nfa.rulebook.rules[i])
   }
   for i := 0; i < len(first_nfa.accept_states); i++ {
      rules = append(
         rules,
         FARule{first_nfa.accept_states[i], 0, second_nfa.start_state},
      )
   }

   return NFADesign{start_state, accept_states, NFARuleBook{rules}}
}

// }}}

type Choose struct { // {{{
   first Pattern
   second Pattern
}

func (s Choose) precedence() int {
   return 0
}

func (s Choose) String() string {
   return bracket(s.first, s.precedence()) + "|" +
      bracket(s.second, s.precedence())
}

func (s Choose) to_nfa_design() NFADesign {
   first_nfa := s.first.to_nfa_design()
   second_nfa := s.second.to_nfa_design()

   // merge accept states
   accept_states := first_nfa.accept_states
   for i := 0; i < len(second_nfa.accept_states); i++ {
      accept_states = append(accept_states, second_nfa.accept_states[i])
   }

   // merge rules
   rules := first_nfa.rulebook.rules
   for i := 0; i < len(second_nfa.rulebook.rules); i++ {
      rules = append(rules, second_nfa.rulebook.rules[i])
   }

   // generate free rules
   var start_state int = unique_int
   unique_int++
   rules = append(
      rules,
      FARule{start_state, 0, first_nfa.start_state},
   )
   rules = append(
      rules,
      FARule{start_state, 0, second_nfa.start_state},
   )

   return NFADesign{start_state, accept_states, NFARuleBook{rules}}
}

// }}}

type Repeat struct { // {{{
   pattern Pattern
}

func (s Repeat) precedence() int {
   return 2
}

func (s Repeat) String() string {
   return bracket(s.pattern, s.precedence()) + "*"
}

func (s Repeat) to_nfa_design() NFADesign {
   nfa := s.pattern.to_nfa_design()
   accept_states := nfa.accept_states
   rules := nfa.rulebook.rules

   // generate accepting start state
   start_state := unique_int
   unique_int++
   accept_states = append(accept_states, start_state)
   rules = append(rules, FARule{start_state, 0, nfa.start_state})

   // generate free moves
   for i := 0; i < len(nfa.accept_states); i++ {
      rules = append(
         rules,
         FARule{nfa.accept_states[i], 0, nfa.start_state},
      )
   }

   return NFADesign{start_state, accept_states, NFARuleBook{rules}}
}

// }}}

func main() {
   rulebook := NFARuleBook{
      []FARule{
         FARule{1, 'a', 1}, FARule{1, 'a', 2}, FARule{1, 0, 2},
         FARule{2, 'b', 3},
         FARule{3, 'b', 1}, FARule{3, 0, 2},
      },
   }
   nfa_design := NFADesign{1, States{3}, rulebook}
   simulation := NFASimulation{nfa_design}

   fmt.Println(simulation.next_state(States{1, 2}, 'a'))
   fmt.Println(simulation.next_state(States{1, 2}, 'b'))
   fmt.Println(simulation.next_state(States{3, 2}, 'b'))
   fmt.Println(simulation.next_state(States{1, 3, 2}, 'b'))
   fmt.Println(simulation.next_state(States{1, 3, 2}, 'a'))
}

// vim: foldmethod=marker
