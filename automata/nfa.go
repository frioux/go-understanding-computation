package automata

type NFA struct {
	CurrentStates_ States
	AcceptStates   States
	Rulebook       NFARuleBook
}

func (s NFA) CurrentStates() States {
	return s.Rulebook.FollowFreeMoves(s.CurrentStates_)
}

func (s NFA) IsAccepting() bool {
	curr := s.CurrentStates()
	for i := 0; i < len(curr); i++ {
		for j := 0; j < len(s.AcceptStates); j++ {
			if curr[i] == s.AcceptStates[j] {
				return true
			}
		}
	}

	return false
}

func (s *NFA) ReadCharacter(character byte) {
	s.CurrentStates_ = s.Rulebook.NextStates(s.CurrentStates(), character)
}

func (s *NFA) ReadString(str string) {
	for i := 0; i < len(str); i++ {
		s.ReadCharacter(str[i])
	}
}
