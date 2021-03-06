package builtin

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Scope buildin function
func Scope(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	// Push a new scope
	env.PushScope(nil)

	// Make sure we pop the scope after completion
	defer env.PopScope()

	var val types.Object = types.NIL
	var err error

	err = args.Iter(func(obj types.Object, index interface{}) (bool, error) {
		val, err = env.Eval(obj, context)
		return false, err
	})

	if err != nil {
		return nil, err
	}

	return val, nil
}

// CreateBuiltinScope creates a builtin function object
func CreateBuiltinScope() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Scope, 0, false)
}
