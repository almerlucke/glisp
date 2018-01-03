package buildin

import (
	"errors"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/array"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
	"github.com/almerlucke/glisp/types/numbers"
)

// Array buildin function creates an array with initial elements
// or a specific size
func Array(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	a := make(array.Array, args.Length())

	_ = args.Iter(func(obj types.Object, index uint64) error {
		a[index] = obj
		return nil
	})

	return a, nil
}

// MakeArray creates an array of a specific length with an optional initial element
func MakeArray(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	n := args.Car
	var initialObject types.Object = types.NIL

	if args.Cdr.Type() == types.Cons {
		initialObject = args.Cdr.(*cons.Cons).Car
	}

	if n.Type() != types.Number {
		return nil, errors.New("make-array expected a number as first argument")
	}

	ln := n.(*numbers.Number).Int64Value()
	if ln < 0 {
		return nil, errors.New("make-array expected a positive number as first argument")
	}

	a := make(array.Array, ln)
	for i := range a {
		a[i] = initialObject
	}

	return a, nil
}

// CreateBuildinMakeArray creates an array function object
func CreateBuildinMakeArray() *functions.BuildinFunction {
	return functions.NewBuildinFunction(MakeArray, 1, true)
}

// CreateBuildinArray creates an array function object
func CreateBuildinArray() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Array, 1, true)
}
