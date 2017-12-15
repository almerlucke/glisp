package strings

import "github.com/almerlucke/glisp/types"

// String type declaration
type String string

// Type String for Object interface
func (str String) Type() types.Type {
	return types.String
}
