package automata

import "github.com/frioux/go-understanding-computation/stack"

type DPDADesign struct {
	StartState      int
	BottomCharacter byte
	AcceptStates    []int
	Rulebook        DPDARulebook
}

func (s DPDADesign) DoesAccept(str string) bool {
	dpda := s.ToDPDA()
	dpda.ReadString(str)
	return dpda.IsAccepting()
}

func (s DPDADesign) ToDPDA() DPDA {
	startStack := stack.Stack{s.BottomCharacter}
	startConfig := PDAConfiguration{s.StartState, startStack}
	return DPDA{startConfig, s.AcceptStates, s.Rulebook}
}
