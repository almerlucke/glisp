package types

// Type of lisp object
type Type uint

const (
	// Null object type
	Null Type = iota
	// Boolean object type
	Boolean
	// Symbol object type
	Symbol
	// Cons object type
	Cons
	// String object type
	String
	// Number object type
	Number
	// Function object type
	Function
)

// Object interface, every Lisp object must implement these methods
type Object interface {
	Type() Type
}

// Nil is an empty struct, there will only be one Nil struct created, the
// singleton NIL
type Nil struct{}

// Bool is an empty struct, there will only be one Bool struct created, the
// singleton T
type Bool struct{}

// NIL is the global NIL object, used as false value and to represent the empty
// list
var NIL = &Nil{}

// T is the global true object, used to explicitly indicate a true value
var T = &Bool{}

// Type Boolean
func (b *Bool) Type() Type {
	return Boolean
}

// String implements the stringer interface
func (b *Bool) String() string {
	return "T"
}

// Type Null
func (n *Nil) Type() Type {
	return Null
}

// String implements the stringer interface
func (n *Nil) String() string {
	return "NIL"
}
