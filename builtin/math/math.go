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

func doubleFloat32MathFunc(obj1 types.Object, obj2 types.Object, name string, fun func(float32, float32) float32) (types.Object, error) {
	num1, ok := obj1.(*numbers.Number)
	if !ok {
		return nil, fmt.Errorf("%v only accepts numbers", name)
	}

	num2, ok := obj2.(*numbers.Number)
	if !ok {
		return nil, fmt.Errorf("%v only accepts numbers", name)
	}

	newNum := numbers.New(reflect.Float32)
	newNum.Value = numbers.Float32(fun(float32(num1.Float64Value()), float32(num2.Float64Value())))

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

// Float32Bits float32bits
func Float32Bits(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("FLOAT32-BITS only accepts numbers")
	}

	if num.Kind != reflect.Float32 {
		return nil, errors.New("FLOAT32-BITS expected a float32 number")
	}

	newNum := numbers.New(reflect.Uint32)
	newNum.Value = numbers.Uint32(math.Float32bits(float32(num.Value.(numbers.Float32))))

	return newNum, nil
}

// Float32FromBits float32frombits
func Float32FromBits(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("FLOAT32-FROM-BITS only accepts numbers")
	}

	if num.Kind != reflect.Uint32 {
		return nil, errors.New("FLOAT32-FROM-BITS expected a uint32 number")
	}

	newNum := numbers.New(reflect.Float32)
	newNum.Value = numbers.Float32(math.Float32frombits(uint32(num.Value.(numbers.Uint32))))

	return newNum, nil
}

// Float64Bits float64bits
func Float64Bits(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("FLOAT64-BITS only accepts numbers")
	}

	if num.Kind != reflect.Float64 {
		return nil, errors.New("FLOAT64-BITS expected a float64 number")
	}

	newNum := numbers.New(reflect.Uint64)
	newNum.Value = numbers.Uint64(math.Float64bits(float64(num.Value.(numbers.Float64))))

	return newNum, nil
}

// Float64FromBits float64frombits
func Float64FromBits(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("FLOAT64-FROM-BITS only accepts numbers")
	}

	if num.Kind != reflect.Uint64 {
		return nil, errors.New("FLOAT64-FROM-BITS expected a uint64 number")
	}

	newNum := numbers.New(reflect.Float64)
	newNum.Value = numbers.Float64(math.Float64frombits(uint64(num.Value.(numbers.Uint64))))

	return newNum, nil
}

// Floor floor
func Floor(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "FLOOR", math.Floor)
}

// Frexp frexp
func Frexp(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("FREXP only accepts numbers")
	}

	fl := num.Float64Value()
	frac, exp := math.Frexp(fl)

	n1 := numbers.New(reflect.Float64)
	n1.Value = numbers.Float64(frac)

	n2 := numbers.New(reflect.Int64)
	n2.Value = numbers.Int64(exp)

	return cons.ListFromSlice([]types.Object{n1, n2}), nil
}

// Gamma gamma
func Gamma(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "GAMMA", math.Gamma)
}

// Hypot hypot
func Hypot(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return doubleFloat64MathFunc(args.Car, args.Cdr.(*cons.Cons).Car, "HYPOT", math.Hypot)
}

// Ilogb ilogb
func Ilogb(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("ILOGB only accepts numbers")
	}

	fl := num.Float64Value()
	i := math.Ilogb(fl)

	n := numbers.New(reflect.Int64)
	n.Value = numbers.Int64(i)

	return n, nil
}

// Inf inf
func Inf(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("INF only accepts numbers")
	}

	fl := num.Int64Value()

	n := numbers.New(reflect.Float64)
	n.Value = numbers.Float64(math.Inf(int(fl)))

	return n, nil
}

// IsInf isinf
func IsInf(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num1, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("IS-INF only accepts numbers")
	}

	num2, ok := args.Cdr.(*cons.Cons).Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("IS-INF only accepts numbers")
	}

	fl := num1.Float64Value()
	sign := num2.Int64Value()

	if math.IsInf(fl, int(sign)) {
		return types.T, nil
	}

	return types.NIL, nil
}

// IsNaN isnan
func IsNaN(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("IS-NAN only accepts numbers")
	}

	fl := num.Float64Value()

	if math.IsNaN(fl) {
		return types.T, nil
	}

	return types.NIL, nil
}

// J0 j0
func J0(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "J0", math.J0)
}

// J1 j1
func J1(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "J1", math.J1)
}

// Jn jn
func Jn(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num1, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("JN only accepts numbers")
	}

	num2, ok := args.Cdr.(*cons.Cons).Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("JN only accepts numbers")
	}

	n := num1.Int64Value()
	x := num2.Float64Value()

	newNum := numbers.New(reflect.Float64)
	newNum.Value = numbers.Float64(math.Jn(int(n), x))

	return newNum, nil
}

// Ldexp ldexp
func Ldexp(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num1, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("LDEXP only accepts numbers")
	}

	num2, ok := args.Cdr.(*cons.Cons).Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("LDEXP only accepts numbers")
	}

	frac := num1.Float64Value()
	exp := num2.Int64Value()

	newNum := numbers.New(reflect.Float64)
	newNum.Value = numbers.Float64(math.Ldexp(frac, int(exp)))

	return newNum, nil
}

// Lgamma lgamma
func Lgamma(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("LGAMMA only accepts numbers")
	}

	fl := num.Float64Value()
	lgamma, sign := math.Lgamma(fl)

	n1 := numbers.New(reflect.Float64)
	n1.Value = numbers.Float64(lgamma)

	n2 := numbers.New(reflect.Int64)
	n2.Value = numbers.Int64(sign)

	return cons.ListFromSlice([]types.Object{n1, n2}), nil
}

// Log log
func Log(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "LOG", math.Log)
}

// Log10 log10
func Log10(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "LOG10", math.Log10)
}

// Log1p log1p
func Log1p(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "LOG1P", math.Log1p)
}

// Log2 log2
func Log2(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "LOG2", math.Log2)
}

// Logb logb
func Logb(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "LOGB", math.Logb)
}

// Modf modf
func Modf(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("MODF only accepts numbers")
	}

	fl := num.Float64Value()
	i, frac := math.Modf(fl)

	n1 := numbers.New(reflect.Float64)
	n1.Value = numbers.Float64(i)

	n2 := numbers.New(reflect.Float64)
	n2.Value = numbers.Float64(frac)

	return cons.ListFromSlice([]types.Object{n1, n2}), nil
}

// NaN nan
func NaN(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num := numbers.New(reflect.Float64)
	num.Value = numbers.Float64(math.NaN())
	return num, nil
}

// Nextafter nextafter
func Nextafter(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return doubleFloat64MathFunc(args.Car, args.Cdr.(*cons.Cons).Car, "NEXT-AFTER", math.Nextafter)
}

// Nextafter32 nextafter32
func Nextafter32(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return doubleFloat32MathFunc(args.Car, args.Cdr.(*cons.Cons).Car, "NEXT-AFTER32", math.Nextafter32)
}

// Pow pow
func Pow(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return doubleFloat64MathFunc(args.Car, args.Cdr.(*cons.Cons).Car, "POW", math.Pow)
}

// Pow10 pow10
func Pow10(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("POW10 only accepts numbers")
	}

	n := num.Int64Value()

	newNum := numbers.New(reflect.Float64)
	newNum.Value = numbers.Float64(math.Pow10(int(n)))

	return newNum, nil
}

// Remainder remainder
func Remainder(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return doubleFloat64MathFunc(args.Car, args.Cdr.(*cons.Cons).Car, "REMAINDER", math.Remainder)
}

// Signbit signbit
func Signbit(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("SIGNBIT only accepts numbers")
	}

	if math.Signbit(num.Float64Value()) {
		return types.T, nil
	}

	return types.NIL, nil
}

// Sin sin
func Sin(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "SIN", math.Sin)
}

// Sincos sincos
func Sincos(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("SINCOS only accepts numbers")
	}

	fl := num.Float64Value()
	s, c := math.Sincos(fl)

	n1 := numbers.New(reflect.Float64)
	n1.Value = numbers.Float64(s)

	n2 := numbers.New(reflect.Float64)
	n2.Value = numbers.Int64(c)

	return cons.ListFromSlice([]types.Object{n1, n2}), nil
}

// Sinh sinh
func Sinh(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "SINH", math.Sinh)
}

// Sqrt sqrt
func Sqrt(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "SQRT", math.Sqrt)
}

// Tan tan
func Tan(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "TAN", math.Tan)
}

// Tanh tanh
func Tanh(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "TANH", math.Tanh)
}

// Trunc trunc
func Trunc(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "TRUNC", math.Trunc)
}

// Y0 y0
func Y0(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "Y0", math.Y0)
}

// Y1 y1
func Y1(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return singleFloat64MathFunc(args.Car, "Y1", math.Y1)
}

// Yn yn
func Yn(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	num1, ok := args.Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("YN only accepts numbers")
	}

	num2, ok := args.Cdr.(*cons.Cons).Car.(*numbers.Number)
	if !ok {
		return nil, errors.New("YN only accepts numbers")
	}

	n := num1.Int64Value()
	x := num2.Float64Value()

	newNum := numbers.New(reflect.Float64)
	newNum.Value = numbers.Float64(math.Yn(int(n), x))

	return newNum, nil
}
