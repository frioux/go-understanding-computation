package main

import "fmt"

type Simple interface {
   simple()
}

type Number struct {
   Num int
}

func (s Number) simple() { }

type Add struct {
   Left Simple
   Right Simple
}

func (s Add) simple() { }

type Multiply struct {
   Left Simple
   Right Simple
}

func (s Multiply) simple() { }

func main() {
   fmt.Println(Add{
      Multiply{Number{1}, Number{2}},
      Multiply{Number{3}, Number{4}},
   })
}
