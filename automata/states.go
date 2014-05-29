package automata

// this should probably be in a more generic set class, probably goset

type States []int

func (s States) IsSubsetOf(other States) bool {
	self_set := make(map[int]bool)
	other_set := make(map[int]bool)
	for i := 0; i < len(s); i++ {
		self_set[s[i]] = true
	}
	for i := 0; i < len(other); i++ {
		other_set[other[i]] = true
	}
	for k := range self_set {
		_, ok := other_set[k]
		if !ok {
			return false
		}
	}
	return true
}

func (s States) Union(other States) States {
	set := make(map[int]bool)
	for i := 0; i < len(s); i++ {
		set[s[i]] = true
	}
	for i := 0; i < len(other); i++ {
		set[other[i]] = true
	}
	ret := States{}
	for k := range set {
		ret = append(ret, k)
	}
	return ret
}

type PDAConfigurations []PDAConfiguration

func (s PDAConfigurations) IsSubsetOf(other PDAConfigurations) bool {
	self_set := make(map[string]PDAConfiguration)
	other_set := make(map[string]PDAConfiguration)
	for i := 0; i < len(s); i++ {
		self_set[s[i].String()] = s[i]
	}
	for i := 0; i < len(other); i++ {
		other_set[other[i].String()] = other[i]
	}
	for k := range self_set {
		_, ok := other_set[k]
		if !ok {
			return false
		}
	}
	return true
}

func (s PDAConfigurations) Union(other PDAConfigurations) PDAConfigurations {
	set := make(map[string]PDAConfiguration)
	for i := 0; i < len(s); i++ {
		set[s[i].String()] = s[i]
	}
	for i := 0; i < len(other); i++ {
		set[other[i].String()] = other[i]
	}
	ret := PDAConfigurations{}
	for _, v := range set {
		ret = append(ret, v)
	}
	return ret
}
