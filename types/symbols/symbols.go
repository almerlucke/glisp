package symbols

import (
	"github.com/almerlucke/glisp/types"
)

// Symbol struct
type Symbol struct {
	Name     string
	Reserved bool
	// Value can be the self referencing value of the symbol,
	// when the reader encounters a symbol with a value, it will
	// return the value instead of the symbol
	Value     types.Object
	Interned  bool
	IsKeyword bool
}

// Type Symbol
func (sym *Symbol) Type() types.Type {
	return types.Symbol
}

// String for stringer interface
func (sym *Symbol) String() string {
	if sym.Interned {
		return sym.Name
	}

	return "#:" + sym.Name
}

// Compare for comparable interface
func (sym *Symbol) Compare(obj types.Comparable) bool {
	return sym == obj
}
