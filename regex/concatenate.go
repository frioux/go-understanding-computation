package regex

import (
	a "github.com/frioux/go-understanding-computation/automata"
)

type Concatenate struct {
	first  Pattern
	second Pattern
}

func (s Concatenate) Precedence() int {
	return 1
}

func (s Concatenate) String() string {
	return Bracket(s.first, s.Precedence()) +
		Bracket(s.second, s.Precedence())
}

func (s Concatenate) ToNFADesign() a.NFADesign {
	first_nfa := s.first.ToNFADesign()
	second_nfa := s.second.ToNFADesign()

	start_state := first_nfa.StartState
	accept_states := second_nfa.AcceptStates
	rules := first_nfa.Rulebook.Rules
	for _, v := range second_nfa.Rulebook.Rules {
		rules = append(rules, v)
	}
	for _, v := range first_nfa.AcceptStates {
		rules = append(
			rules,
			a.FARule{v, 0, second_nfa.StartState},
		)
	}

	return a.NFADesign{start_state, accept_states, a.NFARuleBook{rules}}
}
