package numbers

import "reflect"

// Int8Value returns the int8 representation
func (num *Number) Int8Value() int8 {
	var val = int8(0)

	switch num.Value.(type) {
	case Int8:
		val = int8(num.Value.(Int8))
	case Int16:
		val = int8(num.Value.(Int16))
	case Int32:
		val = int8(num.Value.(Int32))
	case Int64:
		val = int8(num.Value.(Int64))
	case Uint8:
		val = int8(num.Value.(Uint8))
	case Uint16:
		val = int8(num.Value.(Uint16))
	case Uint32:
		val = int8(num.Value.(Uint32))
	case Uint64:
		val = int8(num.Value.(Uint64))
	case Float32:
		val = int8(num.Value.(Float32))
	case Float64:
		val = int8(num.Value.(Float64))
	}

	return val
}

// Int8 converts num to Int8
func (num *Number) Int8() *Number {
	return &Number{
		Kind:  reflect.Int8,
		Value: Int8(num.Int8Value()),
	}
}

// Int16Value returns the int16 representation
func (num *Number) Int16Value() int16 {
	var val = int16(0)

	switch num.Value.(type) {
	case Int8:
		val = int16(num.Value.(Int8))
	case Int16:
		val = int16(num.Value.(Int16))
	case Int32:
		val = int16(num.Value.(Int32))
	case Int64:
		val = int16(num.Value.(Int64))
	case Uint8:
		val = int16(num.Value.(Uint8))
	case Uint16:
		val = int16(num.Value.(Uint16))
	case Uint32:
		val = int16(num.Value.(Uint32))
	case Uint64:
		val = int16(num.Value.(Uint64))
	case Float32:
		val = int16(num.Value.(Float32))
	case Float64:
		val = int16(num.Value.(Float64))
	}

	return val
}

// Int16 converts num to Int16
func (num *Number) Int16() *Number {
	return &Number{
		Kind:  reflect.Int16,
		Value: Int16(num.Int16Value()),
	}
}

// Int32Value returns the int32 representation
func (num *Number) Int32Value() int32 {
	var val = int32(0)

	switch num.Value.(type) {
	case Int8:
		val = int32(num.Value.(Int8))
	case Int16:
		val = int32(num.Value.(Int16))
	case Int32:
		val = int32(num.Value.(Int32))
	case Int64:
		val = int32(num.Value.(Int64))
	case Uint8:
		val = int32(num.Value.(Uint8))
	case Uint16:
		val = int32(num.Value.(Uint16))
	case Uint32:
		val = int32(num.Value.(Uint32))
	case Uint64:
		val = int32(num.Value.(Uint64))
	case Float32:
		val = int32(num.Value.(Float32))
	case Float64:
		val = int32(num.Value.(Float64))
	}

	return val
}

// Int32 converts num to Int32
func (num *Number) Int32() *Number {
	return &Number{
		Kind:  reflect.Int32,
		Value: Int32(num.Int32Value()),
	}
}

// Int64Value returns the int64 representation
func (num *Number) Int64Value() int64 {
	var val = int64(0)

	switch num.Value.(type) {
	case Int8:
		val = int64(num.Value.(Int8))
	case Int16:
		val = int64(num.Value.(Int16))
	case Int32:
		val = int64(num.Value.(Int32))
	case Int64:
		val = int64(num.Value.(Int64))
	case Uint8:
		val = int64(num.Value.(Uint8))
	case Uint16:
		val = int64(num.Value.(Uint16))
	case Uint32:
		val = int64(num.Value.(Uint32))
	case Uint64:
		val = int64(num.Value.(Uint64))
	case Float32:
		val = int64(num.Value.(Float32))
	case Float64:
		val = int64(num.Value.(Float64))
	}

	return val
}

// Int64 converts num to Int64
func (num *Number) Int64() *Number {
	return &Number{
		Kind:  reflect.Int64,
		Value: Int64(num.Int64Value()),
	}
}

// Uint8Value returns the uint8 representation
func (num *Number) Uint8Value() uint8 {
	var val = uint8(0)

	switch num.Value.(type) {
	case Int8:
		val = uint8(num.Value.(Int8))
	case Int16:
		val = uint8(num.Value.(Int16))
	case Int32:
		val = uint8(num.Value.(Int32))
	case Int64:
		val = uint8(num.Value.(Int64))
	case Uint8:
		val = uint8(num.Value.(Uint8))
	case Uint16:
		val = uint8(num.Value.(Uint16))
	case Uint32:
		val = uint8(num.Value.(Uint32))
	case Uint64:
		val = uint8(num.Value.(Uint64))
	case Float32:
		val = uint8(num.Value.(Float32))
	case Float64:
		val = uint8(num.Value.(Float64))
	}

	return val
}

// Uint8 converts num to Uint8
func (num *Number) Uint8() *Number {
	return &Number{
		Kind:  reflect.Uint8,
		Value: Uint8(num.Uint8Value()),
	}
}

// Uint16Value returns the uint16 representation
func (num *Number) Uint16Value() uint16 {
	var val = uint16(0)

	switch num.Value.(type) {
	case Int8:
		val = uint16(num.Value.(Int8))
	case Int16:
		val = uint16(num.Value.(Int16))
	case Int32:
		val = uint16(num.Value.(Int32))
	case Int64:
		val = uint16(num.Value.(Int64))
	case Uint8:
		val = uint16(num.Value.(Uint8))
	case Uint16:
		val = uint16(num.Value.(Uint16))
	case Uint32:
		val = uint16(num.Value.(Uint32))
	case Uint64:
		val = uint16(num.Value.(Uint64))
	case Float32:
		val = uint16(num.Value.(Float32))
	case Float64:
		val = uint16(num.Value.(Float64))
	}

	return val
}

// Uint16 converts num to Uint16
func (num *Number) Uint16() *Number {
	return &Number{
		Kind:  reflect.Uint16,
		Value: Uint16(num.Uint16Value()),
	}
}

// Uint32Value returns the uint32 representation
func (num *Number) Uint32Value() uint32 {
	var val = uint32(0)

	switch num.Value.(type) {
	case Int8:
		val = uint32(num.Value.(Int8))
	case Int16:
		val = uint32(num.Value.(Int16))
	case Int32:
		val = uint32(num.Value.(Int32))
	case Int64:
		val = uint32(num.Value.(Int64))
	case Uint8:
		val = uint32(num.Value.(Uint8))
	case Uint16:
		val = uint32(num.Value.(Uint16))
	case Uint32:
		val = uint32(num.Value.(Uint32))
	case Uint64:
		val = uint32(num.Value.(Uint64))
	case Float32:
		val = uint32(num.Value.(Float32))
	case Float64:
		val = uint32(num.Value.(Float64))
	}

	return val
}

// Uint32 converts num to Uint32
func (num *Number) Uint32() *Number {
	return &Number{
		Kind:  reflect.Uint32,
		Value: Uint32(num.Uint32Value()),
	}
}

// Uint64Value returns the uint64 representation
func (num *Number) Uint64Value() uint64 {
	var val = uint64(0)

	switch num.Value.(type) {
	case Int8:
		val = uint64(num.Value.(Int8))
	case Int16:
		val = uint64(num.Value.(Int16))
	case Int32:
		val = uint64(num.Value.(Int32))
	case Int64:
		val = uint64(num.Value.(Int64))
	case Uint8:
		val = uint64(num.Value.(Uint8))
	case Uint16:
		val = uint64(num.Value.(Uint16))
	case Uint32:
		val = uint64(num.Value.(Uint32))
	case Uint64:
		val = uint64(num.Value.(Uint64))
	case Float32:
		val = uint64(num.Value.(Float32))
	case Float64:
		val = uint64(num.Value.(Float64))
	}

	return val
}

// Uint64 converts num to Uint64
func (num *Number) Uint64() *Number {
	return &Number{
		Kind:  reflect.Uint64,
		Value: Uint64(num.Uint64Value()),
	}
}

// Float32Value returns the float32 representation
func (num *Number) Float32Value() float32 {
	var val = float32(0)

	switch num.Value.(type) {
	case Int8:
		val = float32(num.Value.(Int8))
	case Int16:
		val = float32(num.Value.(Int16))
	case Int32:
		val = float32(num.Value.(Int32))
	case Int64:
		val = float32(num.Value.(Int64))
	case Uint8:
		val = float32(num.Value.(Uint8))
	case Uint16:
		val = float32(num.Value.(Uint16))
	case Uint32:
		val = float32(num.Value.(Uint32))
	case Uint64:
		val = float32(num.Value.(Uint64))
	case Float32:
		val = float32(num.Value.(Float32))
	case Float64:
		val = float32(num.Value.(Float64))
	}

	return val
}

// Float32 converts num to Float32
func (num *Number) Float32() *Number {
	return &Number{
		Kind:  reflect.Float32,
		Value: Float32(num.Float32Value()),
	}
}

// Float64Value returns the float64 representation
func (num *Number) Float64Value() float64 {
	var val = float64(0)

	switch num.Value.(type) {
	case Int8:
		val = float64(num.Value.(Int8))
	case Int16:
		val = float64(num.Value.(Int16))
	case Int32:
		val = float64(num.Value.(Int32))
	case Int64:
		val = float64(num.Value.(Int64))
	case Uint8:
		val = float64(num.Value.(Uint8))
	case Uint16:
		val = float64(num.Value.(Uint16))
	case Uint32:
		val = float64(num.Value.(Uint32))
	case Uint64:
		val = float64(num.Value.(Uint64))
	case Float32:
		val = float64(num.Value.(Float32))
	case Float64:
		val = float64(num.Value.(Float64))
	}

	return val
}

// Float64 converts num to Float64
func (num *Number) Float64() *Number {
	return &Number{
		Kind:  reflect.Float64,
		Value: Float64(num.Float64Value()),
	}
}
