package main

import (
   "fmt"
)

type Env map[string]Expr

// type Stmt interface {
//    is_stmt()
//    evaluate(Env) Env
// }

type Expr interface {
   is_expr()
   asGo() string
}

type Boolean struct { // {{{
   boolean bool
}

func (s Boolean) is_expr() { }

func (s Boolean) asGo() string {
   return fmt.Sprintf("func(Env) bool { return %t }", s.boolean)
}

// }}}

type Number struct { // {{{
   num int
}

func (s Number) is_expr() { }

func (s Number) asGo() string {
   return fmt.Sprintf("func(Env) int { return %d }", s.num)
}

// }}}

type Add struct { // {{{
   Left Expr
   Right Expr
}

func (s Add) is_expr() { }

func (s Add) asGo() string {
   return fmt.Sprintf(
      "func(e Env) int { return %s(e) + %s(e) }",
      s.Left.asGo(), s.Right.asGo(),
   )
}

// }}}

type Multiply struct { // {{{
   Left Expr
   Right Expr
}

func (s Multiply) is_expr() { }

func (s Multiply) asGo() string {
   return fmt.Sprintf(
      "func(e Env) int { return %s(e) * %s(e) }",
      s.Left.asGo(), s.Right.asGo(),
   )
}

// }}}

type LessThan struct { // {{{
   Left Expr
   Right Expr
}

func (s LessThan) is_expr() { }

func (s LessThan) asGo() string {
   return fmt.Sprintf(
      "func(e Env) bool { return %s(e) < %s(e) }",
      s.Left.asGo(), s.Right.asGo(),
   )
}

// }}}

type Variable struct { // {{{
   name string
}

func (s Variable) is_expr() { }

func (s Variable) asGo() string {
   return fmt.Sprintf(
      "func(e Env) string { return e[\"%s\"] }",
      s.name,
   )
}

// }}}

// type DoNothing struct { // {{{
// }

// func (s DoNothing) is_stmt() { }

// func (s DoNothing) String() string {
//    return "do-nothing"
// }

// func (s DoNothing) evaluate(e Env) Env {
//    return e
// }

// // }}}

// type Assign struct { // {{{
//    name string
//    expression Expr
// }

// func (s Assign) is_stmt() { }

// func (s Assign) String() string {
//    return fmt.Sprintf("%s = %s", s.name, s.expression)
// }

// func (s Assign) evaluate(e Env) Env {
//    var new_env Env = Env{}
//    for k, v := range e {
//        new_env[k] = v
//    }
//    new_env[s.name] = s.expression.evaluate(new_env)
//    return new_env
// }

// // }}}

// type If struct { // {{{
//    expression Expr
//    consequence Stmt
//    alternative Stmt
// }

// func (s If) is_stmt() { }

// func (s If) String() string {
//    return fmt.Sprintf(
//       "if (%s) { %s } else { %s }",
//       s.expression, s.consequence, s.alternative,
//    )
// }

// func (s If) evaluate(e Env) Env {
//    if s.expression.asBool(e) {
//       return s.consequence.evaluate(e)
//    } else {
//       return s.alternative.evaluate(e)
//    }
// }

// // }}}

// type Sequence struct { // {{{
//    left Stmt
//    right Stmt
// }

// func (s Sequence) is_stmt() { }

// func (s Sequence) String() string {
//    return fmt.Sprintf(
//       "%s; %s",
//       s.left, s.right,
//    )
// }

// func (s Sequence) evaluate(e Env) Env {
//    return s.right.evaluate(s.left.evaluate(e))
// }

// // }}}

// type While struct { // {{{
//    expression Expr
//    statement Stmt
// }

// func (s While) is_stmt() { }

// func (s While) String() string {
//    return fmt.Sprintf(
//       "while (%s) { %s }",
//       s.expression, s.statement,
//    )
// }

// func (s While) evaluate(e Env) Env {
//    var newEnv Env = e
//    for s.expression.asBool(newEnv) {
//       newEnv = s.statement.evaluate(newEnv)
//    }
//    return newEnv
// }

// // }}}

func main() {
   fmt.Println(
      "package main\n",
      "import (\"fmt\")\n",
      "type Env map[string]string\n",
      "func main() {\n",
      "f := ", LessThan{Number{2}, Number{3}}.asGo(), "\n",
      "fmt.Println(f(Env{}))\n",
      "}\n",
   )
}

// vim: foldmethod=marker
