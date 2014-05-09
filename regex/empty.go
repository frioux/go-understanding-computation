package regex

import (
   a "github.com/frioux/go-understanding-computation/automata"
)
type Empty struct {
}

func (s Empty) Precedence() int {
   return 3
}

func (s Empty) String() string {
   return ""
}

func (s Empty) ToNFADesign() a.NFADesign {
   var start_state int = UniqueInt
   UniqueInt++
   accept_states := []int{start_state}
   rulebook := a.NFARuleBook{}

   return a.NFADesign{start_state, accept_states, rulebook}
}
