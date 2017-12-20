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
