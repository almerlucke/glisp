package sequence

import (
	"github.com/almerlucke/glisp/types"
)

// MapFun function to be mapped
type MapFun func(obj types.Object) (types.Object, error)

// IterFun function to iterate over list
type IterFun func(obj types.Object, index uint64) error

// Sequence interface
type Sequence interface {
	Access(nth uint64) types.Object
	Assign(nth uint64, val types.Object) bool
	Map(fun MapFun) (Sequence, error)
	Iter(fun IterFun) error
	Length() uint64
}
