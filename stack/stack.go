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

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func (s Stack) String() string {
	return "Stack «" + reverse(string(s)) + "»"
}
