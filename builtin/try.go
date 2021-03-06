package builtin

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

// Try builtin function
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

		if r := recover(); r != nil {
			tctx, ok := r.(*tryContext)
			if ok {
				result, err = env.Eval(catchPart, context)
				if err == nil {
					lb := cons.ListBuilder{}
					lb.PushBackObject(catchPart)
					lb.PushBackObject(strings.String(tctx.Err))
					result, err = env.Eval(lb.Head, context)

					// Evaluate always part
					env.Eval(alwaysPart, context)
				}
			} else {
				// Continue to panic
				panic(r)
			}
		} else {
			// Evaluate always part
			env.Eval(alwaysPart, context)
		}
	}()

	return env.Eval(tryPart, context)
}

// Throw an exception
func Throw(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	if !env.HasDepthContext("TryDepth") {
		return nil, errors.New("THROW can only be used inside a try block")
	}

	if args.Car.Type() != types.String {
		return nil, errors.New("THROW expected a string as first argument")
	}

	tctx := &tryContext{
		Err: string(args.Car.(strings.String)),
	}

	panic(tctx)
}

// CreateBuiltinTry creates a builtin function object
func CreateBuiltinTry() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Try, 2, false)
}

// CreateBuiltinThrow creates a builtin function object
func CreateBuiltinThrow() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Throw, 1, true)
}
