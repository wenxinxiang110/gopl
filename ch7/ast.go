//我们的表达式语言由浮点数符号（小数点）；
//二元操作符+，-，*， 和/；
//一元操作符-x和+x；
//调用pow(x,y)，sin(x)，和sqrt(x)的函数；
//例如x和pi的变量；
//当然也有括号和标准的优先级运算符。
package ch7

// An Expr is an arithmetic expression.
type Expr interface {
	Eval(env Env) float64
	Check(vars map[Var]bool) error
}

//!+ast

// A Var identifies a variable, e.g., x.
type Var string

// A literal is a numeric constant, e.g., 3.141.
type literal float64

// A unary represents a unary operator expression, e.g., -x.
type unary struct {
	op rune // one of '+', '-'
	x  Expr
}

// A binary represents a binary operator expression, e.g., x+y.
type binary struct {
	op   rune // one of '+', '-', '*', '/'
	x, y Expr
}

// A call represents a function call expression, e.g., sin(x).
type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}
