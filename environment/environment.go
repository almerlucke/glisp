package environment

import (
	"container/list"

	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/symbols"
)

// DotSymbol is used for dotted lists in the reader
var DotSymbol = &symbols.Symbol{
	Name:     ".",
	Reserved: true,
}

// CloseParenthesisSymbol is used to signal a closing parenthesis
// in the OpenParenthesisMacro
var CloseParenthesisSymbol = &symbols.Symbol{
	Name:     ")",
	Reserved: true,
}

// NILSymbol always references NIL instead of the symbol
var NILSymbol = &symbols.Symbol{
	Name:     "NIL",
	Reserved: true,
	Value:    types.NIL,
}

// TSymbol always references T instead of the symbol
var TSymbol = &symbols.Symbol{
	Name:     "T",
	Reserved: true,
	Value:    types.T,
}

// QuoteSymbol is used for quoted objects
var QuoteSymbol = &symbols.Symbol{
	Name:     "QUOTE",
	Reserved: true,
}

// Scope holds the bindings of a symbol to an object
type Scope map[*symbols.Symbol]types.Object

// Environment holds the currently defined symbols and the binding scopes
type Environment struct {
	// Symbol table holds all defined symbols in the environment
	symTable map[string]*symbols.Symbol

	// Scopes can be nested
	scopes *list.List
}

// New returns a new environment
func New() *Environment {
	globalScope := make(Scope)
	scopes := list.New()

	scopes.PushFront(globalScope)

	env := &Environment{
		symTable: make(map[string]*symbols.Symbol),
		scopes:   scopes,
	}

	env.symTable["NIL"] = NILSymbol
	env.symTable["T"] = TSymbol
	env.symTable["."] = DotSymbol
	env.symTable[")"] = CloseParenthesisSymbol
	env.symTable["QUOTE"] = QuoteSymbol

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

// PushScope push a new scope
func (env *Environment) PushScope() Scope {
	scope := make(Scope)
	env.scopes.PushFront(scope)

	return scope
}

// AddBinding bind object to symbol in the current scope
func (env *Environment) AddBinding(sym *symbols.Symbol, obj types.Object) {
	env.scopes.Front().Value.(Scope)[sym] = obj
}

// GetBinding get binding for symbol in current scope
func (env *Environment) GetBinding(sym *symbols.Symbol) types.Object {
	return env.scopes.Front().Value.(Scope)[sym]
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
