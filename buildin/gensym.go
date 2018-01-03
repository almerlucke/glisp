package buildin

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
)

// Gensym buildin function
func Gensym(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	return env.Gensym(), nil
}

// CreateBuildinGensym creates a buildin function object
func CreateBuildinGensym() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Gensym, 0, false)
}
