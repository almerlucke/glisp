package builtin

import (
	"os"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Exit buildin function
func Exit(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	// Exit program
	os.Exit(0)

	return types.NIL, nil
}

// CreateBuiltinExit creates a builtin function object
func CreateBuiltinExit() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Exit, 0, false)
}
