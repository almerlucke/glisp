package environment

import (
	"container/list"
	"fmt"

	"github.com/almerlucke/glisp/scope"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/symbols"
)

// Environment holds the currently defined symbols and the binding scopes
type Environment struct {
	// Symbol table holds all defined symbols in the environment
	symTable map[string]*symbols.Symbol

	// Scopes can be nested
	scopes *list.List

	// Global scope is always at the end of the list, but also keep
	// a reference here for ease
	globalScope scope.Scope

	// Context can hold all kinds of values
	context map[string]interface{}
}

// New returns a new environment
func New() *Environment {
	globalScope := make(scope.Scope)
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
		scopes:  scopes,
		context: map[string]interface{}{},
	}

	return env
}

// CurrentScope returns the current scope
func (env *Environment) CurrentScope() scope.Scope {
	return env.scopes.Front().Value.(scope.Scope)
}

// PopScope pop a scope
func (env *Environment) PopScope() scope.Scope {
	return env.scopes.Remove(env.scopes.Front()).(scope.Scope)
}

// PushScope push a scope, if scope is nil create a new one
func (env *Environment) PushScope(s scope.Scope) scope.Scope {
	if s == nil {
		s = make(scope.Scope)
	}

	env.scopes.PushFront(s)

	return s
}

// CaptureScope captures the scope stack flattened except for the global scope
// if a symbol is shadowed, only capture the topmost binding
func (env *Environment) CaptureScope() scope.Scope {
	s := make(scope.Scope)

	for e := env.scopes.Front(); e.Next() != nil; e = e.Next() {
		otherScope := e.Value.(scope.Scope)
		for sym, val := range otherScope {
			if s[sym] == nil {
				s[sym] = val
			}
		}
	}

	return s
}

// AddGlobalBinding bind object to symbol in the global scope
func (env *Environment) AddGlobalBinding(sym *symbols.Symbol, obj types.Object) {
	env.globalScope[sym] = obj
}

// AddBinding bind object to symbol in the current scope
func (env *Environment) AddBinding(sym *symbols.Symbol, obj types.Object) {
	env.scopes.Front().Value.(scope.Scope)[sym] = obj
}

// GetBinding get binding for symbol
func (env *Environment) GetBinding(sym *symbols.Symbol) types.Object {
	var obj types.Object

	// Go through scopes to find binding
	for e := env.scopes.Front(); e != nil; e = e.Next() {
		s := e.Value.(scope.Scope)
		obj = s[sym]
		if obj != nil {
			break
		}
	}

	return obj
}

// SetBinding set binding for an already defined symbol
func (env *Environment) SetBinding(sym *symbols.Symbol, obj types.Object) error {
	// Go through scopes to find binding
	for e := env.scopes.Front(); e != nil; e = e.Next() {
		s := e.Value.(scope.Scope)
		_, ok := s[sym]
		if ok {
			s[sym] = obj
			return nil
		}
	}

	return fmt.Errorf("unbound symbol %v", sym)
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

func (env *Environment) Context() map[string]interface{} {
	return env.context
}
