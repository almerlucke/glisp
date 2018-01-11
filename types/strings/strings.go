package strings

import (
	"errors"
	"fmt"
	"strings"

	"github.com/almerlucke/glisp/types"
)

// String type declaration
type String string

// Type String for Object interface
func (str String) Type() types.Type {
	return types.String
}

// String conform to Stringer
func (str String) String() string {
	return fmt.Sprintf("\"%v\"", string(str))
}

// Compare for comparable interface
func (str String) Compare(obj types.Comparable) (int, error) {
	otherStr, ok := obj.(String)

	if !ok {
		return 0, errors.New("unequal types for comparison")
	}

	return strings.Compare(string(str), string(otherStr)), nil
}
