package builtin

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Unquote builtin function
func Unquote(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return args.Car, nil
}

// CreateBuiltinUnquote creates a builtin function object
func CreateBuiltinUnquote() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Unquote, 1, true)
}
