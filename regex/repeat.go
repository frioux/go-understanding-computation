package regex

import (
   a "github.com/frioux/go-understanding-computation/automata"
)

type Repeat struct {
   pattern Pattern
}

func (s Repeat) Precedence() int {
   return 2
}

func (s Repeat) String() string {
   return Bracket(s.pattern, s.Precedence()) + "*"
}

func (s Repeat) ToNFADesign() a.NFADesign {
   nfa := s.pattern.ToNFADesign()
   accept_states := nfa.AcceptStates
   rules := nfa.Rulebook.Rules

   // generate accepting start state
   start_state := UniqueInt
   UniqueInt++
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
