package functions

import (
	"fmt"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
)

// BuildinFunctionImp build in function implementation
type BuildinFunctionImp func(*cons.Cons, environment.Environment) (types.Object, error)

// BuildinFunction object
type BuildinFunction struct {
	imp      BuildinFunctionImp
	numArgs  int
	evalArgs bool
}

// NewBuildinFunction creates a new build in function
func NewBuildinFunction(imp BuildinFunctionImp, numArgs int, evalArgs bool) *BuildinFunction {
	return &BuildinFunction{
		imp:      imp,
		numArgs:  numArgs,
		evalArgs: evalArgs,
	}
}

// NumArgs number of arguments
func (fun *BuildinFunction) NumArgs() int {
	return fun.numArgs
}

// EvalArgs evaluate arguments before calling eval
func (fun *BuildinFunction) EvalArgs() bool {
	return fun.evalArgs
}

// Type of Function
func (fun *BuildinFunction) Type() types.Type {
	return types.Function
}

// Eval evaluates a function
func (fun *BuildinFunction) Eval(args *cons.Cons, env environment.Environment) (types.Object, error) {
	return fun.imp(args, env)
}

// String for Stringer
func (fun *BuildinFunction) String() string {
	return fmt.Sprintf("function(%p)", fun)
}
