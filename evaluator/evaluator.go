package evaluator

import (
	"errors"
	"fmt"

	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions/function"
	"github.com/almerlucke/glisp/types/symbols"
)

// Eval a lisp object in the given environment
func Eval(obj types.Object, env *environment.Environment) (types.Object, error) {
	result := obj

	switch obj.Type() {

	case types.Symbol:
		result = env.GetBinding(obj.(*symbols.Symbol))
		if result == nil {
			return nil, fmt.Errorf("unbound symbol %v", obj)
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
			return nil, fmt.Errorf("eval %v is not a function", r)
		}

		fun := r.(function.Function)

		// Check for pure and get length
		pure, length := c.Info()
		if !pure {
			return nil, errors.New("eval can't evaluate a dotted list")
		}

		// Check if we have enough arguments
		if (length - 1) < int64(fun.NumArgs()) {
			return nil, fmt.Errorf("not enough arguments to function %v", c.Car)
		}

		// If we need to first evaluate all args
		var args *cons.Cons
		if c.Cdr != types.NIL {
			args = c.Cdr.(*cons.Cons)
			if fun.EvalArgs() {
				args, err = args.Map(func(obj types.Object) (types.Object, error) {
					return Eval(obj, env)
				})

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
