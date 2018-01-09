package builtin

import (
	"errors"

	"github.com/almerlucke/glisp/interfaces/collection"
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Elt buildin function
func Elt(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	col, ok := args.Car.(collection.Collection)

	if !ok {
		return nil, errors.New("elt expected a collection as first argument")
	}

	index := args.Cdr.(*cons.Cons).Car

	val, err := col.Access(index)
	if err != nil {
		return nil, err
	}

	return val, nil
}

// EltAssign assignable version of Elt
func EltAssign(args *cons.Cons, val types.Object, env environment.Environment, context interface{}) (types.Object, error) {
	col, ok := args.Car.(collection.Collection)

	if !ok {
		return nil, errors.New("elt expected a collection as first argument")
	}

	index := args.Cdr.(*cons.Cons).Car

	err := col.Assign(index, val)
	if err != nil {
		return nil, err
	}

	return val, nil
}

// CreateBuildinElt creates a assignable function object
func CreateBuildinElt() *functions.AssignableFunction {
	return functions.NewAssignableFunction(Elt, EltAssign, 2, true)
}
