package main

import "fmt"

type Env map[string]Simple

type Simple interface {
   simple()
   is_reducible() bool
   reduce(Env) Simple
   Num() int
   Bool() bool
}

type Boolean struct { // {{{
   boolean bool
}

func (s Boolean) simple() { }

func (s Boolean) String() string {
   if s.boolean {
      return "true"
   } else {
      return "false"
   }
}

func (s Boolean) is_reducible() bool {
   return false
}

func (s Boolean) reduce(Env) Simple {
   return s // this should never get called
}

func (s Boolean) Num() int { // this should never get called
   return -999
}

func (s Boolean) Bool() bool {
   return s.boolean
}

// }}}

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

func (s Number) reduce(Env) Simple {
   return s // this should never get called
}

func (s Number) Num() int {
   return s.num
}

func (s Number) Bool() bool { // this should never get called
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

func (s Add) reduce(environment Env) Simple {
   if s.Left.is_reducible() {
      return Add{s.Left.reduce(environment), s.Right}
   } else if s.Right.is_reducible() {
      return Add{s.Left, s.Right.reduce(environment)}
   } else {
      return Number{s.Left.Num() + s.Right.Num()}
   }
}

func (s Add) Num() int { // this should never get called
   return -999
}

func (s Add) Bool() bool { // this should never get called
   return false
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

func (s Multiply) reduce(environment Env) Simple {
   if s.Left.is_reducible() {
      return Add{s.Left.reduce(environment), s.Right}
   } else if s.Right.is_reducible() {
      return Add{s.Left, s.Right.reduce(environment)}
   } else {
      return Number{s.Left.Num() * s.Right.Num()}
   }
}

func (s Multiply) Num() int { // this should never get called
   return -999
}

func (s Multiply) Bool() bool { // this should never get called
   return false
}

// }}}

type LessThan struct { // {{{
   Left Simple
   Right Simple
}

func (s LessThan) simple() { }

func (s LessThan) String() string {
   return fmt.Sprintf("%s < %s", s.Left, s.Right)
}

func (s LessThan) is_reducible() bool {
   return true
}

func (s LessThan) reduce(environment Env) Simple {
   if s.Left.is_reducible() {
      return LessThan{s.Left.reduce(environment), s.Right}
   } else if s.Right.is_reducible() {
      return LessThan{s.Left, s.Right.reduce(environment)}
   } else {
      if s.Left.Num() < s.Right.Num() {
         return Boolean{true}
      } else {
         return Boolean{false}
      }
   }
}

func (s LessThan) Num() int { // this should never get called
   return -999
}

func (s LessThan) Bool() bool { // this should never get called
   return false
}

// }}}

type Variable struct { // {{{
   name string
}

func (s Variable) simple() { }

func (s Variable) String() string {
   return s.name
}

func (s Variable) is_reducible() bool {
   return true
}

func (s Variable) reduce(environment Env) Simple {
   return environment[s.name]
}

func (s Variable) Num() int { // this should never get called
   return -999
}

func (s Variable) Bool() bool { // this should never get called
   return false
}

// }}}

type DoNothing struct { // {{{
}

func (s DoNothing) simple() { }

func (s DoNothing) String() string {
   return "do-nothing"
}

func (s DoNothing) is_reducible() bool {
   return false
}

func (s DoNothing) reduce(environment Env) Simple { // this should never get called
   return s
}

func (s DoNothing) Num() int { // this should never get called
   return -999
}

func (s DoNothing) Bool() bool { // this should never get called
   return false
}

func (s DoNothing) Equal(o Simple) bool {
   _, ok := o.(DoNothing)
   return ok
}

// }}}

type Machine struct { // {{{
   expression Simple
   environment Env
}

func (m *Machine) step() {
   m.expression = m.expression.reduce(m.environment)
}

func (m Machine) run() {
   for m.expression.is_reducible() {
      fmt.Println(m.expression)
      m.step()
   }
   fmt.Println(m.expression)
}

// }}}

func main() {
   var machine Machine = Machine{Add{
      Multiply{Number{1}, Number{2}},
      Multiply{Number{3}, Number{4}},
   }, map[string]Simple {}};
   machine.run()

   Machine{
      LessThan{
         Multiply{Number{1}, Number{20}},
         Add{Number{100}, Number{-80}},
      }, map[string]Simple {},
   }.run()

   Machine{
      Multiply{
         Variable{"a"}, Number{2},
      }, map[string]Simple {
         "a": Number{2},
      },
   }.run()
}

// vim: foldmethod=marker
