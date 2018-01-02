package hashtables

import (
	"errors"
	"fmt"

	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/numbers"
	"github.com/almerlucke/glisp/types/symbols"
)

// HashTable defines a hash table with a comparable type as key and any object as value
type HashTable map[interface{}]types.Object

// Type HashTable for Object interface
func (m HashTable) Type() types.Type {
	return types.HashTable
}

// String for stringer interface
func (m HashTable) String() string {
	str := "(hashtable "
	first := true

	for k, v := range m {
		if !first {
			str += " "
		} else {
			first = false
		}

		kstr := fmt.Sprintf("%v", k)

		switch k.(type) {
		case numbers.Number:
			num := k.(numbers.Number)
			kstr = fmt.Sprintf("%v", &num)
		case symbols.Symbol:
			sym := k.(symbols.Symbol)
			kstr = fmt.Sprintf("%v", &sym)
		}

		str += fmt.Sprintf("(%v %v)", kstr, v)
	}

	return str + ")"
}

// Get a value
func (m HashTable) Get(k types.Object) (types.Object, error) {
	_, ok := k.(types.Comparable)
	if !ok {
		return nil, errors.New("hashtable key must be comparable")
	}

	var v types.Object

	switch k.(type) {
	case *numbers.Number:
		v = m[*(k.(*numbers.Number))]
	case *symbols.Symbol:
		v = m[*(k.(*symbols.Symbol))]
	default:
		v = m[k]
	}

	return v, nil
}

// Set a value to a key
func (m HashTable) Set(k types.Object, v types.Object) error {
	_, ok := k.(types.Comparable)
	if !ok {
		return errors.New("hashtable key must be comparable")
	}

	switch k.(type) {
	case *numbers.Number:
		m[*(k.(*numbers.Number))] = v
	case *symbols.Symbol:
		m[*(k.(*symbols.Symbol))] = v
	default:
		m[k] = v
	}

	return nil
}
