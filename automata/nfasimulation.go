package automata

type NFASimulation struct {
	NFADesign NFADesign
}

func (s NFASimulation) NextState(states States, character byte) States {
	nfa := s.NFADesign.ToNFA(states)
	nfa.ReadCharacter(character)
	return nfa.CurrentStates()
}
