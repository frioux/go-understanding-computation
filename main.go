package main

import "fmt"

type Simple interface {
   simple()
}

type Number struct {
   Num int
}

func (s Number) simple() { }

func (s Number) String() string {
   return fmt.Sprintf("%d", s.Num)
}

type Add struct {
   Left Simple
   Right Simple
}

func (s Add) simple() { }

func (s Add) String() string {
   return fmt.Sprintf("%s + %s", s.Left, s.Right)
}

type Multiply struct {
   Left Simple
   Right Simple
}

func (s Multiply) simple() { }

func (s Multiply) String() string {
   return fmt.Sprintf("%s * %s", s.Left, s.Right)
}

func main() {
   fmt.Println(Add{
      Multiply{Number{1}, Number{2}},
      Multiply{Number{3}, Number{4}},
   })
}
