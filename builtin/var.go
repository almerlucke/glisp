package builtin

import (
	"errors"
	"fmt"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
	"github.com/almerlucke/glisp/types/symbols"
)

// Var builtin function
func Var(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	if args.Car.Type() != types.Symbol {
		return nil, errors.New("VAR expected a symbol as first argument")
	}

	sym := args.Car.(*symbols.Symbol)

	if sym.Reserved {
		return nil, fmt.Errorf("can't assign to a reserved symbol %v", sym)
	}

	var val types.Object = types.NIL
	var err error

	if args.Cdr.Type() == types.Cons {
		args = args.Cdr.(*cons.Cons)
		val, err = env.Eval(args.Car, context)
		if err != nil {
			return nil, err
		}
	}

	env.AddBinding(sym, val)

	return val, nil
}

// CreateBuiltinVar creates a builtin function object
func CreateBuiltinVar() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Var, 1, false)
}
