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

// Uint64Value returns the uint64 representation
func (num *Number) Uint64Value() uint64 {
	var val = uint64(0)

	switch num.Kind {
	case reflect.Float32:
		val = uint64(num.Value.(float32))
	case reflect.Float64:
		val = uint64(num.Value.(float64))
	case reflect.Int:
		val = uint64(num.Value.(int))
	case reflect.Int8:
		val = uint64(num.Value.(int8))
	case reflect.Int16:
		val = uint64(num.Value.(int16))
	case reflect.Int32:
		val = uint64(num.Value.(int32))
	case reflect.Int64:
		val = uint64(num.Value.(int64))
	case reflect.Uint:
		val = uint64(num.Value.(uint))
	case reflect.Uint8:
		val = uint64(num.Value.(uint8))
	case reflect.Uint16:
		val = uint64(num.Value.(uint16))
	case reflect.Uint32:
		val = uint64(num.Value.(uint32))
	case reflect.Uint64:
		val = num.Value.(uint64)
	}

	return val
}

// Int64Value returns the int64 representation
func (num *Number) Int64Value() int64 {
	var val = int64(0)

	switch num.Kind {
	case reflect.Float32:
		val = int64(num.Value.(float32))
	case reflect.Float64:
		val = int64(num.Value.(float64))
	case reflect.Int:
		val = int64(num.Value.(int))
	case reflect.Int8:
		val = int64(num.Value.(int8))
	case reflect.Int16:
		val = int64(num.Value.(int16))
	case reflect.Int32:
		val = int64(num.Value.(int32))
	case reflect.Int64:
		val = num.Value.(int64)
	case reflect.Uint:
		val = int64(num.Value.(uint))
	case reflect.Uint8:
		val = int64(num.Value.(uint8))
	case reflect.Uint16:
		val = int64(num.Value.(uint16))
	case reflect.Uint32:
		val = int64(num.Value.(uint32))
	case reflect.Uint64:
		val = int64(num.Value.(uint64))
	}

	return val
}
