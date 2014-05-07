package automata

type NFARuleBook struct {
   Rules []FARule
}

func (s NFARuleBook) NextStates(states States, character byte) States {
   set := make(map[int]bool)
   for x := 0; x < len(states); x++ {
      inner_states := s.FollowRulesFor(states[x], character)
      for y := 0; y < len(inner_states); y++ {
         set[inner_states[y]] = true
      }
   }
   ret := States{}
   for k := range set {
      ret = append(ret, k)
   }
   return ret
}

func (s NFARuleBook) FollowRulesFor(state int, character byte) States {
   states := s.RulesFor(state, character)
   ret := States{}
   for x := 0; x < len(states); x++ {
      ret = append(ret, states[x].Follow())
   }
   return ret
}

func (s NFARuleBook) RulesFor(state int, character byte) []FARule {
   ret := []FARule{}
   for x := 0; x < len(s.Rules); x++ {
      if s.Rules[x].DoesApplyTo(state, character) {
         ret = append(ret, s.Rules[x])
      }
   }
   return ret
}

func (s NFARuleBook) FollowFreeMoves(states States) States {
   more_states := s.NextStates(states, 0)

   if more_states.IsSubsetOf(states) {
      return states
   } else {
      return s.FollowFreeMoves(states.Union(more_states))
   }
}
