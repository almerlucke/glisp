package environment

import (
	"github.com/almerlucke/glisp/interfaces/namespace"
	"github.com/almerlucke/glisp/scope"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/symbols"
)

// Environment implements the currently defined symbols and the binding scopes
type Environment interface {
	FindNamespace(name string) namespace.Namespace

	AddNamespace(ns namespace.Namespace) error

	CurrentNamespace() namespace.Namespace

	ChangeCurrentNamespace(name string) namespace.Namespace

	FindExportedSymbolInNamespace(name string, ns string) *symbols.Symbol

	FindInternedSymbolInNamespace(name string, ns string) *symbols.Symbol

	CurrentScope() scope.Scope

	PopScope() scope.Scope

	PushScope(scope scope.Scope) scope.Scope

	CaptureScope() scope.Scope

	AddGlobalBinding(sym *symbols.Symbol, obj types.Object)

	AddBinding(sym *symbols.Symbol, obj types.Object)

	GetBinding(sym *symbols.Symbol) types.Object

	SetBinding(sym *symbols.Symbol, obj types.Object) error

	FindSymbol(name string) *symbols.Symbol

	DefineSymbol(name string, reserved bool, value types.Object) *symbols.Symbol

	InternSymbol(name string) *symbols.Symbol

	InternKeyword(name string) *symbols.Symbol

	Gensym() *symbols.Symbol

	Eval(obj types.Object, context interface{}) (types.Object, error)

	Context() map[string]interface{}

	PushDepthContext(string)

	PopDepthContext(string)

	HasDepthContext(string) bool
}
