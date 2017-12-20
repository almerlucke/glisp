package buildin

import (
	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// List buildin function
func List(args *cons.Cons, env *environment.Environment) (types.Object, error) {
	if args == nil {
		return types.NIL, nil
	}

	return args, nil
}

// CreateBuildinList creates a buildin function object
func CreateBuildinList() *functions.BuildinFunction {
	return functions.NewBuildinFunction(List, 0, true)
}
