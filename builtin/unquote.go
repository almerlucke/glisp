package builtin

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Unquote buildin function
func Unquote(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return args.Car, nil
}

// CreateBuildinUnquote creates a buildin function object
func CreateBuildinUnquote() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Unquote, 1, true)
}
