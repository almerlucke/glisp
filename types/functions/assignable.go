package functions

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
)

// AssignableFunctionImp assign function implementation
type AssignableFunctionImp func(*cons.Cons, types.Object, environment.Environment, interface{}) (types.Object, error)

// AssignableFunction function that implements Assignable
type AssignableFunction struct {
	*BuiltinFunction
	assignImp AssignableFunctionImp
}

// NewAssignableFunction creates a new assignable function
func NewAssignableFunction(imp BuiltinFunctionImp, assignImp AssignableFunctionImp, numArgs int, evalArgs bool) *AssignableFunction {
	return &AssignableFunction{
		BuiltinFunction: NewBuiltinFunction(imp, numArgs, evalArgs),
		assignImp:       assignImp,
	}
}

// Assign call
func (fun *AssignableFunction) Assign(args *cons.Cons, val types.Object, env environment.Environment, context interface{}) (types.Object, error) {
	return fun.assignImp(args, val, env, context)
}
