package stack

type Stack []byte

func (s *Stack) push(a byte) {
   *s = append(*s, a)
}

func (s *Stack) pop() byte {
   end := len(*s) - 1;
   ret := (*s)[end]
   *s = (*s)[:end]
   return ret
}

func (s Stack) peek() byte {
   return s[len(s)-1]
}
