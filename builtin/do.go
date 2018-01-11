package builtin

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Do builtin function
func Do(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	var result types.Object = types.NIL
	var err error

	err = args.Iter(func(obj types.Object, index interface{}) (bool, error) {
		result, err = env.Eval(obj, context)
		return false, err
	})

	return result, err
}

// CreateBuiltinDo creates a builtin function object
func CreateBuiltinDo() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Do, 0, false)
}
