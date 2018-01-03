package buildin

import (
	"fmt"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Print buildin function
func Print(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	if args != nil {
		args.Iter(func(obj types.Object, index uint64) error {
			fmt.Printf("%v\n", obj)
			return nil
		})
	}

	return types.NIL, nil
}

// CreateBuildinPrint creates a buildin function object
func CreateBuildinPrint() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Print, 0, true)
}
