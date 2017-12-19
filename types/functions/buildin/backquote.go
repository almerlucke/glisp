package buildin

import (
	"errors"

	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/evaluator"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

func expansion(obj types.Object, env *environment.Environment) (types.Object, error) {
	if obj.Type() == types.Cons {
		builder := cons.ListBuilder{}

		for e := obj; e.Type() == types.Cons; e = e.(*cons.Cons).Cdr {
			car := e.(*cons.Cons).Car

			if car.Type() == types.Cons {
				l := car.(*cons.Cons)

				if l.Car == environment.UnquoteSymbol {
					// Unquote arg
					if l.Cdr.Type() != types.Cons {
						return nil, errors.New("Unquote needs one argument")
					}

					result, err := evaluator.Eval(l.Cdr.(*cons.Cons).Car, env)
					if err != nil {
						return nil, err
					}

					builder.PushBackObject(result)
				} else if l.Car == environment.SpliceSymbol {
					// Splice arg
					if l.Cdr.Type() != types.Cons {
						return nil, errors.New("Splice needs one argument")
					}

					result, err := evaluator.Eval(l.Cdr.(*cons.Cons).Car, env)
					if err != nil {
						return nil, err
					}

					if result.Type() != types.Cons {
						return nil, errors.New("Splice result must be a list")
					}

					builder.Append(result.(*cons.Cons))
				} else if l.Car == environment.BackquoteSymbol {
					// Recursively call backquote
					if l.Cdr.Type() != types.Cons {
						return nil, errors.New("Backquote needs one argument")
					}

					result, err := Backquote(l.Cdr.(*cons.Cons), env)
					if err != nil {
						return nil, err
					}

					builder.PushBackObject(result)
				} else {
					// Expand list further
					result, err := expansion(l, env)
					if err != nil {
						return nil, err
					}

					builder.PushBackObject(result)
				}
			} else {
				// No expansion needed
				builder.PushBackObject(car)
			}
		}

		return builder.Head, nil
	}

	return obj, nil
}

// Backquote buildin function
func Backquote(args *cons.Cons, env *environment.Environment) (types.Object, error) {
	// If not a list, return object unevaluated
	if args.Car.Type() != types.Cons {
		return args.Car, nil
	}

	l := args.Car.(*cons.Cons)

	if l.Car == environment.UnquoteSymbol {
		// Unquote arg
		if l.Cdr.Type() != types.Cons {
			return nil, errors.New("Unquote needs one argument")
		}

		return evaluator.Eval(l.Cdr.(*cons.Cons).Car, env)
	} else if l.Car == environment.SpliceSymbol {
		// Splice arg outside list context is an error
		return nil, errors.New("Splice can only be evaluated in a list context")
	} else if l.Car == environment.BackquoteSymbol {
		// Recursively call backquote
		if l.Cdr.Type() != types.Cons {
			return nil, errors.New("Backquote needs one argument")
		}

		return Backquote(l.Cdr.(*cons.Cons), env)
	}

	return expansion(l, env)
}

// CreateBuildinBackquote creates a buildin function object
func CreateBuildinBackquote() *functions.Function {
	return &functions.Function{
		NumArgs:  1,
		Imp:      Backquote,
		EvalArgs: false,
	}
}
