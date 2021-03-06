package builtin

import (
	"errors"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Cdr builtin function
func Cdr(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	obj := args.Car

	if obj == types.NIL {
		return types.NIL, nil
	}

	if obj.Type() == types.Cons {
		return obj.(*cons.Cons).Cdr, nil
	}

	return nil, errors.New("CDR expects a list as argument")
}

// CreateBuiltinCdr creates a builtin function object
func CreateBuiltinCdr() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Cdr, 1, true)
}
