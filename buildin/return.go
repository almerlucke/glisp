package buildin

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/interfaces/function"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Return buildin function
func Return(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	returnContext := function.ReturnContext{Object: args.Car}

	panic(&returnContext)
}

// CreateBuildinReturn creates a buildin function object
func CreateBuildinReturn() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Return, 1, true)
}
