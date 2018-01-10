package namespaces

import (
	"github.com/almerlucke/glisp/builtin/math"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/interfaces/namespace"
	"github.com/almerlucke/glisp/types/functions"
	"github.com/almerlucke/glisp/types/namespaces"
)

// CreateMathNamespace create the math namespace
func CreateMathNamespace(env environment.Environment) namespace.Namespace {
	mathNS := namespaces.NewNamespace("MATH", false)

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

	return mathNS
}
