package buildin

import (
	"errors"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
	"github.com/almerlucke/glisp/types/hashtables"
)

// HashTable buildin function
func HashTable(args *cons.Cons, env environment.Environment) (types.Object, error) {
	hash := make(hashtables.HashTable)

	err := args.Iter(func(obj types.Object, index uint64) error {
		if obj.Type() != types.Cons {
			return errors.New("illegal key value pair for hash table")
		}

		pair := obj.(*cons.Cons)
		if pair.Length() != 2 {
			return errors.New("illegal key value pair for hash table")
		}

		key := pair.Car
		value := pair.Cdr.(*cons.Cons).Car

		return hash.Set(key, value)
	})

	if err != nil {
		return nil, err
	}

	return hash, nil
}

// CreateBuildinHashTable creates a buildin function object
func CreateBuildinHashTable() *functions.BuildinFunction {
	return functions.NewBuildinFunction(HashTable, 0, true)
}
