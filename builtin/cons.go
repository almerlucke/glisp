package builtin

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Cons buildin function
func Cons(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	obj1 := args.Car
	obj2 := args.Cdr.(*cons.Cons).Car

	return &cons.Cons{
		Car: obj1,
		Cdr: obj2,
	}, nil
}

// CreateBuildinCons creates a buildin function object
func CreateBuildinCons() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Cons, 2, true)
}
