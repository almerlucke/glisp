package environment

import (
	"container/list"

	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/symbol"
)

// Scope holds the bindings of a symbol to an object
type Scope map[*symbol.Symbol]types.Object

// Environment holds the currently defined symbols and the binding scopes
type Environment struct {
	// Symbol table holds all defined symbols in the environment
	symTable map[string]*symbol.Symbol

	// Scopes can be nested
	scopes *list.List
}

// New returns a new environment
func New() *Environment {
	globalScope := make(Scope)
	scopes := list.New()

	scopes.PushFront(globalScope)

	return &Environment{
		symTable: make(map[string]*symbol.Symbol),
		scopes:   scopes,
	}
}

// CurrentScope returns the current scope
func (env *Environment) CurrentScope() Scope {
	return env.scopes.Front().Value.(Scope)
}

// PopScope pop a scope
func (env *Environment) PopScope() Scope {
	return env.scopes.Remove(env.scopes.Front()).(Scope)
}

// PushScope push a new scope
func (env *Environment) PushScope() Scope {
	scope := make(Scope)
	env.scopes.PushFront(scope)

	return scope
}

// AddBinding bind object to symbol in the current scope
func (env *Environment) AddBinding(sym *symbol.Symbol, obj types.Object) {
	env.scopes.Front().Value.(Scope)[sym] = obj
}

// GetBinding get binding for symbol in current scope
func (env *Environment) GetBinding(sym *symbol.Symbol) types.Object {
	return env.scopes.Front().Value.(Scope)[sym]
}

// GetSymbol returns a symbol or nil if not found
func (env *Environment) GetSymbol(name string) *symbol.Symbol {
	return env.symTable[name]
}

// DefineSymbol adds a new symbol to the environment or returns an
// existing symbol
func (env *Environment) DefineSymbol(name string, reserved bool) *symbol.Symbol {
	sym := env.symTable[name]

	if sym == nil {
		sym = &symbol.Symbol{
			Name:     name,
			Reserved: reserved,
		}

		env.symTable[name] = sym
	}

	return sym
}
