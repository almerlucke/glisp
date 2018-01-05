package buildin

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Progn buildin function
func Progn(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	var result types.Object = types.NIL
	var err error

	err = args.Iter(func(obj types.Object, index interface{}) error {
		result, err = env.Eval(obj, context)
		return err
	})

	return result, err
}

// CreateBuildinProgn creates a buildin function object
func CreateBuildinProgn() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Progn, 0, false)
}
