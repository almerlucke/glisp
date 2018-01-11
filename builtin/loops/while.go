package loops

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// While builtin function
func While(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	condition := args.Car
	var result types.Object = types.NIL
	var body types.Object = types.NIL

	if args.Cdr.Type() == types.Cons {
		body = args.Cdr.(*cons.Cons).Car
	}

	env.PushDepthContext(loopDepth)

	defer func() {
		env.PopDepthContext(loopDepth)

		if r := recover(); r != nil {
			_, ok := r.(*loopContext)
			if ok {
				// do nothing, just catch a break
			} else {
				// Continue to panic
				panic(r)
			}
		}
	}()

	for {
		obj, err := env.Eval(condition, context)
		if err != nil {
			return nil, err
		}

		if obj == types.NIL {
			break
		}

		result, err = env.Eval(body, context)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// CreateBuiltinWhile creates a builtin function object
func CreateBuiltinWhile() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(While, 1, false)
}
