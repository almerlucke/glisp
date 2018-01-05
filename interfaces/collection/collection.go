package collection

import (
	"github.com/almerlucke/glisp/types"
)

// MapFun function to be mapped
type MapFun func(obj types.Object, index interface{}) (types.Object, error)

// IterFun function to iterate over list
type IterFun func(obj types.Object, index interface{}) error

// Collection interface
type Collection interface {
	Access(index interface{}) (types.Object, error)
	Assign(index interface{}, val types.Object) error
	Map(fun MapFun) (Collection, error)
	Iter(fun IterFun) error
	Length() uint64
}
