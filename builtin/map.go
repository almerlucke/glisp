package builtin

import (
	"errors"

	"github.com/almerlucke/glisp/interfaces/collection"
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/interfaces/function"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
	"github.com/almerlucke/glisp/types/numbers"
)

// Map builtin function
func Map(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	col, ok := args.Car.(collection.Collection)
	if !ok {
		return nil, errors.New("MAP expected a collection as first argument")
	}

	fun, ok := args.Cdr.(*cons.Cons).Car.(function.Function)
	if !ok {
		return nil, errors.New("MAP expected a function as second argument")
	}

	newCol, err := col.Map(func(obj types.Object, index interface{}) (types.Object, error) {
		objIndex, ok := index.(types.Object)
		if !ok {
			intIndex, ok := index.(uint64)
			if ok {
				objIndex = numbers.NewUint64(intIndex)
			} else {
				objIndex = types.NIL
			}
		}

		return fun.Eval(cons.ListFromSlice([]types.Object{obj, objIndex}), env, context)
	})

	return newCol, err
}

// CreateBuiltinMap creates a builtin function object
func CreateBuiltinMap() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Map, 2, true)
}
