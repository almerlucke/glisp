package buildin

import (
	"errors"

	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Cdr buildin function
func Cdr(args *cons.Cons, env *environment.Environment) (types.Object, error) {
	obj := args.Car

	if obj == types.NIL {
		return types.NIL, nil
	}

	if obj.Type() == types.Cons {
		return obj.(*cons.Cons).Cdr, nil
	}

	return nil, errors.New("Cdr expects a list as argument")
}

// CreateBuildinCdr creates a buildin function object
func CreateBuildinCdr() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Cdr, 1, true)
}
