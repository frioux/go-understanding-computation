package main

import "fmt"

type Simple interface {
   simple()
   is_reducible() bool
}

type Number struct { // {{{
   Num int
}

func (s Number) simple() { }

func (s Number) String() string {
   return fmt.Sprintf("%d", s.Num)
}

func (s Number) is_reducible() bool {
   return false
}

// }}}

type Add struct { // {{{
   Left Simple
   Right Simple
}

func (s Add) simple() { }

func (s Add) String() string {
   return fmt.Sprintf("%s + %s", s.Left, s.Right)
}

func (s Add) is_reducible() bool {
   return true
}

// }}}

type Multiply struct { // {{{
   Left Simple
   Right Simple
}

func (s Multiply) simple() { }

func (s Multiply) String() string {
   return fmt.Sprintf("%s * %s", s.Left, s.Right)
}

func (s Multiply) is_reducible() bool {
   return true
}

// }}}

func main() {
   fmt.Println(Add{
      Multiply{Number{1}, Number{2}},
      Multiply{Number{3}, Number{4}},
   })
}

// vim: foldmethod=marker
