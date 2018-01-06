package scope

import (
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/symbols"
)

// Scope holds the bindings of a symbol to an object
type Scope map[*symbols.Symbol]types.Object
