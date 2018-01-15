package strings

import (
	"errors"
	"fmt"
	"strings"

	"github.com/almerlucke/glisp/interfaces/collection"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/characters"
	"github.com/almerlucke/glisp/types/numbers"
)

// String type declaration
type String []rune

// Access string element
func (str String) Access(index interface{}) (types.Object, error) {
	num, ok := index.(*numbers.Number)
	if !ok {
		return nil, errors.New("string expects a number index")
	}

	nth := num.Int64Value()
	if nth < 0 {
		return nil, errors.New("index out of bounds")
	}

	unth := uint64(nth)

	if unth > uint64(len(str)-1) {
		return nil, errors.New("index out of bounds")
	}

	chr := characters.Character(str[unth])

	return chr, nil
}

// Assign string element
func (str String) Assign(index interface{}, val types.Object) error {
	num, ok := index.(*numbers.Number)
	if !ok {
		return errors.New("string expects a number index")
	}

	nth := num.Int64Value()
	if nth < 0 {
		return errors.New("index out of bounds")
	}

	unth := uint64(nth)

	if unth > uint64(len(str)-1) {
		return errors.New("index out of bounds")
	}

	v, ok := val.(characters.Character)
	if !ok {
		return errors.New("string expects a character value")
	}

	str[nth] = rune(v)

	return nil
}

// Length of array
func (str String) Length() uint64 {
	return uint64(len(str))
}

// Iter over array elements
func (str String) Iter(f collection.IterFun) error {
	for i, e := range str {
		stop, err := f(characters.Character(e), uint64(i))
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
func (str String) Map(f collection.MapFun) (collection.Collection, error) {
	mstr := make(String, len(str))

	for i, e := range str {
		me, err := f(characters.Character(e), i)
		if err != nil {
			return nil, err
		}

		chr, ok := me.(characters.Character)
		if !ok {
			return nil, errors.New("string map expects a character value")
		}

		mstr[i] = rune(chr)
	}

	return mstr, nil
}

// Type String for Object interface
func (str String) Type() types.Type {
	return types.String
}

// String conform to Stringer
func (str String) String() string {
	return fmt.Sprintf("\"%v\"", string(str))
}

// Eql obj
func (str String) Eql(obj types.Object) bool {
	if obj.Type() == types.String {
		return string(str) == string(obj.(String))
	}

	return false
}

// Equal obj
func (str String) Equal(obj types.Object) bool {
	return str.Eql(obj)
}

// Compare for comparable interface
func (str String) Compare(obj types.Comparable) (int, error) {
	otherStr, ok := obj.(String)

	if !ok {
		return 0, errors.New("unequal types for comparison")
	}

	return strings.Compare(string(str), string(otherStr)), nil
}
