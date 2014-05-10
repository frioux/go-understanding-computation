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
	if followed.State != 2 {
		t.Errorf("should be 2, was ", followed.State)
	}
	if followed.Stack.String() != "Stack «b$»" {
		t.Errorf("should be «b$», was " + followed.Stack.String())
	}
}
