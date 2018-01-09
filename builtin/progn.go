package builtin

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Progn builtin function
func Progn(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	var result types.Object = types.NIL
	var err error

	err = args.Iter(func(obj types.Object, index interface{}) (bool, error) {
		result, err = env.Eval(obj, context)
		return false, err
	})

	return result, err
}

// CreateBuiltinProgn creates a builtin function object
func CreateBuiltinProgn() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Progn, 0, false)
}
