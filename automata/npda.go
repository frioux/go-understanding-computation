package automata

type NPDA struct {
	currentConfigurations PDAConfigurations
	AcceptStates States
	Rulebook NPDARulebook
}

func (s NPDA) IsAccepting() bool {
	configs := s.CurrentConfigurations()
	for i := 0; i < len(configs); i++ {
		for j := 0; j < len(s.AcceptStates); j++ {
			if configs[i].State == s.AcceptStates[j] {
				return true
			}
		}
	}
	return false
}

func (s *NPDA) ReadCharacter(char byte) {
	s.currentConfigurations = s.Rulebook.NextConfigurations(s.CurrentConfigurations(), char)
}

func (s *NPDA) ReadString(str string) {
	for i := 0; i < len(str); i++ {
		s.ReadCharacter(str[i])
	}
}

func (s NPDA) CurrentConfigurations() PDAConfigurations {
	return s.Rulebook.FollowFreeMoves(s.currentConfigurations)
}
