package cons

import (
	"bytes"
	"fmt"

	"github.com/almerlucke/glisp/interfaces/sequence"
	"github.com/almerlucke/glisp/types"
)

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
	var buffer bytes.Buffer

	first := true
	var e types.Object = c

	for ; e.Type() == types.Cons; e = e.(*Cons).Cdr {
		if first {
			first = false
			buffer.WriteString("(")
		} else {
			buffer.WriteString(" ")
		}

		buffer.WriteString(e.(*Cons).Car.String())
	}

	if e != types.NIL {
		buffer.WriteString(fmt.Sprintf(" . %v)", e))
	} else {
		buffer.WriteString(")")
	}

	return buffer.String()
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
func (c *Cons) Length() uint64 {
	length := uint64(0)

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
func (c *Cons) Map(fun sequence.MapFun) (sequence.Sequence, error) {
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
func (c *Cons) Iter(fun sequence.IterFun) error {
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
