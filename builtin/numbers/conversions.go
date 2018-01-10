package numbers

import (
	"errors"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
	"github.com/almerlucke/glisp/types/numbers"
)

// Int8 converts a number to int8 type
func Int8(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("INT8 expected a number as first argument")
	}

	return num.Int8(), nil
}

// CreateBuiltinInt8 creates an int8 function object
func CreateBuiltinInt8() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Int8, 1, true)
}

// Int16 converts a number to int16 type
func Int16(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("INT16 expected a number as first argument")
	}

	return num.Int16(), nil
}

// CreateBuiltinInt16 creates an int16 function object
func CreateBuiltinInt16() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Int16, 1, true)
}

// Int32 converts a number to int32 type
func Int32(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("INT32 expected a number as first argument")
	}

	return num.Int32(), nil
}

// CreateBuiltinInt32 creates an int32 function object
func CreateBuiltinInt32() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Int32, 1, true)
}

// Int64 converts a number to int64 type
func Int64(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("INT64 expected a number as first argument")
	}

	return num.Int64(), nil
}

// CreateBuiltinInt64 creates an int64 function object
func CreateBuiltinInt64() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Int64, 1, true)
}

// Uint8 converts a number to Uint8 type
func Uint8(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("UINT8 expected a number as first argument")
	}

	return num.Uint8(), nil
}

// CreateBuiltinUint8 creates an uint8 function object
func CreateBuiltinUint8() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Uint8, 1, true)
}

// Uint16 converts a number to uint16 type
func Uint16(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("UINT16 expected a number as first argument")
	}

	return num.Uint16(), nil
}

// CreateBuiltinUint16 creates an uint16 function object
func CreateBuiltinUint16() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Uint16, 1, true)
}

// Uint32 converts a number to uint32 type
func Uint32(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("UINT32 expected a number as first argument")
	}

	return num.Uint32(), nil
}

// CreateBuiltinUint32 creates an uint32 function object
func CreateBuiltinUint32() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Uint32, 1, true)
}

// Uint64 converts a number to uint64 type
func Uint64(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("UINT64 expected a number as first argument")
	}

	return num.Uint64(), nil
}

// CreateBuiltinUint64 creates an uint64 function object
func CreateBuiltinUint64() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Uint64, 1, true)
}

// Float32 converts a number to float32 type
func Float32(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("FLOAT32 expected a number as first argument")
	}

	return num.Float32(), nil
}

// CreateBuiltinFloat32 creates a float32 function object
func CreateBuiltinFloat32() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Float32, 1, true)
}

// Float64 converts a number to float64 type
func Float64(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("FLOAT64 expected a number as first argument")
	}

	return num.Float64(), nil
}

// CreateBuiltinFloat64 creates a float64 function object
func CreateBuiltinFloat64() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(Float64, 1, true)
}
