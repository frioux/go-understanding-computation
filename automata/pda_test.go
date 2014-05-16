package automata

import (
	"github.com/frioux/go-understanding-computation/stack"
	"testing"
)

func TestPDA(t *testing.T) {
	rule := PDARule{1, '(', 2, '$', []byte{'b', '$'}}
	config := PDAConfiguration{1, stack.Stack{'$'}}

	if !rule.DoesApplyTo(config, '(') {
		t.Errorf("should be true")
	}
	followed := rule.Follow(config)
	testConfig(t, 2, "b$", followed)
}

func TestDPDARulebook(t *testing.T) {
	config := PDAConfiguration{1, stack.Stack{'$'}}

	rulebook := DPDARulebook{
		[]PDARule{
			{1, '(', 2, '$', []byte{'b', '$'}},
			{2, '(', 2, 'b', []byte{'b', 'b'}},
			{2, ')', 2, 'b', []byte{}},
			{2, 0, 1, '$', []byte{'$'}},
		},
	}
	config = rulebook.NextConfiguration(config, '(')
	testConfig(t, 2, "b$", config)
	config = rulebook.NextConfiguration(config, '(')
	testConfig(t, 2, "bb$", config)
	config = rulebook.NextConfiguration(config, ')')
	testConfig(t, 2, "b$", config)
}

func TestDPDA(t *testing.T) {
	rulebook := DPDARulebook{
		[]PDARule{
			{1, '(', 2, '$', []byte{'b', '$'}},
			{2, '(', 2, 'b', []byte{'b', 'b'}},
			{2, ')', 2, 'b', []byte{}},
			{2, 0, 1, '$', []byte{'$'}},
		},
	}
	dpda := DPDA{
		PDAConfiguration{1, stack.Stack{'$'}}, []int{1}, rulebook,
	}

	if !dpda.IsAccepting() {
		t.Errorf("dpda should be accepting")
	}

	dpda.ReadString("(()")

	if dpda.IsAccepting() {
		t.Errorf("dpda should not be accepting")
	}
	testConfig(t, 2, "b$", dpda.CurrentConfiguration)

	config := PDAConfiguration{2, stack.Stack{'$'}}
	config = rulebook.FollowFreeMoves(config)
	if config.State != 1 {
		t.Errorf("config.State should be 1")
	}
	if config.Stack.String() != "Stack «$»" {
		t.Errorf("should be «$», was " + config.Stack.String())
	}
}

func testConfig(t *testing.T, state int, stack string, config PDAConfiguration) {
	if config.State != state {
		t.Errorf("should be ", state, ", was ", config.State)
	}
	if config.Stack.String() != "Stack «"+stack+"»" {
		t.Errorf("should be «b$», was " + config.Stack.String())
	}
}
