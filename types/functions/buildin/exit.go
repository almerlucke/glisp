package buildin

import (
	"os"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Exit buildin function
func Exit(args *cons.Cons, env environment.Environment) (types.Object, error) {
	// Exit program
	os.Exit(0)

	return types.NIL, nil
}

// CreateBuildinExit creates a buildin function object
func CreateBuildinExit() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Exit, 0, false)
}
