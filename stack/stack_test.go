package stack

import "testing"

func TestStack(t *testing.T) {
	stack := Stack{'$'}
	stack.Push('x')
	stack.Push('y')
	stack.Push('z')

	if stack.Peek() != 'z' {
		t.Errorf("should be 'z'")
	}
	if stack.Pop() != 'z' {
		t.Errorf("should be 'z'")
	}

	if stack.Peek() != 'y' {
		t.Errorf("should be 'y'")
	}
	if stack.Pop() != 'y' {
		t.Errorf("should be 'y'")
	}

	if stack.Peek() != 'x' {
		t.Errorf("should be 'x'")
	}
	if stack.Pop() != 'x' {
		t.Errorf("should be 'x'")
	}
}
