package automata

type NPDARulebook struct {
	Rules []PDARule
}

func (s NPDARulebook) NextConfigurations(configs PDAConfigurations, char byte) PDAConfigurations {
	ret := PDAConfigurations{}

	for i := 0; i < len(configs); i++ {
		configs := s.FollowRulesFor(configs[i], char)
		for j := 0; j < len(configs); j++ {
			ret = append(ret, configs[j])
		}
	}

	return ret
}

func (s NPDARulebook) FollowRulesFor(config PDAConfiguration, char byte) PDAConfigurations {
	rules := s.RulesFor(config, char)

	ret := PDAConfigurations{}
	for i := 0; i < len(rules); i++ {
		ret = append(ret, rules[i].Follow(config))
	}

	return ret
}

func (s NPDARulebook) RulesFor(config PDAConfiguration, char byte) []PDARule {

	rules := []PDARule{}

	for i := 0; i < len(s.Rules); i++ {
		if s.Rules[i].DoesApplyTo(config, char) {
			rules = append(rules, s.Rules[i])
		}
	}

	return rules
}

func (s NPDARulebook) FollowFreeMoves(configs PDAConfigurations) PDAConfigurations {
	moreConfigs := s.NextConfigurations(configs, 0)

	if moreConfigs.IsSubsetOf(configs) {
		return configs
	} else {
		return s.FollowFreeMoves(configs.Union(moreConfigs))
	}
}
