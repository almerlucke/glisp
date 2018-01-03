package buildin

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Eval buildin function
func Eval(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return env.Eval(args.Car, context)
}

// CreateBuildinEval creates a buildin function object
func CreateBuildinEval() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Eval, 1, true)
}
