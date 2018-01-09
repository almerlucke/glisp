package numbers

import (
	"errors"
	"fmt"
	"math"
	"reflect"
)

// Add two numbers
func (num *Number) Add(otherNum *Number) (*Number, error) {
	if num.Kind != otherNum.Kind {
		return nil, fmt.Errorf("can't add %v to %v number", otherNum.Kind, num.Kind)
	}

	newNum := &Number{}

	switch num.Value.(type) {
	case Int8:
		newNum.Value = num.Value.(Int8) + otherNum.Value.(Int8)
		newNum.Kind = reflect.Int8
	case Int16:
		newNum.Value = num.Value.(Int16) + otherNum.Value.(Int16)
		newNum.Kind = reflect.Int16
	case Int32:
		newNum.Value = num.Value.(Int32) + otherNum.Value.(Int32)
		newNum.Kind = reflect.Int32
	case Int64:
		newNum.Value = num.Value.(Int64) + otherNum.Value.(Int64)
		newNum.Kind = reflect.Int64
	case Uint8:
		newNum.Value = num.Value.(Uint8) + otherNum.Value.(Uint8)
		newNum.Kind = reflect.Uint8
	case Uint16:
		newNum.Value = num.Value.(Uint16) + otherNum.Value.(Uint16)
		newNum.Kind = reflect.Uint16
	case Uint32:
		newNum.Value = num.Value.(Uint32) + otherNum.Value.(Uint32)
		newNum.Kind = reflect.Uint32
	case Uint64:
		newNum.Value = num.Value.(Uint64) + otherNum.Value.(Uint64)
		newNum.Kind = reflect.Uint64
	case Float32:
		newNum.Value = num.Value.(Float32) + otherNum.Value.(Float32)
		newNum.Kind = reflect.Float32
	case Float64:
		newNum.Value = num.Value.(Float64) + otherNum.Value.(Float64)
		newNum.Kind = reflect.Float64
	}

	return newNum, nil
}

// Subtract two numbers
func (num *Number) Subtract(otherNum *Number) (*Number, error) {
	if num.Kind != otherNum.Kind {
		return nil, fmt.Errorf("can't subtract %v from %v number", otherNum.Kind, num.Kind)
	}

	newNum := &Number{}

	switch num.Value.(type) {
	case Int8:
		newNum.Value = num.Value.(Int8) - otherNum.Value.(Int8)
		newNum.Kind = reflect.Int8
	case Int16:
		newNum.Value = num.Value.(Int16) - otherNum.Value.(Int16)
		newNum.Kind = reflect.Int16
	case Int32:
		newNum.Value = num.Value.(Int32) - otherNum.Value.(Int32)
		newNum.Kind = reflect.Int32
	case Int64:
		newNum.Value = num.Value.(Int64) - otherNum.Value.(Int64)
		newNum.Kind = reflect.Int64
	case Uint8:
		newNum.Value = num.Value.(Uint8) - otherNum.Value.(Uint8)
		newNum.Kind = reflect.Uint8
	case Uint16:
		newNum.Value = num.Value.(Uint16) - otherNum.Value.(Uint16)
		newNum.Kind = reflect.Uint16
	case Uint32:
		newNum.Value = num.Value.(Uint32) - otherNum.Value.(Uint32)
		newNum.Kind = reflect.Uint32
	case Uint64:
		newNum.Value = num.Value.(Uint64) - otherNum.Value.(Uint64)
		newNum.Kind = reflect.Uint64
	case Float32:
		newNum.Value = num.Value.(Float32) - otherNum.Value.(Float32)
		newNum.Kind = reflect.Float32
	case Float64:
		newNum.Value = num.Value.(Float64) - otherNum.Value.(Float64)
		newNum.Kind = reflect.Float64
	}

	return newNum, nil
}

// Multiply two numbers
func (num *Number) Multiply(otherNum *Number) (*Number, error) {
	if num.Kind != otherNum.Kind {
		return nil, fmt.Errorf("can't multiply %v with %v number", num.Kind, otherNum.Kind)
	}

	newNum := &Number{}

	switch num.Value.(type) {
	case Int8:
		newNum.Value = num.Value.(Int8) - otherNum.Value.(Int8)
		newNum.Kind = reflect.Int8
	case Int16:
		newNum.Value = num.Value.(Int16) - otherNum.Value.(Int16)
		newNum.Kind = reflect.Int16
	case Int32:
		newNum.Value = num.Value.(Int32) - otherNum.Value.(Int32)
		newNum.Kind = reflect.Int32
	case Int64:
		newNum.Value = num.Value.(Int64) - otherNum.Value.(Int64)
		newNum.Kind = reflect.Int64
	case Uint8:
		newNum.Value = num.Value.(Uint8) - otherNum.Value.(Uint8)
		newNum.Kind = reflect.Uint8
	case Uint16:
		newNum.Value = num.Value.(Uint16) - otherNum.Value.(Uint16)
		newNum.Kind = reflect.Uint16
	case Uint32:
		newNum.Value = num.Value.(Uint32) - otherNum.Value.(Uint32)
		newNum.Kind = reflect.Uint32
	case Uint64:
		newNum.Value = num.Value.(Uint64) - otherNum.Value.(Uint64)
		newNum.Kind = reflect.Uint64
	case Float32:
		newNum.Value = num.Value.(Float32) - otherNum.Value.(Float32)
		newNum.Kind = reflect.Float32
	case Float64:
		newNum.Value = num.Value.(Float64) - otherNum.Value.(Float64)
		newNum.Kind = reflect.Float64
	}

	return newNum, nil
}

// Divide two numbers
func (num *Number) Divide(otherNum *Number) (newNum *Number, err error) {
	if num.Kind != otherNum.Kind {
		return nil, fmt.Errorf("can't divide %v by %v number", num.Kind, otherNum.Kind)
	}

	defer func() {
		if r := recover(); r != nil {
			err = errors.New("division by zero detected")
		}
	}()

	newNum = &Number{}

	switch num.Value.(type) {
	case Int8:
		newNum.Value = num.Value.(Int8) / otherNum.Value.(Int8)
		newNum.Kind = reflect.Int8
	case Int16:
		newNum.Value = num.Value.(Int16) / otherNum.Value.(Int16)
		newNum.Kind = reflect.Int16
	case Int32:
		newNum.Value = num.Value.(Int32) / otherNum.Value.(Int32)
		newNum.Kind = reflect.Int32
	case Int64:
		newNum.Value = num.Value.(Int64) / otherNum.Value.(Int64)
		newNum.Kind = reflect.Int64
	case Uint8:
		newNum.Value = num.Value.(Uint8) / otherNum.Value.(Uint8)
		newNum.Kind = reflect.Uint8
	case Uint16:
		newNum.Value = num.Value.(Uint16) / otherNum.Value.(Uint16)
		newNum.Kind = reflect.Uint16
	case Uint32:
		newNum.Value = num.Value.(Uint32) / otherNum.Value.(Uint32)
		newNum.Kind = reflect.Uint32
	case Uint64:
		newNum.Value = num.Value.(Uint64) / otherNum.Value.(Uint64)
		newNum.Kind = reflect.Uint64
	case Float32:
		newNum.Value = num.Value.(Float32) / otherNum.Value.(Float32)
		newNum.Kind = reflect.Float32
	case Float64:
		newNum.Value = num.Value.(Float64) / otherNum.Value.(Float64)
		newNum.Kind = reflect.Float64
	}

	return newNum, nil
}

// Modulo two numbers
func (num *Number) Modulo(otherNum *Number) (newNum *Number, err error) {
	if num.Kind != otherNum.Kind {
		return nil, fmt.Errorf("can't modulo %v with %v number", num.Kind, otherNum.Kind)
	}

	defer func() {
		if r := recover(); r != nil {
			err = errors.New("division by zero detected")
		}
	}()

	newNum = &Number{}

	switch num.Value.(type) {
	case Int8:
		newNum.Value = num.Value.(Int8) % otherNum.Value.(Int8)
		newNum.Kind = reflect.Int8
	case Int16:
		newNum.Value = num.Value.(Int16) % otherNum.Value.(Int16)
		newNum.Kind = reflect.Int16
	case Int32:
		newNum.Value = num.Value.(Int32) % otherNum.Value.(Int32)
		newNum.Kind = reflect.Int32
	case Int64:
		newNum.Value = num.Value.(Int64) % otherNum.Value.(Int64)
		newNum.Kind = reflect.Int64
	case Uint8:
		newNum.Value = num.Value.(Uint8) % otherNum.Value.(Uint8)
		newNum.Kind = reflect.Uint8
	case Uint16:
		newNum.Value = num.Value.(Uint16) % otherNum.Value.(Uint16)
		newNum.Kind = reflect.Uint16
	case Uint32:
		newNum.Value = num.Value.(Uint32) % otherNum.Value.(Uint32)
		newNum.Kind = reflect.Uint32
	case Uint64:
		newNum.Value = num.Value.(Uint64) % otherNum.Value.(Uint64)
		newNum.Kind = reflect.Uint64
	case Float32:
		newNum.Value = Float32(math.Mod(float64(num.Value.(Float32)), float64(otherNum.Value.(Float32))))
		newNum.Kind = reflect.Float32
	case Float64:
		newNum.Value = Float32(math.Mod(float64(num.Value.(Float64)), float64(otherNum.Value.(Float64))))
		newNum.Kind = reflect.Float64
	}

	return newNum, nil
}
