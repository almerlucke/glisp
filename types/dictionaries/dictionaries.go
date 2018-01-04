package dictionaries

import (
	"fmt"

	"github.com/almerlucke/glisp/types"
	"github.com/mitchellh/hashstructure"
)

type dictionaryEntry struct {
	originalKey types.Object
	value       types.Object
}

// Dictionary defines a dictionary with any object as key and any object as value
type Dictionary map[uint64]*dictionaryEntry

// Type Dictionary for Object interface
func (d Dictionary) Type() types.Type {
	return types.Dictionary
}

// String for stringer interface
func (d Dictionary) String() string {
	str := "(dictionary "
	first := true

	for _, v := range d {
		if !first {
			str += " "
		} else {
			first = false
		}

		str += fmt.Sprintf("(%v %v)", v.originalKey, v.value)
	}

	return str + ")"
}

// Get a value
func (d Dictionary) Get(key types.Object) (types.Object, error) {
	hash, err := hashstructure.Hash(key, nil)
	if err != nil {
		return nil, err
	}

	return d[hash].value, nil
}

// Set a value to a key
func (d Dictionary) Set(k types.Object, v types.Object) error {
	hash, err := hashstructure.Hash(k, nil)
	if err != nil {
		return err
	}

	d[hash] = &dictionaryEntry{
		originalKey: k,
		value:       v,
	}

	return nil
}
