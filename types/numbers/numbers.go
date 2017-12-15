package numbers

import (
	"fmt"
	"reflect"

	"github.com/almerlucke/glisp/types"
)

// Number type
type Number struct {
	Kind  reflect.Kind
	Value interface{}
}

// Type Number for Object interface
func (num *Number) Type() types.Type {
	return types.Number
}

// String for stringer interface
func (num *Number) String() string {
	return fmt.Sprintf("%v", num.Value)
}
