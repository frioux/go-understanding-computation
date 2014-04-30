package main

import (
   "fmt"
)

type Env map[string]Expr

type Stmt interface {
   is_stmt()
   evaluate(Env) Env
}

type Expr interface {
   is_expr()
   evaluate(Env) Expr
   asNum(Env) int
   asBool(Env) bool
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

func (s Boolean) evaluate(Env) Expr {
   return s
}

func (s Boolean) asNum(Env) int {
   return -999 // should never get called
}

func (s Boolean) asBool(Env) bool {
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

func (s Number) evaluate(Env) Expr {
   return s
}

func (s Number) asBool(Env) bool {
   return false // should never get called
}

func (s Number) asNum(Env) int {
   return s.num
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

func (s Add) evaluate(e Env) Expr {
   return Number{
      s.Left.evaluate(e).asNum(e) + s.Right.evaluate(e).asNum(e),
   }
}

func (s Add) asNum(e Env) int {
   return s.evaluate(e).asNum(e)
}

func (s Add) asBool(e Env) bool {
   return s.evaluate(e).asBool(e)
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

func (s Multiply) evaluate(e Env) Expr {
   return Number{s.Left.asNum(e) * s.Right.asNum(e)}
}

func (s Multiply) asNum(e Env) int {
   return s.evaluate(e).asNum(e)
}

func (s Multiply) asBool(e Env) bool {
   return s.evaluate(e).asBool(e)
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

func (s LessThan) evaluate(e Env) Expr {
   return Boolean{s.Left.asNum(e) < s.Right.asNum(e)}
}

func (s LessThan) asNum(e Env) int {
   return s.evaluate(e).asNum(e)
}

func (s LessThan) asBool(e Env) bool {
   return s.evaluate(e).asBool(e)
}

// }}}

type Variable struct { // {{{
   name string
}

func (s Variable) is_expr() { }

func (s Variable) String() string {
   return s.name
}

func (s Variable) evaluate(e Env) Expr {
   return e[s.name]
}

func (s Variable) asNum(e Env) int {
   return s.evaluate(e).asNum(e)
}

func (s Variable) asBool(e Env) bool {
   return s.evaluate(e).asBool(e)
}

// }}}

type DoNothing struct { // {{{
}

func (s DoNothing) is_stmt() { }

func (s DoNothing) String() string {
   return "do-nothing"
}

func (s DoNothing) evaluate(e Env) Env {
   return e
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

func (s Assign) evaluate(e Env) Env {
   var new_env Env = Env{}
   for k, v := range e {
       new_env[k] = v
   }
   new_env[s.name] = s.expression.evaluate(new_env)
   return new_env
}

// }}}

type If struct { // {{{
   expression Expr
   consequence Stmt
   alternative Stmt
}

func (s If) is_stmt() { }

func (s If) String() string {
   return fmt.Sprintf(
      "if (%s) { %s } else { %s }",
      s.expression, s.consequence, s.alternative,
   )
}

func (s If) evaluate(e Env) Env {
   if s.expression.asBool(e) {
      return s.consequence.evaluate(e)
   } else {
      return s.alternative.evaluate(e)
   }
}

// }}}

type Sequence struct { // {{{
   left Stmt
   right Stmt
}

func (s Sequence) is_stmt() { }

func (s Sequence) String() string {
   return fmt.Sprintf(
      "%s; %s",
      s.left, s.right,
   )
}

func (s Sequence) evaluate(e Env) Env {
   return s.right.evaluate(s.left.evaluate(e))
}

// }}}

type While struct { // {{{
   expression Expr
   statement Stmt
}

func (s While) is_stmt() { }

func (s While) String() string {
   return fmt.Sprintf(
      "while (%s) { %s }",
      s.expression, s.statement,
   )
}

func (s While) evaluate(e Env) Env {
   var newEnv Env = e
   for s.expression.asBool(newEnv) {
      newEnv = s.statement.evaluate(newEnv)
   }
   return newEnv
}

// }}}

func main() {
   e := Env{"x": Number{1}}
   e = While{
      LessThan{Variable{"x"}, Number{5}},
      Assign{"x", Multiply{Variable{"x"}, Number{3}}},
   }.evaluate(e)
   fmt.Println(e)
}

// vim: foldmethod=marker
