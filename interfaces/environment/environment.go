package environment

import (
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/symbols"
)

// Scope holds the bindings of a symbol to an object
type Scope map[*symbols.Symbol]types.Object

// Environment implements the currently defined symbols and the binding scopes
type Environment interface {
	CurrentScope() Scope

	PopScope() Scope

	PushScope(scope Scope) Scope

	CaptureScope() Scope

	AddGlobalBinding(sym *symbols.Symbol, obj types.Object)

	AddBinding(sym *symbols.Symbol, obj types.Object)

	GetBinding(sym *symbols.Symbol) types.Object

	SetBinding(sym *symbols.Symbol, obj types.Object) error

	GetSymbol(name string) *symbols.Symbol

	DefineSymbol(name string, reserved bool, value types.Object) *symbols.Symbol

	Gensym() *symbols.Symbol

	Eval(obj types.Object, context interface{}) (types.Object, error)

	Context() map[string]interface{}

	PushDepthContext(string)

	PopDepthContext(string)

	HasDepthContext(string) bool
}
