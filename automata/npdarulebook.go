package automata

type NPDARulebook struct {
	Rules []PDARule
}

func (s NPDARulebook) NextConfigurations(configs PDAConfigurations, char byte) PDAConfigurations {
	ret := PDAConfigurations{}

	for _, x := range configs {
		configs := s.FollowRulesFor(x, char)
		for _, y := range configs {
			ret = append(ret, y)
		}
	}

	return ret
}

func (s NPDARulebook) FollowRulesFor(config PDAConfiguration, char byte) PDAConfigurations {
	rules := s.RulesFor(config, char)

	ret := PDAConfigurations{}
	for _, v := range rules {
		ret = append(ret, v.Follow(config))
	}

	return ret
}

func (s NPDARulebook) RulesFor(config PDAConfiguration, char byte) []PDARule {

	rules := []PDARule{}

	for _, v := range s.Rules {
		if v.DoesApplyTo(config, char) {
			rules = append(rules, v)
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
