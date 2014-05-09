package stack

type Stack []byte

func (s *Stack) Push(a byte) {
	*s = append(*s, a)
}

func (s *Stack) Pop() byte {
	end := len(*s) - 1
	ret := (*s)[end]
	*s = (*s)[:end]
	return ret
}

func (s Stack) Peek() byte {
	return s[len(s)-1]
}
