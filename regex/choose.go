package regex

import (
	a "github.com/frioux/go-understanding-computation/automata"
)

type Choose struct {
	first  Pattern
	second Pattern
}

func (s Choose) Precedence() int {
	return 0
}

func (s Choose) String() string {
	return Bracket(s.first, s.Precedence()) + "|" +
		Bracket(s.second, s.Precedence())
}

func (s Choose) ToNFADesign() a.NFADesign {
	first_nfa := s.first.ToNFADesign()
	second_nfa := s.second.ToNFADesign()

	// merge accept states
	accept_states := first_nfa.AcceptStates
	for _, v := range second_nfa.AcceptStates {
		accept_states = append(accept_states, v)
	}

	// merge rules
	rules := first_nfa.Rulebook.Rules
	for _, v := range second_nfa.Rulebook.Rules {
		rules = append(rules, v)
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
