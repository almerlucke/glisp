package function

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
)

// ReturnContext can be passed as context to eval, the RETURN buildin function
// can put an object in this context before panicking
type ReturnContext struct {
	Object types.Object
}

// Function interface
type Function interface {
	types.Object
	NumArgs() int
	EvalArgs() bool
	Eval(*cons.Cons, environment.Environment, interface{}) (types.Object, error)
}

// Assignable allows for a function to be used with = assign, the value to Assign
// is passed as second arg to Assign, Assign should further be equal to Eval
type Assignable interface {
	Function
	Assign(*cons.Cons, types.Object, environment.Environment, interface{}) (types.Object, error)
}
