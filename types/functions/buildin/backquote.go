package buildin

import (
	"errors"

	"github.com/almerlucke/glisp/globals/symbols"
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

func expansion(obj types.Object, env environment.Environment) (types.Object, error) {
	if obj.Type() != types.Cons {
		// If not a cons just return the object unevaluated
		return obj, nil
	}

	builder := cons.ListBuilder{}

	err := obj.(*cons.Cons).Iter(func(car types.Object, index uint64) error {
		if car.Type() == types.Cons {
			l := car.(*cons.Cons)

			if l.Car == symbols.UnquoteSymbol {
				// Unquote arg
				if l.Cdr.Type() != types.Cons {
					return errors.New("unquote needs one argument")
				}

				result, err := env.Eval(l.Cdr.(*cons.Cons).Car)
				if err != nil {
					return err
				}

				builder.PushBackObject(result)
			} else if l.Car == symbols.SpliceSymbol {
				// Splice arg
				if l.Cdr.Type() != types.Cons {
					return errors.New("splice needs one argument")
				}

				result, err := env.Eval(l.Cdr.(*cons.Cons).Car)
				if err != nil {
					return err
				}

				if result.Type() != types.Cons {
					return errors.New("splice result must be a list")
				}

				builder.Append(result.(*cons.Cons))
			} else if l.Car == symbols.BackquoteSymbol {
				// Recursively call backquote
				if l.Cdr.Type() != types.Cons {
					return errors.New("backquote needs one argument")
				}

				result, err := Backquote(l.Cdr.(*cons.Cons), env)
				if err != nil {
					return err
				}

				builder.PushBackObject(result)
			} else {
				// Expand list
				result, err := expansion(l, env)
				if err != nil {
					return err
				}

				builder.PushBackObject(result)
			}
		} else {
			// No expansion needed
			builder.PushBackObject(car)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return builder.Head, nil
}

// Backquote buildin function
func Backquote(args *cons.Cons, env environment.Environment) (types.Object, error) {
	obj := args.Car

	// If not a list, return object unevaluated
	if obj.Type() != types.Cons {
		return obj, nil
	}

	// Cast to list
	l := obj.(*cons.Cons)

	if l.Car == symbols.UnquoteSymbol {
		// Unquote arg
		if l.Cdr.Type() != types.Cons {
			return nil, errors.New("unquote needs one argument")
		}

		return env.Eval(l.Cdr.(*cons.Cons).Car)
	} else if l.Car == symbols.SpliceSymbol {
		// Splice arg outside list context is an error
		return nil, errors.New("splice can only be evaluated in a list context")
	} else if l.Car == symbols.BackquoteSymbol {
		// Recursively call backquote
		if l.Cdr.Type() != types.Cons {
			return nil, errors.New("backquote needs one argument")
		}

		return Backquote(l.Cdr.(*cons.Cons), env)
	}

	return expansion(l, env)
}

// CreateBuildinBackquote creates a buildin function object
func CreateBuildinBackquote() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Backquote, 1, false)
}
