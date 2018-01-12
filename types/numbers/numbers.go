package numbers

import (
	"errors"
	"fmt"
	"math"
	"reflect"

	"github.com/almerlucke/glisp/types"
)

// Numeric type
type Numeric interface {
	isNumeric()
}

// Int8 is int8
type Int8 int8

func (i Int8) isNumeric() {}

// Int16 is int16
type Int16 int16

func (i Int16) isNumeric() {}

// Int32 is int32
type Int32 int32

func (i Int32) isNumeric() {}

// Int64 is int64
type Int64 int64

func (i Int64) isNumeric() {}

// Uint8 is uint8
type Uint8 uint8

func (i Uint8) isNumeric() {}

// Uint16 is uint16
type Uint16 uint16

func (i Uint16) isNumeric() {}

// Uint32 is uint32
type Uint32 uint32

func (i Uint32) isNumeric() {}

// Uint64 is uint64
type Uint64 uint64

func (i Uint64) isNumeric() {}

// Float32 is float32
type Float32 float32

func (i Float32) isNumeric() {}

// Float64 is float64
type Float64 float64

func (i Float64) isNumeric() {}

// Number type
type Number struct {
	Kind  reflect.Kind
	Value Numeric
}

// Type Number for Object interface
func (num *Number) Type() types.Type {
	return types.Number
}

// String for stringer interface
func (num *Number) String() string {
	return fmt.Sprintf("%v", num.Value)
}

// Eql obj
func (num *Number) Eql(obj types.Object) bool {
	if obj.Type() == types.Number {
		return num.Value == obj.(*Number).Value
	}

	return false
}

// Equal obj
func (num *Number) Equal(obj types.Object) bool {
	return num.Eql(obj)
}

// Compare for comparable interface
func (num *Number) Compare(obj types.Comparable) (int, error) {
	otherNum, ok := obj.(*Number)

	if !ok {
		return 0, errors.New("unequal types for comparison")
	}

	g, err := num.GreaterThan(otherNum)
	if err != nil {
		return 0, err
	}

	if g {
		return 1, nil
	}

	g, _ = num.LesserThan(otherNum)

	if g {
		return -1, nil
	}

	return 0, nil
}

// New number
func New(kind reflect.Kind) *Number {
	num := &Number{
		Kind: kind,
	}

	switch kind {
	case reflect.Int8:
		num.Value = Int8(0)
	case reflect.Int16:
		num.Value = Int16(0)
	case reflect.Int32:
		num.Value = Int32(0)
	case reflect.Int64:
		num.Value = Int64(0)
	case reflect.Uint8:
		num.Value = Uint8(0)
	case reflect.Uint16:
		num.Value = Uint16(0)
	case reflect.Uint32:
		num.Value = Uint32(0)
	case reflect.Uint64:
		num.Value = Uint64(0)
	case reflect.Float32:
		num.Value = Float32(0)
	case reflect.Float64:
		num.Value = Float64(0)
	}

	return num
}

// NewInt8 new int8
func NewInt8(val int8) *Number {
	return &Number{
		Kind:  reflect.Int8,
		Value: Int8(val),
	}
}

// NewInt16 new int16
func NewInt16(val int16) *Number {
	return &Number{
		Kind:  reflect.Int16,
		Value: Int16(val),
	}
}

// NewInt32 new int32
func NewInt32(val int32) *Number {
	return &Number{
		Kind:  reflect.Int32,
		Value: Int32(val),
	}
}

// NewInt64 new int64
func NewInt64(val int64) *Number {
	return &Number{
		Kind:  reflect.Int64,
		Value: Int64(val),
	}
}

// NewUint8 new uint8
func NewUint8(val uint8) *Number {
	return &Number{
		Kind:  reflect.Uint8,
		Value: Uint8(val),
	}
}

// NewUint16 new uint16
func NewUint16(val uint16) *Number {
	return &Number{
		Kind:  reflect.Uint16,
		Value: Uint16(val),
	}
}

// NewUint32 new uint32
func NewUint32(val uint32) *Number {
	return &Number{
		Kind:  reflect.Uint32,
		Value: Uint32(val),
	}
}

// NewUint64 new uint64
func NewUint64(val uint64) *Number {
	return &Number{
		Kind:  reflect.Uint64,
		Value: Uint64(val),
	}
}

// NewFloat32 new float32
func NewFloat32(val float32) *Number {
	return &Number{
		Kind:  reflect.Float32,
		Value: Float32(val),
	}
}

// NewFloat64 new float64
func NewFloat64(val float64) *Number {
	return &Number{
		Kind:  reflect.Float64,
		Value: Float64(val),
	}
}

// SetInt64Value set value with int64
func (num *Number) SetInt64Value(val int64) {
	switch num.Kind {
	case reflect.Int8:
		num.Value = Int8(val)
	case reflect.Int16:
		num.Value = Int16(val)
	case reflect.Int32:
		num.Value = Int32(val)
	case reflect.Int64:
		num.Value = Int64(val)
	case reflect.Uint8:
		num.Value = Uint8(val)
	case reflect.Uint16:
		num.Value = Uint16(val)
	case reflect.Uint32:
		num.Value = Uint32(val)
	case reflect.Uint64:
		num.Value = Uint64(val)
	case reflect.Float32:
		num.Value = Float32(val)
	case reflect.Float64:
		num.Value = Float64(val)
	}
}

// SetUint64Value set value with uint64
func (num *Number) SetUint64Value(val uint64) {
	switch num.Kind {
	case reflect.Int8:
		num.Value = Int8(val)
	case reflect.Int16:
		num.Value = Int16(val)
	case reflect.Int32:
		num.Value = Int32(val)
	case reflect.Int64:
		num.Value = Int64(val)
	case reflect.Uint8:
		num.Value = Uint8(val)
	case reflect.Uint16:
		num.Value = Uint16(val)
	case reflect.Uint32:
		num.Value = Uint32(val)
	case reflect.Uint64:
		num.Value = Uint64(val)
	case reflect.Float32:
		num.Value = Float32(val)
	case reflect.Float64:
		num.Value = Float64(val)
	}
}

// SetFloat64Value set value with float64
func (num *Number) SetFloat64Value(val float64) {
	switch num.Kind {
	case reflect.Int8:
		num.Value = Int8(val)
	case reflect.Int16:
		num.Value = Int16(val)
	case reflect.Int32:
		num.Value = Int32(val)
	case reflect.Int64:
		num.Value = Int64(val)
	case reflect.Uint8:
		num.Value = Uint8(val)
	case reflect.Uint16:
		num.Value = Uint16(val)
	case reflect.Uint32:
		num.Value = Uint32(val)
	case reflect.Uint64:
		num.Value = Uint64(val)
	case reflect.Float32:
		num.Value = Float32(val)
	case reflect.Float64:
		num.Value = Float64(val)
	}
}

// IsZero check if num is zero
func (num *Number) IsZero() bool {
	switch num.Kind {
	case reflect.Int8:
		return num.Value.(Int8) == 0
	case reflect.Int16:
		return num.Value.(Int16) == 0
	case reflect.Int32:
		return num.Value.(Int32) == 0
	case reflect.Int64:
		return num.Value.(Int64) == 0
	case reflect.Uint8:
		return num.Value.(Uint8) == 0
	case reflect.Uint16:
		return num.Value.(Uint16) == 0
	case reflect.Uint32:
		return num.Value.(Uint32) == 0
	case reflect.Uint64:
		return num.Value.(Uint64) == 0
	case reflect.Float32:
		return num.Value.(Float32) == 0
	case reflect.Float64:
		return num.Value.(Float64) == 0
	}

	return false
}

// IsNaN is number a NaN
func (num *Number) IsNaN() bool {
	switch num.Kind {
	case reflect.Float32:
		return math.IsNaN(float64(num.Value.(Float32)))
	case reflect.Float64:
		return math.IsNaN(float64(num.Value.(Float64)))
	}

	return false
}

// IsInteger returns true if num is integer
func (num *Number) IsInteger() bool {
	isInteger := true

	switch num.Kind {
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		isInteger = false
	}

	return isInteger
}
