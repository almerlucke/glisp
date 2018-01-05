package buildin

import (
	"errors"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/dictionaries"
	"github.com/almerlucke/glisp/types/functions"
)

// Dictionary buildin function
func Dictionary(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	dictionary := make(dictionaries.Dictionary)

	err := args.Iter(func(obj types.Object, index interface{}) error {
		if obj.Type() != types.Cons {
			return errors.New("illegal key value pair for dictionary")
		}

		pair := obj.(*cons.Cons)
		if pair.Length() != 2 {
			return errors.New("illegal key value pair for dictionary")
		}

		key := pair.Car
		value := pair.Cdr.(*cons.Cons).Car

		return dictionary.Assign(key, value)
	})

	if err != nil {
		return nil, err
	}

	return dictionary, nil
}

// CreateBuildinDictionary creates a buildin function object
func CreateBuildinDictionary() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Dictionary, 0, true)
}
