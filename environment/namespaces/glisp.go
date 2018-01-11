package namespaces

import (
	"github.com/almerlucke/glisp/builtin"
	"github.com/almerlucke/glisp/builtin/loops"
	"github.com/almerlucke/glisp/builtin/numbers"
	"github.com/almerlucke/glisp/globals/symbols"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/interfaces/namespace"
	"github.com/almerlucke/glisp/types/namespaces"
)

// CreateGlispNamespace create the main GLisp namespace
func CreateGlispNamespace(env environment.Environment) namespace.Namespace {
	glispNS := namespaces.NewNamespace("GLISP", false)

	glispNS.Add(symbols.NILSymbol, true)
	glispNS.Add(symbols.TSymbol, true)
	glispNS.Add(symbols.AndRestSymbol, true)
	glispNS.Add(symbols.SelfSymbol, true)
	glispNS.Add(symbols.BackquoteSymbol, true)
	glispNS.Add(symbols.CloseParenthesisSymbol, true)
	glispNS.Add(symbols.DotSymbol, true)
	glispNS.Add(symbols.QuoteSymbol, true)
	glispNS.Add(symbols.SpliceSymbol, true)
	glispNS.Add(symbols.UnquoteSymbol, true)

	env.AddGlobalBinding(symbols.QuoteSymbol, builtin.CreateBuiltinQuote())
	env.AddGlobalBinding(symbols.BackquoteSymbol, builtin.CreateBuiltinBackquote())
	env.AddGlobalBinding(symbols.UnquoteSymbol, builtin.CreateBuiltinUnquote())
	env.AddGlobalBinding(symbols.SpliceSymbol, builtin.CreateBuiltinUnquote())

	env.AddGlobalBinding(glispNS.DefineSymbol("LIST", true, nil, true), builtin.CreateBuiltinList())
	env.AddGlobalBinding(glispNS.DefineSymbol("CDR", true, nil, true), builtin.CreateBuiltinCdr())
	env.AddGlobalBinding(glispNS.DefineSymbol("CAR", true, nil, true), builtin.CreateBuiltinCar())
	env.AddGlobalBinding(glispNS.DefineSymbol("CONS", true, nil, true), builtin.CreateBuiltinCons())
	env.AddGlobalBinding(glispNS.DefineSymbol("LAMBDA", true, nil, true), builtin.CreateBuiltinLambda())
	env.AddGlobalBinding(glispNS.DefineSymbol("MACRO", true, nil, true), builtin.CreateBuiltinMacro())
	env.AddGlobalBinding(glispNS.DefineSymbol("GENSYM", true, nil, true), builtin.CreateBuiltinGensym())
	env.AddGlobalBinding(glispNS.DefineSymbol("PRINT", true, nil, true), builtin.CreateBuiltinPrint())
	env.AddGlobalBinding(glispNS.DefineSymbol("EXIT", true, nil, true), builtin.CreateBuiltinExit())
	env.AddGlobalBinding(glispNS.DefineSymbol("RETURN", true, nil, true), builtin.CreateBuiltinReturn())
	env.AddGlobalBinding(glispNS.DefineSymbol("LOAD", true, nil, true), builtin.CreateBuiltinLoad())
	env.AddGlobalBinding(glispNS.DefineSymbol("VAR", true, nil, true), builtin.CreateBuiltinVar())
	env.AddGlobalBinding(glispNS.DefineSymbol("=", true, nil, true), builtin.CreateBuiltinAssign())
	env.AddGlobalBinding(glispNS.DefineSymbol("SCOPE", true, nil, true), builtin.CreateBuiltinScope())
	env.AddGlobalBinding(glispNS.DefineSymbol("EVAL", true, nil, true), builtin.CreateBuiltinEval())
	env.AddGlobalBinding(glispNS.DefineSymbol("ELT", true, nil, true), builtin.CreateBuiltinElt())
	env.AddGlobalBinding(glispNS.DefineSymbol("ARRAY", true, nil, true), builtin.CreateBuiltinArray())
	env.AddGlobalBinding(glispNS.DefineSymbol("MAKE-ARRAY", true, nil, true), builtin.CreateBuiltinMakeArray())
	env.AddGlobalBinding(glispNS.DefineSymbol("IF", true, nil, true), builtin.CreateBuiltinIf())
	env.AddGlobalBinding(glispNS.DefineSymbol("DO", true, nil, true), builtin.CreateBuiltinDo())
	env.AddGlobalBinding(glispNS.DefineSymbol("TRY", true, nil, true), builtin.CreateBuiltinTry())
	env.AddGlobalBinding(glispNS.DefineSymbol("THROW", true, nil, true), builtin.CreateBuiltinThrow())
	env.AddGlobalBinding(glispNS.DefineSymbol("DICTIONARY", true, nil, true), builtin.CreateBuiltinDictionary())
	env.AddGlobalBinding(glispNS.DefineSymbol("AND", true, nil, true), builtin.CreateBuiltinAnd())
	env.AddGlobalBinding(glispNS.DefineSymbol("OR", true, nil, true), builtin.CreateBuiltinOr())
	env.AddGlobalBinding(glispNS.DefineSymbol("NOT", true, nil, true), builtin.CreateBuiltinNot())
	env.AddGlobalBinding(glispNS.DefineSymbol("NAMESPACE", true, nil, true), builtin.CreateBuiltinNamespace())
	env.AddGlobalBinding(glispNS.DefineSymbol("IN-NAMESPACE", true, nil, true), builtin.CreateBuiltinInNamespace())
	env.AddGlobalBinding(glispNS.DefineSymbol("USE-NAMESPACE", true, nil, true), builtin.CreateBuiltinUseNamespace())
	env.AddGlobalBinding(glispNS.DefineSymbol("MAP", true, nil, true), builtin.CreateBuiltinMap())

	env.AddGlobalBinding(glispNS.DefineSymbol("WHILE", true, nil, true), loops.CreateBuiltinWhile())
	env.AddGlobalBinding(glispNS.DefineSymbol("BREAK", true, nil, true), loops.CreateBuiltinBreak())

	env.AddGlobalBinding(glispNS.DefineSymbol("INT8", true, nil, true), numbers.CreateBuiltinInt8())
	env.AddGlobalBinding(glispNS.DefineSymbol("INT16", true, nil, true), numbers.CreateBuiltinInt16())
	env.AddGlobalBinding(glispNS.DefineSymbol("INT32", true, nil, true), numbers.CreateBuiltinInt32())
	env.AddGlobalBinding(glispNS.DefineSymbol("INT64", true, nil, true), numbers.CreateBuiltinInt64())
	env.AddGlobalBinding(glispNS.DefineSymbol("UINT8", true, nil, true), numbers.CreateBuiltinUint8())
	env.AddGlobalBinding(glispNS.DefineSymbol("UINT16", true, nil, true), numbers.CreateBuiltinUint16())
	env.AddGlobalBinding(glispNS.DefineSymbol("UINT32", true, nil, true), numbers.CreateBuiltinUint32())
	env.AddGlobalBinding(glispNS.DefineSymbol("UINT64", true, nil, true), numbers.CreateBuiltinUint64())
	env.AddGlobalBinding(glispNS.DefineSymbol("FLOAT32", true, nil, true), numbers.CreateBuiltinFloat32())
	env.AddGlobalBinding(glispNS.DefineSymbol("FLOAT64", true, nil, true), numbers.CreateBuiltinFloat64())

	env.AddGlobalBinding(glispNS.DefineSymbol("+", true, nil, true), numbers.CreateBuiltinNumberAdd())
	env.AddGlobalBinding(glispNS.DefineSymbol("-", true, nil, true), numbers.CreateBuiltinNumberSubtract())
	env.AddGlobalBinding(glispNS.DefineSymbol("*", true, nil, true), numbers.CreateBuiltinNumberMultiply())
	env.AddGlobalBinding(glispNS.DefineSymbol("/", true, nil, true), numbers.CreateBuiltinNumberDivide())
	env.AddGlobalBinding(glispNS.DefineSymbol("%", true, nil, true), numbers.CreateBuiltinNumberModulo())
	env.AddGlobalBinding(glispNS.DefineSymbol(">", true, nil, true), numbers.CreateBuiltinNumberGreaterThan())
	env.AddGlobalBinding(glispNS.DefineSymbol(">=", true, nil, true), numbers.CreateBuiltinNumberGreaterThanOrEqual())
	env.AddGlobalBinding(glispNS.DefineSymbol("<", true, nil, true), numbers.CreateBuiltinNumberLesserThan())
	env.AddGlobalBinding(glispNS.DefineSymbol("<=", true, nil, true), numbers.CreateBuiltinNumberLesserThanOrEqual())
	env.AddGlobalBinding(glispNS.DefineSymbol("MAX", true, nil, true), numbers.CreateBuiltinNumberMax())
	env.AddGlobalBinding(glispNS.DefineSymbol("MIN", true, nil, true), numbers.CreateBuiltinNumberMin())

	return glispNS
}
