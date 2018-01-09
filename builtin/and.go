package builtin

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// And buildin function
func And(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	var result types.Object = types.T

	var err error

	err = args.Iter(func(obj types.Object, index interface{}) (bool, error) {
		result, err = env.Eval(obj, context)
		if err != nil {
			return false, err
		}

		if result == types.NIL {
			// Signal to stop iteration
			return true, nil
		}

		return false, nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

// CreateBuiltinAnd creates a builtin function object
func CreateBuiltinAnd() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(And, 0, false)
}
