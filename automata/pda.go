package automata

import "github.com/frioux/go-understanding-computation/stack"

type PDAConfiguration struct {
	State int
	Stack stack.Stack
}

type PDARule struct {
	State          int
	Character      byte
	NextState      int
	PopCharacter   byte
	PushCharacters []byte
}

func (s PDARule) DoesApplyTo(config PDAConfiguration, character byte) bool {
	return s.State == config.State &&
		s.PopCharacter == config.Stack.Peek() &&
		s.Character == character
}

func (s PDARule) Follow(config PDAConfiguration) PDAConfiguration {
	return PDAConfiguration{s.NextState, s.NextStack(config)}
}

func (s PDARule) NextStack(config PDAConfiguration) stack.Stack {
	ret := config.Stack
	ret.Pop()
	for i := len(s.PushCharacters) - 1; i > -1; i-- {
		ret.Push(s.PushCharacters[i])
	}
	return ret
}
