package environment

import (
	"container/list"
	"errors"
	"fmt"

	globals "github.com/almerlucke/glisp/globals/symbols"

	"github.com/almerlucke/glisp/buildin"
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/interfaces/function"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/symbols"
)

// Environment holds the currently defined symbols and the binding scopes
type Environment struct {
	// Symbol table holds all defined symbols in the environment
	symTable map[string]*symbols.Symbol

	// gensymCounter is used to create a unique symbol
	gensymCounter uint64

	// Scopes can be nested
	scopes *list.List

	// Global scope is always at the end of the list, but also keep
	// a reference here for ease
	globalScope environment.Scope

	// Context can hold all kinds of values
	context map[string]interface{}
}

// New returns a new default environment
func New() *Environment {
	globalScope := make(environment.Scope)
	scopes := list.New()
	scopes.PushFront(globalScope)

	env := &Environment{
		globalScope: globalScope,
		symTable: map[string]*symbols.Symbol{
			globals.NILSymbol.Name:              globals.NILSymbol,
			globals.TSymbol.Name:                globals.TSymbol,
			globals.DotSymbol.Name:              globals.DotSymbol,
			globals.CloseParenthesisSymbol.Name: globals.CloseParenthesisSymbol,
			globals.QuoteSymbol.Name:            globals.QuoteSymbol,
			globals.BackquoteSymbol.Name:        globals.BackquoteSymbol,
			globals.UnquoteSymbol.Name:          globals.UnquoteSymbol,
			globals.SpliceSymbol.Name:           globals.SpliceSymbol,
			globals.AndRestSymbol.Name:          globals.AndRestSymbol,
		},
		scopes:  scopes,
		context: map[string]interface{}{},
	}

	env.AddGlobalBinding(globals.QuoteSymbol, buildin.CreateBuildinQuote())
	env.AddGlobalBinding(globals.BackquoteSymbol, buildin.CreateBuildinBackquote())
	env.AddGlobalBinding(globals.UnquoteSymbol, buildin.CreateBuildinUnquote())
	env.AddGlobalBinding(globals.SpliceSymbol, buildin.CreateBuildinUnquote())

	env.AddGlobalBinding(env.DefineSymbol("LIST", true, nil), buildin.CreateBuildinList())
	env.AddGlobalBinding(env.DefineSymbol("CDR", true, nil), buildin.CreateBuildinCdr())
	env.AddGlobalBinding(env.DefineSymbol("CAR", true, nil), buildin.CreateBuildinCar())
	env.AddGlobalBinding(env.DefineSymbol("CONS", true, nil), buildin.CreateBuildinCons())
	env.AddGlobalBinding(env.DefineSymbol("LAMBDA", true, nil), buildin.CreateBuildinLambda())
	env.AddGlobalBinding(env.DefineSymbol("MACRO", true, nil), buildin.CreateBuildinMacro())
	env.AddGlobalBinding(env.DefineSymbol("GENSYM", true, nil), buildin.CreateBuildinGensym())
	env.AddGlobalBinding(env.DefineSymbol("PRINT", true, nil), buildin.CreateBuildinPrint())
	env.AddGlobalBinding(env.DefineSymbol("EXIT", true, nil), buildin.CreateBuildinExit())
	env.AddGlobalBinding(env.DefineSymbol("LOAD", true, nil), buildin.CreateBuildinLoad())
	env.AddGlobalBinding(env.DefineSymbol("VAR", true, nil), buildin.CreateBuildinVar())
	env.AddGlobalBinding(env.DefineSymbol("=", true, nil), buildin.CreateBuildinAssign())
	env.AddGlobalBinding(env.DefineSymbol("SCOPE", true, nil), buildin.CreateBuildinScope())
	env.AddGlobalBinding(env.DefineSymbol("EVAL", true, nil), buildin.CreateBuildinEval())
	env.AddGlobalBinding(env.DefineSymbol("ELT", true, nil), buildin.CreateBuildinElt())
	env.AddGlobalBinding(env.DefineSymbol("ARRAY", true, nil), buildin.CreateBuildinArray())
	env.AddGlobalBinding(env.DefineSymbol("MAKE-ARRAY", true, nil), buildin.CreateBuildinMakeArray())
	env.AddGlobalBinding(env.DefineSymbol("HASHTABLE", true, nil), buildin.CreateBuildinHashTable())

	return env
}

// Gensym creates a unique uninterned symbol
func (env *Environment) Gensym() *symbols.Symbol {
	name := fmt.Sprintf("SYSTEM::G%d", env.gensymCounter)
	env.gensymCounter++

	return &symbols.Symbol{
		Name: name,
	}
}

// CurrentScope returns the current scope
func (env *Environment) CurrentScope() environment.Scope {
	return env.scopes.Front().Value.(environment.Scope)
}

// PopScope pop a scope
func (env *Environment) PopScope() environment.Scope {
	return env.scopes.Remove(env.scopes.Front()).(environment.Scope)
}

// PushScope push a scope, if scope is nil create a new one
func (env *Environment) PushScope(s environment.Scope) environment.Scope {
	if s == nil {
		s = make(environment.Scope)
	}

	env.scopes.PushFront(s)

	return s
}

// CaptureScope captures the scope stack flattened except for the global scope
// if a symbol is shadowed, only capture the topmost binding
func (env *Environment) CaptureScope() environment.Scope {
	s := make(environment.Scope)

	for e := env.scopes.Front(); e.Next() != nil; e = e.Next() {
		otherScope := e.Value.(environment.Scope)
		for sym, val := range otherScope {
			if s[sym] == nil {
				s[sym] = val
			}
		}
	}

	return s
}

// AddGlobalBinding bind object to symbol in the global scope
func (env *Environment) AddGlobalBinding(sym *symbols.Symbol, obj types.Object) {
	env.globalScope[sym] = obj
}

// AddBinding bind object to symbol in the current scope
func (env *Environment) AddBinding(sym *symbols.Symbol, obj types.Object) {
	env.scopes.Front().Value.(environment.Scope)[sym] = obj
}

// GetBinding get binding for symbol
func (env *Environment) GetBinding(sym *symbols.Symbol) types.Object {
	var obj types.Object

	// Go through scopes to find binding
	for e := env.scopes.Front(); e != nil; e = e.Next() {
		s := e.Value.(environment.Scope)
		obj = s[sym]
		if obj != nil {
			break
		}
	}

	return obj
}

// SetBinding set binding for an already defined symbol
func (env *Environment) SetBinding(sym *symbols.Symbol, obj types.Object) error {
	// Go through scopes to find binding
	for e := env.scopes.Front(); e != nil; e = e.Next() {
		s := e.Value.(environment.Scope)
		_, ok := s[sym]
		if ok {
			s[sym] = obj
			return nil
		}
	}

	return fmt.Errorf("unbound symbol %v", sym)
}

// GetSymbol returns a symbol or nil if not found
func (env *Environment) GetSymbol(name string) *symbols.Symbol {
	return env.symTable[name]
}

// DefineSymbol adds a new symbol to the environment or returns an
// existing symbol
func (env *Environment) DefineSymbol(name string, reserved bool, value types.Object) *symbols.Symbol {
	sym := env.symTable[name]

	if sym == nil {
		sym = &symbols.Symbol{
			Name:     name,
			Reserved: reserved,
			Value:    value,
		}

		env.symTable[name] = sym
	}

	return sym
}

// Context returns the environment context
func (env *Environment) Context() map[string]interface{} {
	return env.context
}

// Eval evaluates an object with this environment
func (env *Environment) Eval(obj types.Object) (types.Object, error) {
	result := obj

	switch obj.Type() {

	case types.Symbol:
		result = env.GetBinding(obj.(*symbols.Symbol))
		if result == nil {
			return nil, fmt.Errorf("unbound symbol %v", obj)
		}

	case types.Cons:
		// List to evaluate
		c := obj.(*cons.Cons)

		// Evaluate first elem
		r, err := env.Eval(c.Car)
		if err != nil {
			return nil, err
		}

		// Must be a function
		if r.Type() != types.Function {
			return nil, fmt.Errorf("eval %v is not a function", r)
		}

		fun := r.(function.Function)

		// Check for pure and get length
		pure, length := c.Info()
		if !pure {
			return nil, errors.New("eval can't evaluate a dotted list")
		}

		// Check if we have enough arguments
		if (length - 1) < int64(fun.NumArgs()) {
			return nil, fmt.Errorf("not enough arguments to function %v", c.Car)
		}

		// If we need to first evaluate all args
		var args *cons.Cons
		if c.Cdr != types.NIL {
			args = c.Cdr.(*cons.Cons)
			if fun.EvalArgs() {
				seq, serr := args.Map(func(obj types.Object) (types.Object, error) {
					return env.Eval(obj)
				})

				if serr != nil {
					return nil, serr
				}

				args = seq.(*cons.Cons)
			}
		}

		// Evaluate function call
		result, err = fun.Eval(args, env)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
