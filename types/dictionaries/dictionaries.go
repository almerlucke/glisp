package dictionaries

import (
	"fmt"

	"github.com/almerlucke/glisp/interfaces/collection"
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

// Length of dictionary
func (d Dictionary) Length() uint64 {
	return uint64(len(d))
}

// Map maps over dictionary
func (d Dictionary) Map(fun collection.MapFun) (collection.Collection, error) {
	nd := make(Dictionary)

	for k, v := range d {
		nv, err := fun(v.value, v.originalKey)
		if err != nil {
			return nil, err
		}

		nd[k] = &dictionaryEntry{
			originalKey: v.originalKey,
			value:       nv,
		}
	}

	return nd, nil
}

// Iter iterates over dictionary
func (d Dictionary) Iter(fun collection.IterFun) error {
	for _, v := range d {
		stop, err := fun(v.value, v.originalKey)
		if err != nil {
			return err
		}

		if stop {
			break
		}
	}

	return nil
}

// Access a dictionary element
func (d Dictionary) Access(key interface{}) (types.Object, error) {
	hash, err := hashstructure.Hash(key, nil)
	if err != nil {
		return nil, err
	}

	v := d[hash]

	if v == nil {
		return types.NIL, nil
	}

	return v.value, nil
}

// Assign to a dictionary key
func (d Dictionary) Assign(key interface{}, val types.Object) error {
	hash, err := hashstructure.Hash(key, nil)
	if err != nil {
		return err
	}

	d[hash] = &dictionaryEntry{
		originalKey: key.(types.Object),
		value:       val,
	}

	return nil
}
