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
