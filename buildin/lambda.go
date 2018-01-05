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

// Lambda buildin function
func Lambda(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	argType := args.Car.Type()

	// Arg list must be cons or nil
	if argType != types.Cons && argType != types.Null {
		return nil, errors.New("lambda expected an arg list as first argument")
	}

	symList := []*symbols.Symbol{}

	if argType == types.Cons {
		argList := args.Car.(*cons.Cons)
		pure, length := argList.Info()
		if !pure {
			return nil, errors.New("lambda arg list must be a pure list")
		}

		symList = make([]*symbols.Symbol, length)

		// Build sym arg slice
		err := argList.Iter(func(obj types.Object, index interface{}) error {
			if obj.Type() != types.Symbol {
				return errors.New("lambda arg list must contain only symbols")
			}

			if obj.(*symbols.Symbol).Reserved {
				return fmt.Errorf("lambda arg list contains reserved symbol %v", obj.(*symbols.Symbol))
			}

			symList[index.(uint64)] = obj.(*symbols.Symbol)

			return nil
		})

		if err != nil {
			return nil, err
		}
	}

	var body *cons.Cons
	if args.Cdr.Type() == types.Cons {
		body = args.Cdr.(*cons.Cons)
	}

	return functions.NewLambdaFunction(symList, env.CaptureScope(), body), nil
}

// CreateBuildinLambda creates a buildin function object
func CreateBuildinLambda() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Lambda, 1, false)
}
