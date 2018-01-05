package buildin

import (
	"errors"
	"fmt"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/interfaces/function"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
	"github.com/almerlucke/glisp/types/symbols"
)

func symbolAssign(sym *symbols.Symbol, val types.Object, env environment.Environment, context interface{}) (types.Object, error) {
	if sym.Reserved {
		return nil, fmt.Errorf("can't assign to a reserved symbol %v", sym)
	}

	val, err := env.Eval(val, context)
	if err != nil {
		return nil, err
	}

	err = env.SetBinding(sym, val)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func expressionAssign(c *cons.Cons, val types.Object, env environment.Environment, context interface{}) (types.Object, error) {
	r, err := env.Eval(c.Car, context)
	if err != nil {
		return nil, err
	}

	if r.Type() != types.Function {
		return nil, fmt.Errorf("can't assign to %v", r)
	}

	assignable, ok := r.(function.Assignable)
	if !ok {
		return nil, fmt.Errorf("can't assign to %v", r)
	}

	// Check for pure and get length
	pure, length := c.Info()
	if !pure {
		return nil, errors.New("assign can't evaluate a dotted list")
	}

	// Check if we have enough arguments
	if (length - 1) < int64(assignable.NumArgs()) {
		return nil, fmt.Errorf("not enough arguments to function %v", c.Car)
	}

	// If we need to first evaluate all args
	var args *cons.Cons
	if c.Cdr != types.NIL {
		args = c.Cdr.(*cons.Cons)
		if assignable.EvalArgs() {
			col, serr := args.Map(func(obj types.Object, index interface{}) (types.Object, error) {
				return env.Eval(obj, context)
			})

			if serr != nil {
				return nil, serr
			}

			args = col.(*cons.Cons)
		}
	}

	// Evaluate assign value if needed
	if assignable.EvalArgs() {
		val, err = env.Eval(val, context)
		if err != nil {
			return nil, err
		}
	}

	return assignable.Assign(args, val, env, context)
}

// Assign buildin function
func Assign(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	switch args.Car.Type() {
	case types.Symbol:
		return symbolAssign(args.Car.(*symbols.Symbol), args.Cdr.(*cons.Cons).Car, env, context)
	case types.Cons:
		return expressionAssign(args.Car.(*cons.Cons), args.Cdr.(*cons.Cons).Car, env, context)
	}

	return nil, fmt.Errorf("can't assign to %v", args.Car)
}

// CreateBuildinAssign creates a buildin function object
func CreateBuildinAssign() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Assign, 2, false)
}
