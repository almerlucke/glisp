package functions

import (
	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
)

// BuildinFunctionImp build in function implementation
type BuildinFunctionImp func(*cons.Cons, *environment.Environment) (types.Object, error)

// Function object
type Function struct {
	Imp      BuildinFunctionImp
	NumArgs  int
	EvalArgs bool
}

// Type of Function
func (fun *Function) Type() types.Type {
	return types.Function
}

// Eval evaluates a function
func (fun *Function) Eval(args *cons.Cons, env *environment.Environment) (types.Object, error) {
	return fun.Imp(args, env)
}
