package number

import (
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
