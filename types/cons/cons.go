package cons

import (
	"fmt"

	"github.com/almerlucke/glisp/types"
)

// MapFun function to be mapped
type MapFun func(obj types.Object) (types.Object, error)

// Cons is the main list structure
type Cons struct {
	Car types.Object
	Cdr types.Object
}

// Type Cons for Object interface
func (c *Cons) Type() types.Type {
	return types.Cons
}

// String for stringer interface
func (c *Cons) String() string {
	str := ""

	var e types.Object = c

	for ; e.Type() == types.Cons; e = e.(*Cons).Cdr {
		if str == "" {
			str = "("
		} else {
			str += " "
		}

		str += fmt.Sprintf("%v", e.(*Cons).Car)
	}

	if e != types.NIL {
		str += fmt.Sprintf(" . %v)", e)
	} else {
		str += ")"
	}

	return str
}

// IsPureList performs a check if the list is a pure list, so last atom
// is NIL
func (c *Cons) IsPureList() bool {
	var e types.Object = c

	for ; e.Type() == types.Cons; e = e.(*Cons).Cdr {
	}

	return e == types.NIL
}

// Length of the list
func (c *Cons) Length() int64 {
	length := int64(0)

	var e types.Object = c

	for ; e.Type() == types.Cons; e = e.(*Cons).Cdr {
		length++
	}

	return length
}

// Info about the list, if it is pure and it's length
func (c *Cons) Info() (bool, int64) {
	length := int64(0)

	var e types.Object = c

	for ; e.Type() == types.Cons; e = e.(*Cons).Cdr {
		length++
	}

	return e == types.NIL, length
}

// Map maps a function over a cons and returns a new cons
func (c *Cons) Map(fun MapFun) (*Cons, error) {
	builder := ListBuilder{}
	var e types.Object = c

	for ; e.Type() == types.Cons; e = e.(*Cons).Cdr {
		car, err := fun(e.(*Cons).Car)
		if err != nil {
			return nil, err
		}

		builder.PushBack(&Cons{
			Car: car,
			Cdr: types.NIL,
		})
	}

	return builder.Head, nil
}

// ListBuilder can be used to build lists of cons objects
type ListBuilder struct {
	Head *Cons
	Tail *Cons
}

// PushBack a new cons on the list
func (builder *ListBuilder) PushBack(c *Cons) {
	if builder.Head == nil {
		builder.Head = c
		builder.Tail = c
	} else {
		builder.Tail.Cdr = c
		builder.Tail = c
	}
}
