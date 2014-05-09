package regex

import (
	"fmt"
	a "github.com/frioux/go-understanding-computation/automata"
)

type Literal struct {
	character byte
}

func (s Literal) Precedence() int {
	return 3
}

func (s Literal) String() string {
	return fmt.Sprintf("%c", s.character)
}

func (s Literal) ToNFADesign() a.NFADesign {
	var start_state int = unique_int
	unique_int++
	accept_states := unique_int
	unique_int++
	rulebook := a.NFARuleBook{
		[]a.FARule{a.FARule{start_state, s.character, accept_states}}}

	return a.NFADesign{start_state, a.States{accept_states}, rulebook}
}
