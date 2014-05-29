package automata

import "github.com/frioux/go-understanding-computation/stack"

type NPDADesign struct {
	StartState      int
	BottomCharacter byte
	AcceptStates    States
	Rulebook        NPDARulebook
}

func (s NPDADesign) DoesAccept(str string) bool {
	npda := s.ToNPDA()
	npda.ReadString(str)
	return npda.IsAccepting()
}

func (s NPDADesign) ToNPDA() NPDA {
	startStack := stack.Stack{s.BottomCharacter}
	startConfig := PDAConfiguration{s.StartState, startStack}

	return NPDA{
		PDAConfigurations{startConfig},
		s.AcceptStates, s.Rulebook,
	}
}
