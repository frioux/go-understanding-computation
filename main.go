package main

import (
   "fmt"
   "github.com/frioux/go-understanding-computation/stack"
   a "github.com/frioux/go-understanding-computation/automata"
)

var unique_int int = 0

type NFASimulation struct { // {{{
   nfa_design a.NFADesign
}

func (s NFASimulation) next_state(states a.States, character byte) a.States {
   nfa := s.nfa_design.ToNFA(states)
   nfa.ReadCharacter(character)
   return nfa.CurrentStates()
}

// }}}

type Pattern interface { // {{{
   to_nfa_design() a.NFADesign
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
   return s.to_nfa_design().DoesAccept(str)
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

func (s Empty) to_nfa_design() a.NFADesign {
   var start_state int = unique_int
   unique_int++
   accept_states := []int{start_state}
   rulebook := a.NFARuleBook{}

   return a.NFADesign{start_state, accept_states, rulebook}
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

func (s Literal) to_nfa_design() a.NFADesign {
   var start_state int = unique_int
   unique_int++
   accept_states := unique_int
   unique_int++
   rulebook := a.NFARuleBook{
      []a.FARule{a.FARule{start_state, s.character, accept_states},
   }}

   return a.NFADesign{start_state, a.States{accept_states}, rulebook}
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

func (s Concatenate) to_nfa_design() a.NFADesign {
   first_nfa := s.first.to_nfa_design()
   second_nfa := s.second.to_nfa_design()

   start_state := first_nfa.StartState
   accept_states := second_nfa.AcceptStates
   rules := first_nfa.Rulebook.Rules
   for i := 0; i < len(second_nfa.Rulebook.Rules); i++ {
      rules = append(rules, second_nfa.Rulebook.Rules[i])
   }
   for i := 0; i < len(first_nfa.AcceptStates); i++ {
      rules = append(
         rules,
         a.FARule{first_nfa.AcceptStates[i], 0, second_nfa.StartState},
      )
   }

   return a.NFADesign{start_state, accept_states, a.NFARuleBook{rules}}
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

func (s Choose) to_nfa_design() a.NFADesign {
   first_nfa := s.first.to_nfa_design()
   second_nfa := s.second.to_nfa_design()

   // merge accept states
   accept_states := first_nfa.AcceptStates
   for i := 0; i < len(second_nfa.AcceptStates); i++ {
      accept_states = append(accept_states, second_nfa.AcceptStates[i])
   }

   // merge rules
   rules := first_nfa.Rulebook.Rules
   for i := 0; i < len(second_nfa.Rulebook.Rules); i++ {
      rules = append(rules, second_nfa.Rulebook.Rules[i])
   }

   // generate free rules
   var start_state int = unique_int
   unique_int++
   rules = append(
      rules,
      a.FARule{start_state, 0, first_nfa.StartState},
   )
   rules = append(
      rules,
      a.FARule{start_state, 0, second_nfa.StartState},
   )

   return a.NFADesign{start_state, accept_states, a.NFARuleBook{rules}}
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

func (s Repeat) to_nfa_design() a.NFADesign {
   nfa := s.pattern.to_nfa_design()
   accept_states := nfa.AcceptStates
   rules := nfa.Rulebook.Rules

   // generate accepting start state
   start_state := unique_int
   unique_int++
   accept_states = append(accept_states, start_state)
   rules = append(rules, a.FARule{start_state, 0, nfa.StartState})

   // generate free moves
   for i := 0; i < len(nfa.AcceptStates); i++ {
      rules = append(
         rules,
         a.FARule{nfa.AcceptStates[i], 0, nfa.StartState},
      )
   }

   return a.NFADesign{start_state, accept_states, a.NFARuleBook{rules}}
}

// }}}

func main() {
   rulebook := a.NFARuleBook{
      []a.FARule{
         a.FARule{1, 'a', 1}, a.FARule{1, 'a', 2}, a.FARule{1, 0, 2},
         a.FARule{2, 'b', 3},
         a.FARule{3, 'b', 1}, a.FARule{3, 0, 2},
      },
   }
   nfa_design := a.NFADesign{1, a.States{3}, rulebook}
   simulation := NFASimulation{nfa_design}

   fmt.Println(simulation.next_state(a.States{1, 2}, 'a'))
   fmt.Println(simulation.next_state(a.States{1, 2}, 'b'))
   fmt.Println(simulation.next_state(a.States{3, 2}, 'b'))
   fmt.Println(simulation.next_state(a.States{1, 3, 2}, 'b'))
   fmt.Println(simulation.next_state(a.States{1, 3, 2}, 'a'))

   stack := stack.Stack{'a', 'b', 'c', 'd', 'e'}
   fmt.Println(stack.Peek())
   stack.Pop()
   stack.Pop()
   fmt.Println(stack.Peek())
   stack.Push('x')
   stack.Push('y')
   fmt.Println(stack.Peek())
}

// vim: foldmethod=marker
