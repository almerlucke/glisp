package function

import (
	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
)

// Function interface
type Function interface {
	types.Object
	NumArgs() int
	EvalArgs() bool
	Eval(*cons.Cons, *environment.Environment) (types.Object, error)
}

// Assignable allows for a function to be used with = assign, the value to Assign
// is passed as second arg to Assign, Assign should further be equal to Eval
type Assignable interface {
	Function
	Assign(*cons.Cons, types.Object, *environment.Environment) (types.Object, error)
}
