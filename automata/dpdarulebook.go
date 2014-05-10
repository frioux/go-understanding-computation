package automata

type DPDARulebook struct {
	Rules []PDARule
}

func (s DPDARulebook) NextConfiguration(config PDAConfiguration, char byte) PDAConfiguration {
	return s.RuleFor(config, char).Follow(config)
}

func (s DPDARulebook) RuleFor(config PDAConfiguration, char byte) PDARule {

	for i := 0; i < len(s.Rules); i++ {
		if s.Rules[i].DoesApplyTo(config, char) {
			return s.Rules[i]
		}
	}

	return PDARule{0, 0, 0, 0, []byte{}}
}
