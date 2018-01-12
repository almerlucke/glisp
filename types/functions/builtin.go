package functions

import (
	"fmt"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
)

// BuiltinFunctionImp built in function implementation
type BuiltinFunctionImp func(*cons.Cons, environment.Environment, interface{}) (types.Object, error)

// BuiltinFunction object
type BuiltinFunction struct {
	imp      BuiltinFunctionImp
	numArgs  int
	evalArgs bool
}

// NewBuiltinFunction creates a new builtin function
func NewBuiltinFunction(imp BuiltinFunctionImp, numArgs int, evalArgs bool) *BuiltinFunction {
	return &BuiltinFunction{
		imp:      imp,
		numArgs:  numArgs,
		evalArgs: evalArgs,
	}
}

// NumArgs number of arguments
func (fun *BuiltinFunction) NumArgs() int {
	return fun.numArgs
}

// EvalArgs evaluate arguments before calling eval
func (fun *BuiltinFunction) EvalArgs() bool {
	return fun.evalArgs
}

// Type of Function
func (fun *BuiltinFunction) Type() types.Type {
	return types.Function
}

// Eval evaluates a function
func (fun *BuiltinFunction) Eval(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return fun.imp(args, env, context)
}

// String for Stringer
func (fun *BuiltinFunction) String() string {
	return fmt.Sprintf("function(%p)", fun)
}

// Eql obj
func (fun *BuiltinFunction) Eql(obj types.Object) bool {
	return fun == obj
}

// Equal obj
func (fun *BuiltinFunction) Equal(obj types.Object) bool {
	return fun == obj
}
