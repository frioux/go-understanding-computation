package automata

import "fmt"

type FARule struct {
	State     int
	Character byte
	NextState int
}

func (s FARule) DoesApplyTo(state int, character byte) bool {
	return s.State == state && s.Character == character
}

func (s FARule) Follow() int {
	return s.NextState
}

func (s FARule) String() string {
	return fmt.Sprintf(
		"#<FARule %s -- %c--> %s",
		s.State, s.Character, s.NextState,
	)
}
