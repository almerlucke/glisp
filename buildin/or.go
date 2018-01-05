package buildin

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Or buildin function
func Or(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	var result types.Object = types.NIL

	var err error

	err = args.Iter(func(obj types.Object, index interface{}) (bool, error) {
		result, err = env.Eval(obj, context)
		if err != nil {
			return false, err
		}

		if result != types.NIL {
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

// CreateBuildinOr creates a buildin function object
func CreateBuildinOr() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Or, 0, false)
}
