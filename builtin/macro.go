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

// Macro buildin function
func Macro(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	argType := args.Car.Type()

	// Arg list must be cons or nil
	if argType != types.Cons && argType != types.Null {
		return nil, errors.New("macro expected an arg list as first argument")
	}

	symList := []*symbols.Symbol{}

	if argType == types.Cons {
		argList := args.Car.(*cons.Cons)
		pure, length := argList.Info()
		if !pure {
			return nil, errors.New("macro arg list must be a pure list")
		}

		symList = make([]*symbols.Symbol, length)

		// Build sym arg slice
		err := argList.Iter(func(obj types.Object, index interface{}) (bool, error) {
			if obj.Type() != types.Symbol {
				return false, errors.New("macro arg list must contain only symbols")
			}

			if obj.(*symbols.Symbol).Reserved {
				return false, fmt.Errorf("macro arg list contains reserved symbol %v", obj.(*symbols.Symbol))
			}

			symList[index.(uint64)] = obj.(*symbols.Symbol)

			return false, nil
		})

		if err != nil {
			return nil, err
		}
	}

	var body *cons.Cons
	if args.Cdr.Type() == types.Cons {
		body = args.Cdr.(*cons.Cons)
	}

	return functions.NewMacroFunction(symList, env.CaptureScope(), body), nil
}

// CreateBuildinMacro creates a buildin function object
func CreateBuildinMacro() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Macro, 1, false)
}
