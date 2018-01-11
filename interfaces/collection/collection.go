package collection

import (
	"github.com/almerlucke/glisp/types"
)

// MapFun function to be mapped
type MapFun func(obj types.Object, index interface{}) (types.Object, error)

// IterFun function to iterate over list, the bool return value can be used
// to indicate a stop, when it is true the iteration should stop
type IterFun func(obj types.Object, index interface{}) (bool, error)

// Collection interface
type Collection interface {
	types.Object
	Access(index interface{}) (types.Object, error)
	Assign(index interface{}, val types.Object) error
	Map(fun MapFun) (Collection, error)
	Iter(fun IterFun) error
	Length() uint64

	// Reduce types.Object
}
