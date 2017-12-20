package environment

import (
	"container/list"

	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/symbols"
)

// Scope holds the bindings of a symbol to an object
type Scope map[*symbols.Symbol]types.Object

// Environment holds the currently defined symbols and the binding scopes
type Environment struct {
	// Symbol table holds all defined symbols in the environment
	symTable map[string]*symbols.Symbol

	// Scopes can be nested
	scopes *list.List

	// Global scope is always at the end of the list, but also keep
	// a reference here for ease
	globalScope Scope
}

// New returns a new environment
func New() *Environment {
	globalScope := make(Scope)
	scopes := list.New()
	scopes.PushFront(globalScope)

	env := &Environment{
		globalScope: globalScope,
		symTable: map[string]*symbols.Symbol{
			NILSymbol.Name:              NILSymbol,
			TSymbol.Name:                TSymbol,
			DotSymbol.Name:              DotSymbol,
			CloseParenthesisSymbol.Name: CloseParenthesisSymbol,
			QuoteSymbol.Name:            QuoteSymbol,
			BackquoteSymbol.Name:        BackquoteSymbol,
			UnquoteSymbol.Name:          UnquoteSymbol,
			SpliceSymbol.Name:           SpliceSymbol,
			AndRestSymbol.Name:          AndRestSymbol,
		},
		scopes: scopes,
	}

	return env
}

// CurrentScope returns the current scope
func (env *Environment) CurrentScope() Scope {
	return env.scopes.Front().Value.(Scope)
}

// PopScope pop a scope
func (env *Environment) PopScope() Scope {
	return env.scopes.Remove(env.scopes.Front()).(Scope)
}

// PushScope push a scope, if scope is nil create a new one
func (env *Environment) PushScope(scope Scope) Scope {
	if scope == nil {
		scope = make(Scope)
	}

	env.scopes.PushFront(scope)

	return scope
}

// CaptureScope captures the scope stack flattened except for the global scope
// if a symbol is shadowed, only capture the topmost binding
func (env *Environment) CaptureScope() Scope {
	scope := make(Scope)

	for e := env.scopes.Front(); e.Next() != nil; e = e.Next() {
		otherScope := e.Value.(Scope)
		for sym, val := range otherScope {
			if scope[sym] == nil {
				scope[sym] = val
			}
		}
	}

	return scope
}

// AddGlobalBinding bind object to symbol in the global scope
func (env *Environment) AddGlobalBinding(sym *symbols.Symbol, obj types.Object) {
	env.globalScope[sym] = obj
}

// AddBinding bind object to symbol in the current scope
func (env *Environment) AddBinding(sym *symbols.Symbol, obj types.Object) {
	env.scopes.Front().Value.(Scope)[sym] = obj
}

// GetBinding get binding for symbol
func (env *Environment) GetBinding(sym *symbols.Symbol) types.Object {
	var obj types.Object

	// Go through scopes to find binding
	for e := env.scopes.Front(); e != nil; e = e.Next() {
		scope := e.Value.(Scope)
		obj = scope[sym]
		if obj != nil {
			break
		}
	}

	return obj
}

// GetSymbol returns a symbol or nil if not found
func (env *Environment) GetSymbol(name string) *symbols.Symbol {
	return env.symTable[name]
}

// DefineSymbol adds a new symbol to the environment or returns an
// existing symbol
func (env *Environment) DefineSymbol(name string, reserved bool, value types.Object) *symbols.Symbol {
	sym := env.symTable[name]

	if sym == nil {
		sym = &symbols.Symbol{
			Name:     name,
			Reserved: reserved,
			Value:    value,
		}

		env.symTable[name] = sym
	}

	return sym
}
