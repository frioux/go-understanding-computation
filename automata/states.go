package automata

// this should probably be in a more generic set class, probably goset

type States []int

func (s States) IsSubsetOf(other States) bool {
	self_set := make(map[int]bool)
	other_set := make(map[int]bool)
	for _, v := range s {
		self_set[v] = true
	}
	for _, v := range other {
		other_set[v] = true
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
	for _, v := range s {
		set[v] = true
	}
	for _, v := range other {
		set[v] = true
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
	for _, v := range s {
		self_set[v.String()] = v
	}
	for _, v := range other {
		other_set[v.String()] = v
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
	for _, v := range s {
		set[v.String()] = v
	}
	for _, v := range other {
		set[v.String()] = v
	}
	ret := PDAConfigurations{}
	for _, v := range set {
		ret = append(ret, v)
	}
	return ret
}
