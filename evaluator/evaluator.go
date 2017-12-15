package evaluator

import (
	"errors"
	"fmt"

	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
	"github.com/almerlucke/glisp/types/symbols"
)

// EvalArgs evaluate a list of arguments
func EvalArgs(c *cons.Cons, env *environment.Environment) (*cons.Cons, error) {
	// Evaluate arguments first by mapping eval
	return c.Map(func(obj types.Object) (types.Object, error) {
		r, err := Eval(obj, env)
		if err != nil {
			return nil, err
		}

		return r, nil
	})
}

// Eval a lisp object in the given environment
func Eval(obj types.Object, env *environment.Environment) (types.Object, error) {
	result := obj

	switch obj.Type() {
	case types.Symbol:
		result = env.GetBinding(obj.(*symbols.Symbol))
		if result == nil {
			return nil, fmt.Errorf("Unbound symbol %v", obj)
		}
	case types.Cons:
		// List to evaluate
		c := obj.(*cons.Cons)

		// Evaluate first elem
		r, err := Eval(c.Car, env)
		if err != nil {
			return nil, err
		}

		// Must be a function
		if r.Type() != types.Function {
			return nil, fmt.Errorf("Eval %v is not a function", r)
		}

		fun := r.(*functions.Function)

		// Check for pure and get length
		pure, length := c.Info()
		if !pure {
			return nil, errors.New("Eval can't evaluate a dotted list")
		}

		// Check if we have enough arguments
		if (length - 1) < int64(fun.NumArgs) {
			return nil, fmt.Errorf("Not enough arguments to function %v", c.Car)
		}

		// If we need to first evaluate all args
		var args *cons.Cons
		if c.Cdr != types.NIL {
			args = c.Cdr.(*cons.Cons)
			if fun.EvalArgs {
				args, err = EvalArgs(args, env)
				if err != nil {
					return nil, err
				}
			}
		}

		// Evaluate function call
		result, err = fun.Eval(args, env)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
