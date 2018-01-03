package cons

import "github.com/almerlucke/glisp/types"

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
