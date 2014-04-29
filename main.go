package main

import "fmt"

type Env map[string]Expr

type Stmt interface {
   is_stmt()
   is_reducible() bool
   reduce(Env) (Stmt, Env)
}

type Expr interface {
   is_expr()
   is_reducible() bool
   reduce(Env) Expr
   Num() int
   Bool() bool
}

type Boolean struct { // {{{
   boolean bool
}

func (s Boolean) is_expr() { }

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

func (s Boolean) reduce(Env) Expr {
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

func (s Number) is_expr() { }

func (s Number) String() string {
   return fmt.Sprintf("%d", s.num)
}

func (s Number) is_reducible() bool {
   return false
}

func (s Number) reduce(Env) Expr {
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
   Left Expr
   Right Expr
}

func (s Add) is_expr() { }

func (s Add) String() string {
   return fmt.Sprintf("%s + %s", s.Left, s.Right)
}

func (s Add) is_reducible() bool {
   return true
}

func (s Add) reduce(environment Env) Expr {
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
   Left Expr
   Right Expr
}

func (s Multiply) is_expr() { }

func (s Multiply) String() string {
   return fmt.Sprintf("%s * %s", s.Left, s.Right)
}

func (s Multiply) is_reducible() bool {
   return true
}

func (s Multiply) reduce(environment Env) Expr {
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
   Left Expr
   Right Expr
}

func (s LessThan) is_expr() { }

func (s LessThan) String() string {
   return fmt.Sprintf("%s < %s", s.Left, s.Right)
}

func (s LessThan) is_reducible() bool {
   return true
}

func (s LessThan) reduce(environment Env) Expr {
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

func (s Variable) is_expr() { }

func (s Variable) String() string {
   return s.name
}

func (s Variable) is_reducible() bool {
   return true
}

func (s Variable) reduce(environment Env) Expr {
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

func (s DoNothing) is_stmt() { }

func (s DoNothing) String() string {
   return "do-nothing"
}

func (s DoNothing) is_reducible() bool {
   return false
}

func (s DoNothing) reduce(environment Env) (Stmt, Env) { // this should never get called
   return s, environment
}

func (s DoNothing) Num() int { // this should never get called
   return -999
}

func (s DoNothing) Bool() bool { // this should never get called
   return false
}

func (s DoNothing) Equal(o Stmt) bool {
   _, ok := o.(DoNothing)
   return ok
}

// }}}

type Assign struct { // {{{
   name string
   expression Expr
}

func (s Assign) is_stmt() { }

func (s Assign) String() string {
   return fmt.Sprintf("%s = %s", s.name, s.expression)
}

func (s Assign) is_reducible() bool {
   return true
}

func (s Assign) reduce(environment Env) (Stmt, Env) {
   if s.expression.is_reducible() {
      return Assign{s.name, s.expression.reduce(environment)}, environment
   } else {
      var new_env Env = Env{}
      for k, v := range environment {
          new_env[k] = v
      }
      new_env[s.name] = s.expression
      return DoNothing{}, new_env
   }
}

// }}}

type Machine struct { // {{{
   statement Stmt
   environment Env
}

func (m *Machine) step() {
   m.statement, m.environment = m.statement.reduce(m.environment)
}

func (m Machine) run() {
   for m.statement.is_reducible() {
      fmt.Println(m.statement)
      fmt.Println(m.environment)
      m.step()
   }
   fmt.Println(m.statement)
   fmt.Println(m.environment)
}

// }}}

func main() {
   Machine{
      Assign{
         "a", Add{Number{2}, Variable{"a"}},
      }, map[string]Expr {
         "a": Number{2},
      },
   }.run()
}

// vim: foldmethod=marker
