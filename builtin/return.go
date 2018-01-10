package builtin

import (
	"errors"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/interfaces/function"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Return builtin function
func Return(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	if !env.HasDepthContext("CallDepth") {
		return nil, errors.New("RETURN can only be used inside a macro or lambda body")
	}

	returnContext := function.ReturnContext{Object: args.Car}

	panic(&returnContext)
}

// CreateBuiltinReturn creates a builtin function object
func CreateBuiltinReturn() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Return, 1, true)
}
