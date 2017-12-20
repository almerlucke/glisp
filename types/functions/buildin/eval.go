package buildin

import (
	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/evaluator"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Eval buildin function
func Eval(args *cons.Cons, env *environment.Environment) (types.Object, error) {
	return evaluator.Eval(args.Car, env)
}

// CreateBuildinEval creates a buildin function object
func CreateBuildinEval() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Eval, 1, true)
}
