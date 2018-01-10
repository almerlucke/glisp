package namespaces

import (
	goMath "math"

	"github.com/almerlucke/glisp/builtin/math"
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/interfaces/namespace"
	"github.com/almerlucke/glisp/types/functions"
	"github.com/almerlucke/glisp/types/namespaces"
	"github.com/almerlucke/glisp/types/numbers"
)

// CreateMathNamespace create the math namespace
func CreateMathNamespace(env environment.Environment) namespace.Namespace {
	mathNS := namespaces.NewNamespace("MATH", false)

	env.AddGlobalBinding(mathNS.DefineSymbol("E", true, nil, true), numbers.NewFloat64(goMath.E))
	env.AddGlobalBinding(mathNS.DefineSymbol("PI", true, nil, true), numbers.NewFloat64(goMath.Pi))
	env.AddGlobalBinding(mathNS.DefineSymbol("PHI", true, nil, true), numbers.NewFloat64(goMath.Phi))
	env.AddGlobalBinding(mathNS.DefineSymbol("SQRT2", true, nil, true), numbers.NewFloat64(goMath.Sqrt2))
	env.AddGlobalBinding(mathNS.DefineSymbol("SQRTE", true, nil, true), numbers.NewFloat64(goMath.SqrtE))
	env.AddGlobalBinding(mathNS.DefineSymbol("SQRTPI", true, nil, true), numbers.NewFloat64(goMath.SqrtPi))
	env.AddGlobalBinding(mathNS.DefineSymbol("SQRTPHI", true, nil, true), numbers.NewFloat64(goMath.SqrtPhi))
	env.AddGlobalBinding(mathNS.DefineSymbol("LN2", true, nil, true), numbers.NewFloat64(goMath.Ln2))
	env.AddGlobalBinding(mathNS.DefineSymbol("LOG2E", true, nil, true), numbers.NewFloat64(goMath.Log2E))
	env.AddGlobalBinding(mathNS.DefineSymbol("LN10", true, nil, true), numbers.NewFloat64(goMath.Ln10))
	env.AddGlobalBinding(mathNS.DefineSymbol("LOG10E", true, nil, true), numbers.NewFloat64(goMath.Log10E))

	env.AddGlobalBinding(mathNS.DefineSymbol("MAX-FLOAT32", true, nil, true), numbers.NewFloat32(goMath.MaxFloat32))
	env.AddGlobalBinding(mathNS.DefineSymbol("MIN-FLOAT32", true, nil, true), numbers.NewFloat32(goMath.SmallestNonzeroFloat32))
	env.AddGlobalBinding(mathNS.DefineSymbol("MAX-FLOAT64", true, nil, true), numbers.NewFloat64(goMath.MaxFloat64))
	env.AddGlobalBinding(mathNS.DefineSymbol("MIN-FLOAT64", true, nil, true), numbers.NewFloat64(goMath.SmallestNonzeroFloat64))

	env.AddGlobalBinding(mathNS.DefineSymbol("MAX-INT8", true, nil, true), numbers.NewInt8(goMath.MaxInt8))
	env.AddGlobalBinding(mathNS.DefineSymbol("MIN-INT8", true, nil, true), numbers.NewInt8(goMath.MinInt8))
	env.AddGlobalBinding(mathNS.DefineSymbol("MAX-INT16", true, nil, true), numbers.NewInt16(goMath.MaxInt16))
	env.AddGlobalBinding(mathNS.DefineSymbol("MIN-INT16", true, nil, true), numbers.NewInt16(goMath.MinInt16))
	env.AddGlobalBinding(mathNS.DefineSymbol("MAX-INT32", true, nil, true), numbers.NewInt32(goMath.MaxInt32))
	env.AddGlobalBinding(mathNS.DefineSymbol("MIN-INT32", true, nil, true), numbers.NewInt32(goMath.MinInt32))
	env.AddGlobalBinding(mathNS.DefineSymbol("MAX-INT64", true, nil, true), numbers.NewInt64(goMath.MaxInt64))
	env.AddGlobalBinding(mathNS.DefineSymbol("MIN-INT64", true, nil, true), numbers.NewInt64(goMath.MinInt64))

	env.AddGlobalBinding(mathNS.DefineSymbol("MAX-UINT8", true, nil, true), numbers.NewUint8(goMath.MaxUint8))
	env.AddGlobalBinding(mathNS.DefineSymbol("MAX-UINT16", true, nil, true), numbers.NewUint16(goMath.MaxUint16))
	env.AddGlobalBinding(mathNS.DefineSymbol("MAX-UINT32", true, nil, true), numbers.NewUint32(goMath.MaxUint32))
	env.AddGlobalBinding(mathNS.DefineSymbol("MAX-UINT64", true, nil, true), numbers.NewUint64(goMath.MaxUint64))

	env.AddGlobalBinding(mathNS.DefineSymbol("ABS", true, nil, true), functions.NewBuiltinFunction(math.Abs, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("ACOS", true, nil, true), functions.NewBuiltinFunction(math.Acos, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("ACOSH", true, nil, true), functions.NewBuiltinFunction(math.Acosh, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("ASIN", true, nil, true), functions.NewBuiltinFunction(math.Asin, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("ASINH", true, nil, true), functions.NewBuiltinFunction(math.Asinh, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("ATAN", true, nil, true), functions.NewBuiltinFunction(math.Atan, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("ATAN2", true, nil, true), functions.NewBuiltinFunction(math.Atan2, 2, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("ATANH", true, nil, true), functions.NewBuiltinFunction(math.Atanh, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("CBRT", true, nil, true), functions.NewBuiltinFunction(math.Cbrt, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("CEIL", true, nil, true), functions.NewBuiltinFunction(math.Ceil, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("COPYSIGN", true, nil, true), functions.NewBuiltinFunction(math.Copysign, 2, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("COS", true, nil, true), functions.NewBuiltinFunction(math.Cos, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("COSH", true, nil, true), functions.NewBuiltinFunction(math.Cosh, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("DIM", true, nil, true), functions.NewBuiltinFunction(math.Dim, 2, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("ERF", true, nil, true), functions.NewBuiltinFunction(math.Erf, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("ERFC", true, nil, true), functions.NewBuiltinFunction(math.Erfc, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("EXP", true, nil, true), functions.NewBuiltinFunction(math.Exp, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("EXP2", true, nil, true), functions.NewBuiltinFunction(math.Exp2, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("EXPM1", true, nil, true), functions.NewBuiltinFunction(math.Expm1, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("FLOAT32-BITS", true, nil, true), functions.NewBuiltinFunction(math.Float32Bits, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("FLOAT32-FROM-BITS", true, nil, true), functions.NewBuiltinFunction(math.Float32FromBits, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("FLOAT64-BITS", true, nil, true), functions.NewBuiltinFunction(math.Float64Bits, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("FLOAT64-FROM-BITS", true, nil, true), functions.NewBuiltinFunction(math.Float64FromBits, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("FLOOR", true, nil, true), functions.NewBuiltinFunction(math.Floor, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("FREXP", true, nil, true), functions.NewBuiltinFunction(math.Frexp, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("GAMMA", true, nil, true), functions.NewBuiltinFunction(math.Gamma, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("HYPOT", true, nil, true), functions.NewBuiltinFunction(math.Hypot, 2, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("ILOGB", true, nil, true), functions.NewBuiltinFunction(math.Ilogb, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("INF", true, nil, true), functions.NewBuiltinFunction(math.Inf, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("IS-INF", true, nil, true), functions.NewBuiltinFunction(math.IsInf, 2, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("IS-NAN", true, nil, true), functions.NewBuiltinFunction(math.IsNaN, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("J0", true, nil, true), functions.NewBuiltinFunction(math.J0, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("J1", true, nil, true), functions.NewBuiltinFunction(math.J1, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("JN", true, nil, true), functions.NewBuiltinFunction(math.Jn, 2, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("LDEXP", true, nil, true), functions.NewBuiltinFunction(math.Ldexp, 2, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("LGAMMA", true, nil, true), functions.NewBuiltinFunction(math.Lgamma, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("LOG", true, nil, true), functions.NewBuiltinFunction(math.Log, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("LOG10", true, nil, true), functions.NewBuiltinFunction(math.Log10, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("LOG1P", true, nil, true), functions.NewBuiltinFunction(math.Log1p, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("LOG2", true, nil, true), functions.NewBuiltinFunction(math.Log2, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("LOGB", true, nil, true), functions.NewBuiltinFunction(math.Logb, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("MODF", true, nil, true), functions.NewBuiltinFunction(math.Modf, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("NAN", true, nil, true), functions.NewBuiltinFunction(math.NaN, 0, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("NEXT-AFTER", true, nil, true), functions.NewBuiltinFunction(math.Nextafter, 2, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("NEXT-AFTER32", true, nil, true), functions.NewBuiltinFunction(math.Nextafter32, 2, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("POW", true, nil, true), functions.NewBuiltinFunction(math.Pow, 2, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("POW10", true, nil, true), functions.NewBuiltinFunction(math.Pow10, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("REMAINDER", true, nil, true), functions.NewBuiltinFunction(math.Remainder, 2, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("SIGNBIT", true, nil, true), functions.NewBuiltinFunction(math.Signbit, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("SIN", true, nil, true), functions.NewBuiltinFunction(math.Sin, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("SINCOS", true, nil, true), functions.NewBuiltinFunction(math.Sincos, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("SINH", true, nil, true), functions.NewBuiltinFunction(math.Sinh, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("SQRT", true, nil, true), functions.NewBuiltinFunction(math.Sqrt, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("TAN", true, nil, true), functions.NewBuiltinFunction(math.Tan, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("TANH", true, nil, true), functions.NewBuiltinFunction(math.Tanh, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("TRUNC", true, nil, true), functions.NewBuiltinFunction(math.Trunc, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("Y0", true, nil, true), functions.NewBuiltinFunction(math.Y0, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("Y1", true, nil, true), functions.NewBuiltinFunction(math.Y1, 1, true))
	env.AddGlobalBinding(mathNS.DefineSymbol("YN", true, nil, true), functions.NewBuiltinFunction(math.Yn, 2, true))

	return mathNS
}
