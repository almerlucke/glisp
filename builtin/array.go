package builtin

import (
	"errors"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/arrays"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
	"github.com/almerlucke/glisp/types/numbers"
)

// Array builtin function creates an array with initial elements
// or a specific size
func Array(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	a := make(arrays.Array, args.Length())

	_ = args.Iter(func(obj types.Object, index interface{}) (bool, error) {
		a[index.(uint64)] = obj
		return false, nil
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

	a := make(arrays.Array, ln)
	for i := range a {
		a[i] = initialObject
	}

	return a, nil
}

// CreateBuiltinMakeArray creates an array function object
func CreateBuiltinMakeArray() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(MakeArray, 1, true)
}

// CreateBuiltinArray creates an array function object
func CreateBuiltinArray() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Array, 1, true)
}
