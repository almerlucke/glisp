package builtin

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Not builtin function
func Not(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	if args.Car == types.NIL {
		return types.T, nil
	}

	return types.NIL, nil
}

// CreateBuiltinNot creates a builtin function object
func CreateBuiltinNot() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Not, 1, true)
}
