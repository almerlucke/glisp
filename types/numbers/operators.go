package numbers

import (
	"errors"
	"fmt"
	"math"
	"reflect"
)

// LesserThanOrEqual lesser or equal
func (num *Number) LesserThanOrEqual(otherNum *Number) (bool, error) {
	if num.Kind != otherNum.Kind {
		return false, fmt.Errorf("can't compare %v with %v number", otherNum.Kind, num.Kind)
	}

	greaterThan := false

	switch num.Value.(type) {
	case Int8:
		greaterThan = num.Value.(Int8) <= otherNum.Value.(Int8)
	case Int16:
		greaterThan = num.Value.(Int16) <= otherNum.Value.(Int16)
	case Int32:
		greaterThan = num.Value.(Int32) <= otherNum.Value.(Int32)
	case Int64:
		greaterThan = num.Value.(Int64) <= otherNum.Value.(Int64)
	case Uint8:
		greaterThan = num.Value.(Uint8) <= otherNum.Value.(Uint8)
	case Uint16:
		greaterThan = num.Value.(Uint16) <= otherNum.Value.(Uint16)
	case Uint32:
		greaterThan = num.Value.(Uint32) <= otherNum.Value.(Uint32)
	case Uint64:
		greaterThan = num.Value.(Uint64) <= otherNum.Value.(Uint64)
	case Float32:
		greaterThan = num.Value.(Float32) <= otherNum.Value.(Float32)
	case Float64:
		greaterThan = num.Value.(Float64) <= otherNum.Value.(Float64)
	}

	return greaterThan, nil
}

// LesserThan lesser than
func (num *Number) LesserThan(otherNum *Number) (bool, error) {
	if num.Kind != otherNum.Kind {
		return false, fmt.Errorf("can't compare %v with %v number", otherNum.Kind, num.Kind)
	}

	greaterThan := false

	switch num.Value.(type) {
	case Int8:
		greaterThan = num.Value.(Int8) < otherNum.Value.(Int8)
	case Int16:
		greaterThan = num.Value.(Int16) < otherNum.Value.(Int16)
	case Int32:
		greaterThan = num.Value.(Int32) < otherNum.Value.(Int32)
	case Int64:
		greaterThan = num.Value.(Int64) < otherNum.Value.(Int64)
	case Uint8:
		greaterThan = num.Value.(Uint8) < otherNum.Value.(Uint8)
	case Uint16:
		greaterThan = num.Value.(Uint16) < otherNum.Value.(Uint16)
	case Uint32:
		greaterThan = num.Value.(Uint32) < otherNum.Value.(Uint32)
	case Uint64:
		greaterThan = num.Value.(Uint64) < otherNum.Value.(Uint64)
	case Float32:
		greaterThan = num.Value.(Float32) < otherNum.Value.(Float32)
	case Float64:
		greaterThan = num.Value.(Float64) < otherNum.Value.(Float64)
	}

	return greaterThan, nil
}

// GreaterThan greater than
func (num *Number) GreaterThan(otherNum *Number) (bool, error) {
	if num.Kind != otherNum.Kind {
		return false, fmt.Errorf("can't compare %v with %v number", otherNum.Kind, num.Kind)
	}

	greaterThan := false

	switch num.Value.(type) {
	case Int8:
		greaterThan = num.Value.(Int8) > otherNum.Value.(Int8)
	case Int16:
		greaterThan = num.Value.(Int16) > otherNum.Value.(Int16)
	case Int32:
		greaterThan = num.Value.(Int32) > otherNum.Value.(Int32)
	case Int64:
		greaterThan = num.Value.(Int64) > otherNum.Value.(Int64)
	case Uint8:
		greaterThan = num.Value.(Uint8) > otherNum.Value.(Uint8)
	case Uint16:
		greaterThan = num.Value.(Uint16) > otherNum.Value.(Uint16)
	case Uint32:
		greaterThan = num.Value.(Uint32) > otherNum.Value.(Uint32)
	case Uint64:
		greaterThan = num.Value.(Uint64) > otherNum.Value.(Uint64)
	case Float32:
		greaterThan = num.Value.(Float32) > otherNum.Value.(Float32)
	case Float64:
		greaterThan = num.Value.(Float64) > otherNum.Value.(Float64)
	}

	return greaterThan, nil
}

// GreaterThanOrEqual greater than or equal
func (num *Number) GreaterThanOrEqual(otherNum *Number) (bool, error) {
	if num.Kind != otherNum.Kind {
		return false, fmt.Errorf("can't compare %v with %v number", otherNum.Kind, num.Kind)
	}

	greaterThan := false

	switch num.Value.(type) {
	case Int8:
		greaterThan = num.Value.(Int8) >= otherNum.Value.(Int8)
	case Int16:
		greaterThan = num.Value.(Int16) >= otherNum.Value.(Int16)
	case Int32:
		greaterThan = num.Value.(Int32) >= otherNum.Value.(Int32)
	case Int64:
		greaterThan = num.Value.(Int64) >= otherNum.Value.(Int64)
	case Uint8:
		greaterThan = num.Value.(Uint8) >= otherNum.Value.(Uint8)
	case Uint16:
		greaterThan = num.Value.(Uint16) >= otherNum.Value.(Uint16)
	case Uint32:
		greaterThan = num.Value.(Uint32) >= otherNum.Value.(Uint32)
	case Uint64:
		greaterThan = num.Value.(Uint64) >= otherNum.Value.(Uint64)
	case Float32:
		greaterThan = num.Value.(Float32) >= otherNum.Value.(Float32)
	case Float64:
		greaterThan = num.Value.(Float64) >= otherNum.Value.(Float64)
	}

	return greaterThan, nil
}

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
		newNum.Value = num.Value.(Int8) * otherNum.Value.(Int8)
		newNum.Kind = reflect.Int8
	case Int16:
		newNum.Value = num.Value.(Int16) * otherNum.Value.(Int16)
		newNum.Kind = reflect.Int16
	case Int32:
		newNum.Value = num.Value.(Int32) * otherNum.Value.(Int32)
		newNum.Kind = reflect.Int32
	case Int64:
		newNum.Value = num.Value.(Int64) * otherNum.Value.(Int64)
		newNum.Kind = reflect.Int64
	case Uint8:
		newNum.Value = num.Value.(Uint8) * otherNum.Value.(Uint8)
		newNum.Kind = reflect.Uint8
	case Uint16:
		newNum.Value = num.Value.(Uint16) * otherNum.Value.(Uint16)
		newNum.Kind = reflect.Uint16
	case Uint32:
		newNum.Value = num.Value.(Uint32) * otherNum.Value.(Uint32)
		newNum.Kind = reflect.Uint32
	case Uint64:
		newNum.Value = num.Value.(Uint64) * otherNum.Value.(Uint64)
		newNum.Kind = reflect.Uint64
	case Float32:
		newNum.Value = num.Value.(Float32) * otherNum.Value.(Float32)
		newNum.Kind = reflect.Float32
	case Float64:
		newNum.Value = num.Value.(Float64) * otherNum.Value.(Float64)
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

// Max of two numbers
func (num *Number) Max(otherNum *Number) (*Number, error) {
	if num.Kind != otherNum.Kind {
		return nil, fmt.Errorf("can't MAX %v with %v number", num.Kind, otherNum.Kind)
	}

	max := num

	switch num.Value.(type) {
	case Int8:
		if num.Value.(Int8) < otherNum.Value.(Int8) {
			max = otherNum
		}
	case Int16:
		if num.Value.(Int16) < otherNum.Value.(Int16) {
			max = otherNum
		}
	case Int32:
		if num.Value.(Int32) < otherNum.Value.(Int32) {
			max = otherNum
		}
	case Int64:
		if num.Value.(Int64) < otherNum.Value.(Int64) {
			max = otherNum
		}
	case Uint8:
		if num.Value.(Uint8) < otherNum.Value.(Uint8) {
			max = otherNum
		}
	case Uint16:
		if num.Value.(Uint16) < otherNum.Value.(Uint16) {
			max = otherNum
		}
	case Uint32:
		if num.Value.(Uint32) < otherNum.Value.(Uint32) {
			max = otherNum
		}
	case Uint64:
		if num.Value.(Uint64) < otherNum.Value.(Uint64) {
			max = otherNum
		}
	case Float32:
		if num.Value.(Float32) < otherNum.Value.(Float32) {
			max = otherNum
		}
	case Float64:
		if num.Value.(Float64) < otherNum.Value.(Float64) {
			max = otherNum
		}
	}

	return max, nil
}

// Min of two numbers
func (num *Number) Min(otherNum *Number) (*Number, error) {
	if num.Kind != otherNum.Kind {
		return nil, fmt.Errorf("can't MIN %v with %v number", num.Kind, otherNum.Kind)
	}

	min := num

	switch num.Value.(type) {
	case Int8:
		if num.Value.(Int8) > otherNum.Value.(Int8) {
			min = otherNum
		}
	case Int16:
		if num.Value.(Int16) > otherNum.Value.(Int16) {
			min = otherNum
		}
	case Int32:
		if num.Value.(Int32) > otherNum.Value.(Int32) {
			min = otherNum
		}
	case Int64:
		if num.Value.(Int64) > otherNum.Value.(Int64) {
			min = otherNum
		}
	case Uint8:
		if num.Value.(Uint8) > otherNum.Value.(Uint8) {
			min = otherNum
		}
	case Uint16:
		if num.Value.(Uint16) > otherNum.Value.(Uint16) {
			min = otherNum
		}
	case Uint32:
		if num.Value.(Uint32) > otherNum.Value.(Uint32) {
			min = otherNum
		}
	case Uint64:
		if num.Value.(Uint64) > otherNum.Value.(Uint64) {
			min = otherNum
		}
	case Float32:
		if num.Value.(Float32) > otherNum.Value.(Float32) {
			min = otherNum
		}
	case Float64:
		if num.Value.(Float64) > otherNum.Value.(Float64) {
			min = otherNum
		}
	}

	return min, nil
}
