package namespaces

import (
	"fmt"

	"github.com/almerlucke/glisp/interfaces/namespace"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/symbols"
)

// Namespace implementation
type Namespace struct {
	name            string
	internedSymbols map[string]*symbols.Symbol
	exportedSymbols map[string]*symbols.Symbol
	canIntern       bool
}

// NewNamespace creates a new namespace
func NewNamespace(name string, canIntern bool) *Namespace {
	return &Namespace{
		name:            name,
		canIntern:       canIntern,
		internedSymbols: make(map[string]*symbols.Symbol),
		exportedSymbols: make(map[string]*symbols.Symbol),
	}
}

// Use copies symbols from another namespace
func (n *Namespace) Use(ns namespace.Namespace) {
	// Copy namespace symbols to this namespace
	exportedSymbols := ns.ExportedSymbols()
	for k, v := range exportedSymbols {
		n.internedSymbols[k] = v
	}
}

// Shadow shadows a used symbol
func (n *Namespace) Shadow(name string) {
	// Actually just defines a new symbol
	n.DefineSymbol(name, false, nil, false)
}

// Import a single symbol from another namespace
func (n *Namespace) Import(name string, ns namespace.Namespace) bool {
	exportedSymbols := ns.ExportedSymbols()

	sym := exportedSymbols[name]
	if sym != nil {
		n.internedSymbols[name] = sym
		return true
	}

	return false
}

// Type Namespace for Object interface
func (n *Namespace) Type() types.Type {
	return types.Namespace
}

// String for stringer interface
func (n *Namespace) String() string {
	return fmt.Sprintf("#<The %v namespace>", n.name)
}

// Eql obj
func (n *Namespace) Eql(obj types.Object) bool {
	return n == obj
}

// Equal obj
func (n *Namespace) Equal(obj types.Object) bool {
	return n == obj
}

// Name of the namespace
func (n *Namespace) Name() string {
	return n.name
}

// FindSymbol finds an interned symbol
func (n *Namespace) FindSymbol(name string) *symbols.Symbol {
	return n.internedSymbols[name]
}

// Intern returns an existing symbol with name or creates a new one
func (n *Namespace) Intern(name string) *symbols.Symbol {
	sym := n.internedSymbols[name]
	if sym != nil {
		return sym
	}

	sym = &symbols.Symbol{
		Name:     name,
		Interned: true,
	}

	n.internedSymbols[name] = sym

	return sym
}

// CanIntern check if namespace is locked
func (n *Namespace) CanIntern() bool {
	return n.canIntern
}

// DefineSymbol define a symbol
func (n *Namespace) DefineSymbol(name string, reserved bool, value types.Object, export bool) *symbols.Symbol {
	sym := &symbols.Symbol{
		Name:     name,
		Reserved: reserved,
		Interned: true,
		Value:    value,
	}

	n.internedSymbols[name] = sym

	if export {
		n.exportedSymbols[name] = sym
	}

	return sym
}

// ExportedSymbols gets all exported symbols
func (n *Namespace) ExportedSymbols() map[string]*symbols.Symbol {
	return n.exportedSymbols
}

// InternedSymbols gets all interned symbols
func (n *Namespace) InternedSymbols() map[string]*symbols.Symbol {
	return n.internedSymbols
}

// Export exports a symbol
func (n *Namespace) Export(sym *symbols.Symbol) {
	n.exportedSymbols[sym.Name] = sym
}

// Add an existing symbol
func (n *Namespace) Add(sym *symbols.Symbol, export bool) {
	n.internedSymbols[sym.Name] = sym

	if export {
		n.exportedSymbols[sym.Name] = sym
	}
}
