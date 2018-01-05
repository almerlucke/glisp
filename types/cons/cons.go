package cons

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/almerlucke/glisp/interfaces/collection"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/numbers"
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
func (c *Cons) Map(fun collection.MapFun) (collection.Collection, error) {
	builder := ListBuilder{}

	err := c.Iter(func(obj types.Object, index interface{}) error {
		car, err := fun(obj, index)
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
func (c *Cons) Iter(fun collection.IterFun) error {
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
func (c *Cons) Access(index interface{}) (types.Object, error) {
	num, ok := index.(*numbers.Number)
	if !ok {
		return nil, errors.New("list expects a number index")
	}

	nth := num.Int64Value()
	if nth < 0 {
		return nil, errors.New("index out of bounds")
	}

	unth := uint64(nth)
	cnt := uint64(0)

	for e := types.Object(c); e.Type() == types.Cons; e = e.(*Cons).Cdr {
		if cnt == unth {
			return e.(*Cons).Car, nil
		}

		cnt++
	}

	return nil, errors.New("index out of bounds")
}

// Assign a new value to the nth cons car
func (c *Cons) Assign(index interface{}, val types.Object) error {
	num, ok := index.(*numbers.Number)
	if !ok {
		return errors.New("list expects a number index")
	}

	nth := num.Int64Value()
	if nth < 0 {
		return errors.New("index out of bounds")
	}

	unth := uint64(nth)
	cnt := uint64(0)

	for e := types.Object(c); e.Type() == types.Cons; e = e.(*Cons).Cdr {
		if cnt == unth {
			e.(*Cons).Car = val
			return nil
		}

		cnt++
	}

	return errors.New("index out of bounds")
}
