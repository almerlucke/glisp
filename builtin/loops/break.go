package loops

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Break builtin function
func Break(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	panic(&loopContext{})
}

// CreateBuiltinBreak creates a builtin function object
func CreateBuiltinBreak() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Break, 0, false)
}
