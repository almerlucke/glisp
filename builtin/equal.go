package builtin

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Eql builtin function
func Eql(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	eql := args.Car.Eql(args.Cdr.(*cons.Cons).Car)

	if eql {
		return types.T, nil
	}

	return types.NIL, nil
}

// Equal builtin function
func Equal(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	equal := args.Car.Equal(args.Cdr.(*cons.Cons).Car)

	if equal {
		return types.T, nil
	}

	return types.NIL, nil
}

// CreateBuiltinEql creates a builtin function object
func CreateBuiltinEql() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Eql, 2, true)
}

// CreateBuiltinEqual creates a builtin function object
func CreateBuiltinEqual() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Equal, 2, true)
}
