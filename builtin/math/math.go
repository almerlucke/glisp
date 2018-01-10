package math

import (
	"errors"
	"fmt"
	"math"
	"reflect"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/numbers"
)

// genericSingleArgMathFunc generic math function
func genericSingleArgMathFunc(num *numbers.Number, fun func(float64) float64, env environment.Environment, context interface{}) (types.Object, error) {
	switch num.Kind {
	case reflect.Int8:
		newNum := numbers.New(num.Kind)
		newNum.Value = numbers.Int8(fun(float64(num.Value.(numbers.Int8))))
		num = newNum
	case reflect.Int16:
		newNum := numbers.New(num.Kind)
		newNum.Value = numbers.Int16(fun(float64(num.Value.(numbers.Int16))))
		num = newNum
	case reflect.Int32:
		newNum := numbers.New(num.Kind)
		newNum.Value = numbers.Int32(fun(float64(num.Value.(numbers.Int32))))
		num = newNum
	case reflect.Int64:
		newNum := numbers.New(num.Kind)
		newNum.Value = numbers.Int64(fun(float64(num.Value.(numbers.Int64))))
		num = newNum
	case reflect.Uint8:
		newNum := numbers.New(num.Kind)
		newNum.Value = numbers.Uint8(fun(float64(num.Value.(numbers.Uint8))))
		num = newNum
	case reflect.Uint16:
		newNum := numbers.New(num.Kind)
		newNum.Value = numbers.Uint16(fun(float64(num.Value.(numbers.Uint16))))
		num = newNum
	case reflect.Uint32:
		newNum := numbers.New(num.Kind)
		newNum.Value = numbers.Uint32(fun(float64(num.Value.(numbers.Uint32))))
		num = newNum
	case reflect.Uint64:
		newNum := numbers.New(num.Kind)
		newNum.Value = numbers.Uint64(fun(float64(num.Value.(numbers.Uint64))))
		num = newNum
	case reflect.Float32:
		newNum := numbers.New(num.Kind)
		newNum.Value = numbers.Float32(fun(float64(num.Value.(numbers.Float32))))
		num = newNum
	case reflect.Float64:
		newNum := numbers.New(num.Kind)
		newNum.Value = numbers.Float64(fun(float64(num.Value.(numbers.Float64))))
		num = newNum
	}

	return num, nil
}

func singleFloat64MathFunc(obj types.Object, name string, fun func(float64) float64) (types.Object, error) {
	num, ok := obj.(*numbers.Number)
	if !ok {
		return nil, fmt.Errorf("%v only accepts numbers", name)
	}

	newNum := numbers.New(reflect.Float64)
	newNum.Value = numbers.Float64(fun(num.Float64Value()))

	return newNum, nil
}

func doubleFloat64MathFunc(obj1 types.Object, obj2 types.Object, name string, fun func(float64, float64) float64) (types.Object, error) {
	num1, ok := obj1.(*numbers.Number)
	if !ok {
		return nil, fmt.Errorf("%v only accepts numbers", name)
	}

	num2, ok := obj2.(*numbers.Number)
	if !ok {
		return nil, fmt.Errorf("%v only accepts numbers", name)
	}

	newNum := numbers.New(reflect.Float64)
	newNum.Value = numbers.Float64(fun(float64(num1.Float64Value()), float64(num2.Float64Value())))

	return newNum, nil
}

// Abs get absolute value
func Abs(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("ABS only accepts numbers")
	}

	return genericSingleArgMathFunc(num, math.Abs, env, context)
}

// Acos acos
func Acos(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "ACOS", math.Acos)
}

// Acosh acosh
func Acosh(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "ACOSH", math.Acosh)
}

// Asin asin
func Asin(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "ASIN", math.Asin)
}

// Asinh asinh
func Asinh(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "ASINH", math.Asinh)
}

// Atan atan
func Atan(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "ATAN", math.Atan)
}

// Atan2 atan2
func Atan2(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return doubleFloat64MathFunc(args.Car, args.Cdr.(*cons.Cons).Car, "ATAN2", math.Atan2)
}

// Atanh atanh
func Atanh(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "ATANH", math.Atanh)
}

// Cbrt cbrt
func Cbrt(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "CBRT", math.Cbrt)
}

// Ceil ceil
func Ceil(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "CEIL", math.Ceil)
}

// Copysign copysign
func Copysign(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return doubleFloat64MathFunc(args.Car, args.Cdr.(*cons.Cons).Car, "COPYSIGN", math.Copysign)
}

// Cos cos
func Cos(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "COS", math.Cos)
}

// Cosh cosh
func Cosh(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "COSH", math.Cosh)
}

// Dim dim
func Dim(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return doubleFloat64MathFunc(args.Car, args.Cdr.(*cons.Cons).Car, "DIM", math.Dim)
}

// Erf erf
func Erf(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "ERF", math.Erf)
}

// Erfc erfc
func Erfc(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "ERFC", math.Erfc)
}

// Exp exp
func Exp(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "EXP", math.Exp)
}

// Exp2 exp2
func Exp2(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "EXP2", math.Exp2)
}

// Expm1 expm1
func Expm1(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "EXPM1", math.Expm1)
}
