package namespace

import (
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/symbols"
)

// Namespace interface
type Namespace interface {
	types.Object
	Name() string
	FindSymbol(name string) *symbols.Symbol
	Intern(name string) *symbols.Symbol
	Use(ns Namespace)
	Shadow(name string)
	Import(name string, ns Namespace) bool
	CanIntern() bool
	DefineSymbol(name string, reserved bool, value types.Object, export bool) *symbols.Symbol
	ExportedSymbols() map[string]*symbols.Symbol
	InternedSymbols() map[string]*symbols.Symbol
	Export(sym *symbols.Symbol)
	Add(sym *symbols.Symbol, export bool)
}
