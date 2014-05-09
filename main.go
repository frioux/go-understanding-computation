package main

import (
	"fmt"
	a "github.com/frioux/go-understanding-computation/automata"
	"github.com/frioux/go-understanding-computation/stack"
)

func main() {
	rulebook := a.NFARuleBook{
		[]a.FARule{
			{1, 'a', 1}, {1, 'a', 2}, {1, 0, 2},
			{2, 'b', 3},
			{3, 'b', 1}, {3, 0, 2},
		},
	}
	nfa_design := a.NFADesign{1, a.States{3}, rulebook}
	simulation := a.NFASimulation{nfa_design}

	fmt.Println(simulation.NextState(a.States{1, 2}, 'a'))
	fmt.Println(simulation.NextState(a.States{1, 2}, 'b'))
	fmt.Println(simulation.NextState(a.States{3, 2}, 'b'))
	fmt.Println(simulation.NextState(a.States{1, 3, 2}, 'b'))
	fmt.Println(simulation.NextState(a.States{1, 3, 2}, 'a'))

	stack := stack.Stack{'a', 'b', 'c', 'd', 'e'}
	fmt.Println(stack.Peek())
	stack.Pop()
	stack.Pop()
	fmt.Println(stack.Peek())
	stack.Push('x')
	stack.Push('y')
	fmt.Println(stack.Peek())
}

// vim: foldmethod=marker
