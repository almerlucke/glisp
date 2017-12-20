package buildin

import (
	"errors"
	"fmt"

	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/evaluator"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
	"github.com/almerlucke/glisp/types/symbols"
)

// Assign buildin function
func Assign(args *cons.Cons, env *environment.Environment) (types.Object, error) {
	if args.Car.Type() != types.Symbol {
		return nil, errors.New("= expected a symbol as first argument")
	}

	sym := args.Car.(*symbols.Symbol)

	if sym.Reserved {
		return nil, fmt.Errorf("can't assign to a reserved symbol %v", sym)
	}

	args = args.Cdr.(*cons.Cons)
	val, err := evaluator.Eval(args.Car, env)
	if err != nil {
		return nil, err
	}

	err = env.SetBinding(sym, val)
	if err != nil {
		return nil, err
	}

	return val, nil
}

// CreateBuildinAssign creates a buildin function object
func CreateBuildinAssign() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Assign, 2, false)
}
