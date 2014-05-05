package main

import "fmt"

type Statey interface { //
   is_statey()
}

type State int
func (s State) is_statey() { }

func (s State) String() string {
   return fmt.Sprintf("%d", s)
}

// }}}

var unique_int State = 0

type FARule struct { // {{{
   state Statey
   character byte
   next_state Statey
}

func (s FARule) does_apply_to(state Statey, character byte) bool {
   return s.state == state && s.character == character
}

func (s FARule) follow() Statey {
   return s.next_state
}

func (s FARule) String() string {
   return fmt.Sprintf(
      "#<FARule %s --%c--> %s",
      s.state, s.character, s.next_state,
   )
}

// }}}

type States []Statey // {{{
func (s States) is_statey() {}

func (s States) is_subset_of(other States) bool {
   self_set := make(map[Statey]bool)
   other_set := make(map[Statey]bool)
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

func union(s, other []Statey) []Statey {
   set := make(map[Statey]bool)
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

func is_subset_of(s, other []Statey) bool {
   self_set := make(map[Statey]bool)
   other_set := make(map[Statey]bool)
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
   set := make(map[Statey]bool)
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
   set := make(map[Statey]bool)
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

func (s NFARuleBook) follow_rules_for(state Statey, character byte) States {
   states := s.rules_for(state, character)
   ret := States{}
   for x := 0; x < len(states); x++ {
      ret = append(ret, states[x].follow())
   }
   return ret
}

func (s NFARuleBook) rules_for(state Statey, character byte) []FARule {
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

func (s NFARuleBook) alphabet() []byte {
   chars := make(map[byte]struct{})
   for i := 0; i < len(s.rules); i++ {
      char := s.rules[i].character
      if char != 0 {
         chars[char] = struct{}{}
      }
   }
   ret := []byte{}
   for k := range chars {
      ret = append(ret, k)
   }
   return ret
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
   start_state Statey
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

func (s NFADesign) to_nfa(start Statey) NFA {
   return NFA{start.(States), s.accept_states, s.rulebook}
}

// }}}

type NFASimulation struct { // {{{
   nfa_design NFADesign
}

func (s NFASimulation) next_state(state Statey, character byte) Statey {
   nfa := s.nfa_design.to_nfa(state)
   nfa.read_character(character)
   return nfa.CurrentStates()
}

func (s NFASimulation) rules_for(state Statey) []FARule {
   az := s.nfa_design.rulebook.alphabet()
   ret := []FARule{}
   for i := 0; i < len(az); i++ {
      ret = append(ret, FARule{state, az[i], s.next_state(state, az[i])})
   }

   return ret
}

func (s NFASimulation) discover_states_and_rules(states States) (States, []FARule) {

   // map s.rules_for($_), states
   rules := []FARule{}
   for i := 0; i < len(states); i++ {
      to_append := s.rules_for(states[i])
      for j := 0; j < len(to_append); j++ {
         rules = append(rules, to_append[j])
      }
   }

   // uniq, map $_.follow(), rules
   more_states_map := make(map[Statey]struct{})
   for i := 0; i < len(rules); i++ {
      more_states_map[rules[i].follow()] = struct{}{}
   }
   more_states := []Statey{}
   for k := range more_states_map {
      more_states = append(more_states, k)
   }

   if is_subset_of(more_states, states) {
      return states, rules
   } else {
      return s.discover_states_and_rules(union(states, more_states))
   }

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
   var start_state Statey = unique_int
   unique_int++
   accept_states := []Statey{start_state}
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
   var start_state Statey = unique_int
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
   var start_state Statey = unique_int
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
         FARule{State(1), 'a', State(1)}, FARule{State(1), 'a', State(2)}, FARule{State(1), 0, State(2)},
         FARule{State(2), 'b', State(3)},
         FARule{State(3), 'b', State(1)}, FARule{State(3), 0, State(2)},
      },
   }
   nfa_design := NFADesign{State(1), States{State(3)}, rulebook}
   // simulation := NFASimulation{nfa_design}

   start_state := nfa_design.to_nfa_default().CurrentStates()
   fmt.Println(start_state)
}

// vim: foldmethod=marker
