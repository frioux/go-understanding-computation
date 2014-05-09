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
