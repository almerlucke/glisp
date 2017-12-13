package cons

import "github.com/almerlucke/glisp/types"

// Cons is the main list structure
type Cons struct {
	Car types.Object
	Cdr types.Object
}

// Type Cons for Object interface
func (cons *Cons) Type() types.Type {
	return types.Cons
}
