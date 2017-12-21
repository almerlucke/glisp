package cons

import (
	"fmt"

	"github.com/almerlucke/glisp/types"
)

// MapFun function to be mapped
type MapFun func(obj types.Object) (types.Object, error)

// IterFun function to iterate over list
type IterFun func(obj types.Object, index uint64) error

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

// Last returns last linked cons
func (c *Cons) Last() *Cons {
	// Get last cons from c
	last := c

	for ; last.Cdr.Type() == types.Cons; last = last.Cdr.(*Cons) {
	}

	return last
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

	err := c.Iter(func(obj types.Object, index uint64) error {
		car, err := fun(obj)
		if err != nil {
			return err
		}

		builder.PushBackObject(car)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return builder.Head, nil
}

// Iter over a list
func (c *Cons) Iter(fun IterFun) error {
	index := uint64(0)

	for e := types.Object(c); e.Type() == types.Cons; e = e.(*Cons).Cdr {
		err := fun(e.(*Cons).Car, index)
		if err != nil {
			return err
		}

		index++
	}

	return nil
}

// Access access the nth element
func (c *Cons) Access(nth uint64) types.Object {
	index := uint64(0)

	for e := types.Object(c); e.Type() == types.Cons; e = e.(*Cons).Cdr {
		if index == nth {
			return e.(*Cons).Car
		}

		index++
	}

	return nil
}

// Assign a new value to the nth cons car
func (c *Cons) Assign(nth uint64, val types.Object) bool {
	index := uint64(0)

	for e := types.Object(c); e.Type() == types.Cons; e = e.(*Cons).Cdr {
		if index == nth {
			e.(*Cons).Car = val
			return true
		}

		index++
	}

	return false
}

// ListBuilder can be used to build lists of cons objects
type ListBuilder struct {
	Head *Cons
	Tail *Cons
}

// PushBackObject push back an object
func (builder *ListBuilder) PushBackObject(obj types.Object) {
	c := &Cons{
		Car: obj,
		Cdr: types.NIL,
	}

	if builder.Head == nil {
		builder.Head = c
	} else {
		builder.Tail.Cdr = c
	}

	builder.Tail = c
}

// PushBack a new cons on the list
func (builder *ListBuilder) PushBack(c *Cons) {
	if builder.Head == nil {
		builder.Head = c
	} else {
		builder.Tail.Cdr = c
	}

	builder.Tail = c
}

// Append a list of conses
func (builder *ListBuilder) Append(c *Cons) {
	if builder.Head == nil {
		builder.Head = c
	} else {
		builder.Tail.Cdr = c
	}

	builder.Tail = c.Last()
}
