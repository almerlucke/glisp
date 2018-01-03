package array

import (
	"bytes"

	"github.com/almerlucke/glisp/interfaces/sequence"
	"github.com/almerlucke/glisp/types"
)

// Array type is just a slice typedef
type Array []types.Object

// Access array element
func (a Array) Access(nth uint64) types.Object {
	if nth > uint64(len(a)-1) {
		return nil
	}

	return a[nth]
}

// Assign array element
func (a Array) Assign(nth uint64, val types.Object) bool {
	if nth > uint64(len(a)-1) {
		return false
	}

	a[nth] = val

	return true
}

// Length of array
func (a Array) Length() uint64 {
	return uint64(len(a))
}

// Iter over array elements
func (a Array) Iter(f sequence.IterFun) error {
	for i, e := range a {
		err := f(e, uint64(i))
		if err != nil {
			return err
		}
	}

	return nil
}

// Map over array elements
func (a Array) Map(f sequence.MapFun) (sequence.Sequence, error) {
	ma := make(Array, len(a))

	for i, e := range a {
		me, err := f(e)
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
