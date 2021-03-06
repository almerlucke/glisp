package builtin

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// If else
func If(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	result, err := env.Eval(args.Car, context)
	if err != nil {
		return nil, err
	}

	ifPart := args.Cdr.(*cons.Cons)

	if result != types.NIL {
		result, err = env.Eval(ifPart.Car, context)
	} else {
		elsePart := ifPart.Cdr
		if elsePart.Type() == types.Cons {
			result, err = env.Eval(elsePart.(*cons.Cons).Car, context)
		}
	}

	return result, err
}

// CreateBuiltinIf creates a builtin function object
func CreateBuiltinIf() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(If, 2, false)
}
