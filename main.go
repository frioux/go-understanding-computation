package main

import "fmt"

type Simple interface {
   simple()
   is_reducible() bool
   reduce() Simple
   Num() int
}

type Number struct { // {{{
   num int
}

func (s Number) simple() { }

func (s Number) String() string {
   return fmt.Sprintf("%d", s.num)
}

func (s Number) is_reducible() bool {
   return false
}

func (s Number) reduce() Simple {
   return s // this should never get called
}

func (s Number) Num() int {
   return s.num
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

func (s Add) reduce() Simple {
   if s.Left.is_reducible() {
      return Add{s.Left.reduce(), s.Right}
   } else if s.Right.is_reducible() {
      return Add{s.Left, s.Right.reduce()}
   } else {
      return Number{s.Left.Num() + s.Right.Num()}
   }
}

func (s Add) Num() int { // this should never get called
   return -999
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

func (s Multiply) reduce() Simple {
   if s.Left.is_reducible() {
      return Add{s.Left.reduce(), s.Right}
   } else if s.Right.is_reducible() {
      return Add{s.Left, s.Right.reduce()}
   } else {
      return Number{s.Left.Num() * s.Right.Num()}
   }
}

func (s Multiply) Num() int { // this should never get called
   return -999
}

// }}}

func main() {
   var expression Simple = Add{
      Multiply{Number{1}, Number{2}},
      Multiply{Number{3}, Number{4}},
   };
   fmt.Println(expression)
   fmt.Println(expression.is_reducible())
   expression = expression.reduce()
   fmt.Println(expression)
   fmt.Println(expression.is_reducible())
   expression = expression.reduce()
   fmt.Println(expression)
   fmt.Println(expression.is_reducible())
   expression = expression.reduce()
   fmt.Println(expression)
   fmt.Println(expression.is_reducible())
}

// vim: foldmethod=marker
