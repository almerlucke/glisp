package glisp

import (
	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/types/functions/buildin"
)

// CreateDefaultEnvironment creates a default GLisp environment
func CreateDefaultEnvironment() *environment.Environment {
	env := environment.New()
	env.AddBinding(env.DefineSymbol("QUOTE", true, nil), buildin.CreateBuildinQuote())
	return env
}
