package builtin

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Not buildin function
func Not(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	if args.Car == types.NIL {
		return types.T, nil
	}

	return types.NIL, nil
}

// CreateBuildinNot creates a buildin function object
func CreateBuildinNot() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Not, 1, true)
}
