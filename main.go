package main

import (
   "fmt"
   "github.com/frioux/go-understanding-computation/stack"
   a "github.com/frioux/go-understanding-computation/automata"
   r "github.com/frioux/go-understanding-computation/regex"
)

type Choose struct { // {{{
   first r.Pattern
   second r.Pattern
}

func (s Choose) Precedence() int {
   return 0
}

func (s Choose) String() string {
   return r.Bracket(s.first, s.Precedence()) + "|" +
      r.Bracket(s.second, s.Precedence())
}

func (s Choose) ToNFADesign() a.NFADesign {
   first_nfa := s.first.ToNFADesign()
   second_nfa := s.second.ToNFADesign()

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
   var start_state int = r.UniqueInt
   r.UniqueInt++
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
   pattern r.Pattern
}

func (s Repeat) Precedence() int {
   return 2
}

func (s Repeat) String() string {
   return r.Bracket(s.pattern, s.Precedence()) + "*"
}

func (s Repeat) ToNFADesign() a.NFADesign {
   nfa := s.pattern.ToNFADesign()
   accept_states := nfa.AcceptStates
   rules := nfa.Rulebook.Rules

   // generate accepting start state
   start_state := r.UniqueInt
   r.UniqueInt++
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
   simulation := a.NFASimulation{nfa_design}

   fmt.Println(simulation.NextState(a.States{1, 2}, 'a'))
   fmt.Println(simulation.NextState(a.States{1, 2}, 'b'))
   fmt.Println(simulation.NextState(a.States{3, 2}, 'b'))
   fmt.Println(simulation.NextState(a.States{1, 3, 2}, 'b'))
   fmt.Println(simulation.NextState(a.States{1, 3, 2}, 'a'))

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
