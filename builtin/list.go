package builtin

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// List buildin function
func List(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	if args == nil {
		return types.NIL, nil
	}

	return args, nil
}

// CreateBuiltinList creates a builtin function object
func CreateBuiltinList() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(List, 0, true)
}
