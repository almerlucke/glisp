package array

import (
	"bytes"
	"errors"

	"github.com/almerlucke/glisp/interfaces/collection"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/numbers"
)

// Array type is just a slice typedef
type Array []types.Object

// Access array element
func (a Array) Access(index interface{}) (types.Object, error) {
	num, ok := index.(*numbers.Number)
	if !ok {
		return nil, errors.New("array expects a number index")
	}

	nth := num.Int64Value()
	if nth < 0 {
		return nil, errors.New("index out of bounds")
	}

	unth := uint64(nth)

	if unth > uint64(len(a)-1) {
		return nil, errors.New("index out of bounds")
	}

	return a[unth], nil
}

// Assign array element
func (a Array) Assign(index interface{}, val types.Object) error {
	num, ok := index.(*numbers.Number)
	if !ok {
		return errors.New("array expects a number index")
	}

	nth := num.Int64Value()
	if nth < 0 {
		return errors.New("index out of bounds")
	}

	unth := uint64(nth)

	if unth > uint64(len(a)-1) {
		return errors.New("index out of bounds")
	}

	a[nth] = val

	return nil
}

// Length of array
func (a Array) Length() uint64 {
	return uint64(len(a))
}

// Iter over array elements
func (a Array) Iter(f collection.IterFun) error {
	for i, e := range a {
		stop, err := f(e, uint64(i))
		if err != nil {
			return err
		}

		if stop {
			break
		}
	}

	return nil
}

// Map over array elements
func (a Array) Map(f collection.MapFun) (collection.Collection, error) {
	ma := make(Array, len(a))

	for i, e := range a {
		me, err := f(e, i)
		if err != nil {
			return nil, err
		}

		ma[i] = me
	}

	return ma, nil
}

// Type Array
func (a Array) Type() types.Type {
	return types.Array
}

// String implements the stringer interface
func (a Array) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("[")

	end := len(a) - 1

	for i, e := range a {
		buffer.WriteString(e.String())
		if i != end {
			buffer.WriteString(" ")
		}
	}

	buffer.WriteString("]")

	return buffer.String()
}
