package builtin

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Gensym builtin function
func Gensym(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return env.Gensym(), nil
}

// CreateBuiltinGensym creates a builtin function object
func CreateBuiltinGensym() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Gensym, 0, false)
}
