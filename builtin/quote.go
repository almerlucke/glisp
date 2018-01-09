package builtin

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Quote builtin function
func Quote(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return args.Car, nil
}

// CreateBuiltinQuote creates a builtin function object
func CreateBuiltinQuote() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Quote, 1, false)
}
