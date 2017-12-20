package buildin

import (
	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Quote buildin function
func Quote(args *cons.Cons, env *environment.Environment) (types.Object, error) {
	return args.Car, nil
}

// CreateBuildinQuote creates a buildin function object
func CreateBuildinQuote() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Quote, 1, false)
}
