package automata

import "testing"

func TestNFA(t *testing.T) {
	nfa := NFA{
		States{1, 2},
		States{2},
		NFARuleBook{
			[]FARule{
				{1, 'a', 2},
			},
		},
	}

	if !nfa.IsAccepting() {
		t.Errorf("should be accepting")
	}
}

