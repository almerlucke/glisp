package strings

import (
	"fmt"

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