package builtin

import (
	"fmt"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Print builtin function
func Print(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	if args != nil {
		args.Iter(func(obj types.Object, index interface{}) (bool, error) {
			fmt.Printf("%v\n", obj)
			return false, nil
		})
	}

	return types.NIL, nil
}

// CreateBuiltinPrint creates a builtin function object
func CreateBuiltinPrint() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Print, 0, true)
}
