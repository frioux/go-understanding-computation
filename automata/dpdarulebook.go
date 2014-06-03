package automata

type DPDARulebook struct {
	Rules []PDARule
}

func (s DPDARulebook) NextConfiguration(config PDAConfiguration, char byte) PDAConfiguration {
	return s.RuleFor(config, char).Follow(config)
}

func (s DPDARulebook) RuleFor(config PDAConfiguration, char byte) PDARule {

	for _, v := range s.Rules {
		if v.DoesApplyTo(config, char) {
			return v
		}
	}

	return PDARule{0, 0, 0, 0, []byte{}}
}

func (s DPDARulebook) DoesApplyTo(config PDAConfiguration, char byte) bool {
	r := s.RuleFor(config, char)
	return !(r.State == 0 &&
		r.Character == 0 &&
		r.NextState == 0 &&
		r.PopCharacter == 0 &&
		len(r.PushCharacters) == 0)
}

func (s DPDARulebook) FollowFreeMoves(config PDAConfiguration) PDAConfiguration {
	if s.DoesApplyTo(config, 0) {
		return s.FollowFreeMoves(s.NextConfiguration(config, 0))
	} else {
		return config
	}
}
