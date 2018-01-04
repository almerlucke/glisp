package buildin

import (
	"errors"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
	"github.com/almerlucke/glisp/types/strings"
)

type tryContext struct {
	Err string
}

// Try buildin function
func Try(args *cons.Cons, env environment.Environment, context interface{}) (result types.Object, err error) {
	tryPart := args.Car
	catchPart := args.Cdr.(*cons.Cons).Car
	var alwaysPart types.Object = types.NIL

	if args.Cdr.(*cons.Cons).Cdr.Type() == types.Cons {
		alwaysPart = args.Cdr.(*cons.Cons).Cdr.(*cons.Cons).Car
	}

	env.PushDepthContext("TryDepth")

	defer func() {
		env.PopDepthContext("TryDepth")

		// Alway evaluate always part
		env.Eval(alwaysPart, context)

		if r := recover(); r != nil {
			tctx, ok := r.(*tryContext)
			if ok {
				result, err = env.Eval(catchPart, context)
				if err == nil {
					lb := cons.ListBuilder{}
					lb.PushBackObject(catchPart)
					lb.PushBackObject(strings.String(tctx.Err))
					result, err = env.Eval(lb.Head, context)
				}
			} else {
				// Continue to panic
				panic(r)
			}
		}
	}()

	return env.Eval(tryPart, context)
}

// Throw an exception
func Throw(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	if !env.HasDepthContext("TryDepth") {
		return nil, errors.New("throw can only be used inside a try block")
	}

	if args.Car.Type() != types.String {
		return nil, errors.New("throw expected a string as first argument")
	}

	tctx := &tryContext{
		Err: string(args.Car.(strings.String)),
	}

	panic(tctx)
}

// CreateBuildinTry creates a buildin function object
func CreateBuildinTry() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Try, 2, false)
}

// CreateBuildinThrow creates a buildin function object
func CreateBuildinThrow() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Throw, 1, true)
}
