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
	testConfig(t, 2, "b$", dpda.currentConfiguration)

	config := PDAConfiguration{2, stack.Stack{'$'}}
	config = rulebook.FollowFreeMoves(config)
	testConfig(t, 1, "$", config)

	dpda = DPDA{
		PDAConfiguration{1, stack.Stack{'$'}}, []int{1}, rulebook,
	}
	dpda.ReadString("(()(")
	if dpda.IsAccepting() {
		t.Errorf("dpda should not be accepting")
	}
	testConfig(t, 2, "bb$", dpda.CurrentConfiguration())
	dpda.ReadString("))()")
	if !dpda.IsAccepting() {
		t.Errorf("dpda should be accepting")
	}
	testConfig(t, 1, "$", dpda.CurrentConfiguration())

	dpda = DPDA{
		PDAConfiguration{
			1,
			stack.Stack{'$'},
		},
		[]int{1},
		DPDARulebook{
			[]PDARule{
				{1, '(', 2, '$', []byte{'b', '$'}},
				{2, '(', 2, 'b', []byte{'b', 'b'}},
				{2, ')', 2, 'b', []byte{}},
				{2, 0, 1, '$', []byte{'$'}},
			},
		},
	}

	dpda.ReadString("())")
	testConfig(t, StuckState, "$", dpda.CurrentConfiguration())
	if dpda.IsAccepting() {
		t.Errorf("dpda should not be accepting")
	}
	if !dpda.IsStuck() {
		t.Errorf("dpda should be stuck")
	}

	rulebook = DPDARulebook{
		[]PDARule{
			{1, 'a', 2, '$', []byte{'a', '$'}},
			{1, 'b', 2, '$', []byte{'b', '$'}},
			{2, 'a', 2, 'a', []byte{'a', 'a'}},
			{2, 'b', 2, 'b', []byte{'b', 'b'}},
			{2, 'a', 2, 'b', []byte{}},
			{2, 'b', 2, 'a', []byte{}},
			{2, 0, 1, '$', []byte{'$'}},
		},
	}
	dpdaDesign := DPDADesign{1, '$', []int{1}, rulebook}
	testAccept(dpdaDesign, "ababab", true, t)
	testAccept(dpdaDesign, "bbbaaaab", true, t)
	testAccept(dpdaDesign, "baa", false, t)

	rulebook = DPDARulebook{
		[]PDARule{
			{1, 'a', 1, '$', []byte{'a', '$'}},
			{1, 'a', 1, 'a', []byte{'a', 'a'}},
			{1, 'a', 1, 'b', []byte{'a', 'b'}},
			{1, 'b', 1, '$', []byte{'b', '$'}},
			{1, 'b', 1, 'a', []byte{'b', 'a'}},
			{1, 'b', 1, 'b', []byte{'b', 'b'}},
			{1, 'm', 2, '$', []byte{'$'}},
			{1, 'm', 2, 'a', []byte{'a'}},
			{1, 'm', 2, 'b', []byte{'b'}},
			{2, 'a', 2, 'a', []byte{}},
			{2, 'b', 2, 'b', []byte{}},
			{2, 0, 3, '$', []byte{'$'}},
		},
	}
	dpdaDesign = DPDADesign{1, '$', []int{3}, rulebook}
	testAccept(dpdaDesign, "abmba", true, t)
	testAccept(dpdaDesign, "babbamabbab", true, t)
	testAccept(dpdaDesign, "abmb", false, t)
	testAccept(dpdaDesign, "baambaa", false, t)
}

type DoesAccepter interface {
	DoesAccept(string) bool
}

func testAccept(d DoesAccepter, str string, should bool, t *testing.T) {
	if should {
		if !d.DoesAccept(str) {
			t.Errorf("acceptor should accept", str)
		}
	} else {
		if d.DoesAccept(str) {
			t.Errorf("acceptor not should accept", str)
		}
	}
}

func TestDPDADesign(t *testing.T) {
	rulebook := DPDARulebook{
		[]PDARule{
			{1, '(', 2, '$', []byte{'b', '$'}},
			{2, '(', 2, 'b', []byte{'b', 'b'}},
			{2, ')', 2, 'b', []byte{}},
			{2, 0, 1, '$', []byte{'$'}},
		},
	}
	dpdaDesign := DPDADesign{1, '$', []int{1}, rulebook}
	if !dpdaDesign.DoesAccept("(((((((((())))))))))") {
		t.Errorf("dpdaDesign should be accepting")
	}
	if !dpdaDesign.DoesAccept("()(())((()))(()(()))") {
		t.Errorf("dpdaDesign should be accepting")
	}
	if dpdaDesign.DoesAccept("(()(()(()()(()()))()") {
		t.Errorf("dpdaDesign should not be accepting")
	}
	if dpdaDesign.DoesAccept("())") {
		t.Errorf("dpdaDesign should not be accepting")
	}
}

func TestNPDADesign(t *testing.T) {
	rulebook := NPDARulebook{
		[]PDARule{
			{1, 'a', 1, '$', []byte{'a', '$'}},
			{1, 'a', 1, 'a', []byte{'a', 'a'}},
			{1, 'a', 1, 'b', []byte{'a', 'b'}},

			{1, 'b', 1, '$', []byte{'b', '$'}},
			{1, 'b', 1, 'a', []byte{'b', 'a'}},
			{1, 'b', 1, 'b', []byte{'b', 'b'}},

			{1, 0, 2, '$', []byte{'$'}},
			{1, 0, 2, 'a', []byte{'a'}},
			{1, 0, 2, 'b', []byte{'b'}},

			{2, 'a', 2, 'a', []byte{}},
			{2, 'b', 2, 'b', []byte{}},
			{2, 0, 3, '$', []byte{'$'}},
		},
	}
	config := PDAConfiguration{1, stack.Stack{'$'}}
	npda := NPDA{PDAConfigurations{config}, States{3}, rulebook}
	if !npda.IsAccepting() {
		t.Errorf("npda should be accepting")
	}
	testConfigs(t, npda.CurrentConfigurations(), PDAConfigurations{
		{1, stack.Stack{'$'}},
		{2, stack.Stack{'$'}},
		{3, stack.Stack{'$'}},
	})
	npda.ReadString("abb")
	if npda.IsAccepting() {
		t.Errorf("npda should not be accepting")
	}
	testConfigs(t, npda.CurrentConfigurations(), PDAConfigurations{
		{1, stack.Stack{'$', 'a', 'b', 'b'}},
		{2, stack.Stack{'$', 'a'}},
		{2, stack.Stack{'$', 'a', 'b', 'b'}},
	})
	npda.ReadCharacter('a')
	if !npda.IsAccepting() {
		t.Errorf("npda should be accepting")
	}
	testConfigs(t, npda.CurrentConfigurations(), PDAConfigurations{
		{1, stack.Stack{'$', 'a', 'b', 'b', 'a'}},
		{2, stack.Stack{'$'}},
		{2, stack.Stack{'$', 'a', 'b', 'b', 'a'}},
		{3, stack.Stack{'$'}},
	})
}

func testConfigs(t *testing.T, got, expected PDAConfigurations) {
	if !got.
		IsSubsetOf(expected) && len(got) != len(expected) {
		t.Errorf("unexpected configurations")
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
