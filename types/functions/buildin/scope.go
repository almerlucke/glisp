package buildin

import (
	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/evaluator"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Scope buildin function
func Scope(args *cons.Cons, env *environment.Environment) (types.Object, error) {
	// Push a new scope
	env.PushScope(nil)

	// Make sure we pop the scope after completion
	defer env.PopScope()

	var val types.Object = types.NIL
	var err error

	err = args.Iter(func(obj types.Object, index uint64) error {
		val, err = evaluator.Eval(obj, env)
		return err
	})

	if err != nil {
		return nil, err
	}

	return val, nil
}

// CreateBuildinScope creates a buildin function object
func CreateBuildinScope() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Scope, 0, false)
}
