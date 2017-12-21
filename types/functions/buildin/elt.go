package buildin

import (
	"errors"

	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
	"github.com/almerlucke/glisp/types/numbers"
)

// Elt buildin function
func Elt(args *cons.Cons, env *environment.Environment) (types.Object, error) {
	if args.Car.Type() != types.Cons {
		return nil, errors.New("elt expected a list as first argument")
	}

	l := args.Car.(*cons.Cons)

	args = args.Cdr.(*cons.Cons)
	if args.Car.Type() != types.Number {
		return nil, errors.New("elt expected an integer as second argument")
	}

	num := args.Car.(*numbers.Number)
	index := num.Int64Value()
	if index < 0 {
		return nil, errors.New("index out of bounds")
	}

	val := l.Access(uint64(index))
	if val == nil {
		return nil, errors.New("index out of bounds")
	}

	return val, nil
}

// EltAssign assignable version of Elt
func EltAssign(args *cons.Cons, val types.Object, env *environment.Environment) (types.Object, error) {
	if args.Car.Type() != types.Cons {
		return nil, errors.New("elt expected a list as first argument")
	}

	l := args.Car.(*cons.Cons)

	args = args.Cdr.(*cons.Cons)
	if args.Car.Type() != types.Number {
		return nil, errors.New("elt expected an integer as second argument")
	}

	num := args.Car.(*numbers.Number)
	index := num.Int64Value()
	if index < 0 {
		return nil, errors.New("index out of bounds")
	}

	success := l.Assign(uint64(index), val)
	if !success {
		return nil, errors.New("index out of bounds")
	}

	return val, nil
}

// CreateBuildinElt creates a assignable function object
func CreateBuildinElt() *functions.AssignableFunction {
	return functions.NewAssignableFunction(Elt, EltAssign, 2, true)
}
