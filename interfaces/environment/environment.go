package environment

import (
	"github.com/almerlucke/glisp/scope"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/symbols"
)

// Environment implements the currently defined symbols and the binding scopes
type Environment interface {
	CurrentScope() scope.Scope

	PopScope() scope.Scope

	PushScope(scope scope.Scope) scope.Scope

	CaptureScope() scope.Scope

	AddGlobalBinding(sym *symbols.Symbol, obj types.Object)

	AddBinding(sym *symbols.Symbol, obj types.Object)

	GetBinding(sym *symbols.Symbol) types.Object

	SetBinding(sym *symbols.Symbol, obj types.Object) error

	GetSymbol(name string) *symbols.Symbol

	DefineSymbol(name string, reserved bool, value types.Object) *symbols.Symbol

	Context() map[string]interface{}
}
