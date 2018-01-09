package builtin

import (
	"errors"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/interfaces/function"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Return buildin function
func Return(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	if !env.HasDepthContext("CallDepth") {
		return nil, errors.New("return can only be used inside a macro or lambda body")
	}

	returnContext := function.ReturnContext{Object: args.Car}

	panic(&returnContext)
}

// CreateBuildinReturn creates a buildin function object
func CreateBuildinReturn() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Return, 1, true)
}
