package buildin

import (
	"errors"
	"fmt"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
	"github.com/almerlucke/glisp/types/symbols"
)

// Var buildin function
func Var(args *cons.Cons, env environment.Environment) (types.Object, error) {
	if args.Car.Type() != types.Symbol {
		return nil, errors.New("var expected a symbol as first argument")
	}

	sym := args.Car.(*symbols.Symbol)

	if sym.Reserved {
		return nil, fmt.Errorf("can't assign to a reserved symbol %v", sym)
	}

	var val types.Object = types.NIL
	var err error

	if args.Cdr.Type() == types.Cons {
		args = args.Cdr.(*cons.Cons)
		val, err = env.Eval(args.Car)
		if err != nil {
			return nil, err
		}
	}

	env.AddBinding(sym, val)

	return val, nil
}

// CreateBuildinVar creates a buildin function object
func CreateBuildinVar() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Var, 1, false)
}
