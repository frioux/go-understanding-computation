package automata

type DPDA struct {
	currentConfiguration PDAConfiguration
	AcceptStates         []int
	Rulebook             DPDARulebook
}

func (s DPDA) IsAccepting() bool {
	for i := 0; i < len(s.AcceptStates); i++ {
		if s.AcceptStates[i] == s.CurrentConfiguration().State {
			return true
		}
	}
	return false
}

func (s *DPDA) ReadCharacter(char byte) {
	s.currentConfiguration = s.NextConfiguration(char)
}

func (s *DPDA) ReadString(str string) {
	for i := 0; !s.IsStuck() && i < len(str); i++ {
		s.ReadCharacter(str[i])
	}
}

func (s *DPDA) CurrentConfiguration() PDAConfiguration {
	return s.Rulebook.FollowFreeMoves(s.currentConfiguration)
}

func (s DPDA) NextConfiguration(char byte) PDAConfiguration {
	if s.Rulebook.DoesApplyTo(s.CurrentConfiguration(), char) {
		return s.Rulebook.NextConfiguration(s.CurrentConfiguration(), char)
	} else {
		return s.CurrentConfiguration().Stuck()
	}
}

func (s DPDA) IsStuck() bool {
	return s.CurrentConfiguration().IsStuck()
}
