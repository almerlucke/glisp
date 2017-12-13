package symbol

import "github.com/almerlucke/glisp/types"

// Symbol struct
type Symbol struct {
	Name     string
	Reserved bool
}

// Type Symbol
func (sym *Symbol) Type() types.Type {
	return types.Symbol
}
