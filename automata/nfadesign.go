package automata

type NFADesign struct {
   StartState int
   AcceptStates States
   Rulebook NFARuleBook
}

func (s NFADesign) DoesAccept(str string) bool {
   nfa := s.ToNFADefault()
   nfa.ReadString(str)
   return nfa.IsAccepting()
}

func (s NFADesign) ToNFADefault() NFA {
   return s.ToNFA(States{s.StartState})
}

func (s NFADesign) ToNFA(start States) NFA {
   return NFA{start, s.AcceptStates, s.Rulebook}
}
