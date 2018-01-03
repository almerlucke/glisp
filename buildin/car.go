package buildin

import (
	"errors"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Car buildin function
func Car(args *cons.Cons, env environment.Environment) (types.Object, error) {
	obj := args.Car

	if obj == types.NIL {
		return types.NIL, nil
	}

	if obj.Type() == types.Cons {
		return obj.(*cons.Cons).Car, nil
	}

	return nil, errors.New("car expects a list as argument")
}

// CreateBuildinCar creates a buildin function object
func CreateBuildinCar() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Car, 1, true)
}
