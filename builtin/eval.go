package builtin

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Eval builtin function
func Eval(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return env.Eval(args.Car, context)
}

// CreateBuiltinEval creates a builtin function object
func CreateBuiltinEval() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Eval, 1, true)
}
