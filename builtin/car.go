package builtin

import (
	"errors"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Car buildin function
func Car(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	obj := args.Car

	if obj == types.NIL {
		return types.NIL, nil
	}

	if obj.Type() == types.Cons {
		return obj.(*cons.Cons).Car, nil
	}

	return nil, errors.New("CAR expects a list as argument")
}

// CreateBuiltinCar creates a builtin function object
func CreateBuiltinCar() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Car, 1, true)
}
