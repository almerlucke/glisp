package environment

import (
	"container/list"
	"errors"
	"fmt"

	namespacesSetup "github.com/almerlucke/glisp/environment/namespaces"

	"github.com/almerlucke/glisp/interfaces/function"
	"github.com/almerlucke/glisp/interfaces/namespace"
	"github.com/almerlucke/glisp/scope"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/namespaces"
	"github.com/almerlucke/glisp/types/symbols"
)

// Environment holds the currently defined symbols and the binding scopes
type Environment struct {
	// Symbol table holds all defined symbols in the environment
	symTable map[string]*symbols.Symbol

	// gensymCounter is used to create a unique uninterned symbol
	gensymCounter uint64

	// Scopes can be nested
	scopes *list.List

	// Global scope is always at the end of the list, but also keep
	// a reference here for ease
	globalScope scope.Scope

	// Context can hold all kinds of values
	context map[string]interface{}

	// Currently used namespace
	currentNamespace namespace.Namespace

	// Keyword namespace
	keywordNamespace namespace.Namespace

	// all namespaces
	namespaces map[string]namespace.Namespace
}

// New returns a new default environment
func New() *Environment {
	globalScope := make(scope.Scope)
	scopes := list.New()
	scopes.PushFront(globalScope)

	env := &Environment{
		globalScope: globalScope,
		namespaces:  map[string]namespace.Namespace{},
		scopes:      scopes,
		context:     map[string]interface{}{},
	}

	glispNS := namespacesSetup.CreateGlispNamespace(env)
	keywordNS := namespaces.NewNamespace("KEYWORD", false)
	glispUserNS := namespaces.NewNamespace("GLISP-USER", true)
	mathNS := namespacesSetup.CreateMathNamespace(env)

	env.namespaces[glispNS.Name()] = glispNS
	env.namespaces[keywordNS.Name()] = keywordNS
	env.namespaces[glispUserNS.Name()] = glispUserNS
	env.namespaces[mathNS.Name()] = mathNS

	// Set keyword namespace
	env.keywordNamespace = keywordNS

	// Let glisp-user namespace use the glisp namespace
	glispUserNS.Use(glispNS)

	// Set current namespace to glisp user
	env.currentNamespace = glispUserNS

	return env
}

// FindNamespace find a namespace
func (env *Environment) FindNamespace(name string) namespace.Namespace {
	return env.namespaces[name]
}

// AddNamespace add a namespace
func (env *Environment) AddNamespace(ns namespace.Namespace) error {
	ens := env.namespaces[ns.Name()]
	if ens != nil {
		return fmt.Errorf("namespace %v already exists", ns.Name())
	}

	env.namespaces[ns.Name()] = ns

	return nil
}

// ChangeCurrentNamespace changes the current namespace
func (env *Environment) ChangeCurrentNamespace(name string) namespace.Namespace {
	ns := env.namespaces[name]
	if ns != nil {
		env.currentNamespace = ns
		return ns
	}

	return nil
}

// CurrentNamespace gets the current namespace
func (env *Environment) CurrentNamespace() namespace.Namespace {
	return env.currentNamespace
}

// Gensym creates a unique uninterned symbol
func (env *Environment) Gensym() *symbols.Symbol {
	name := fmt.Sprintf("G%d", env.gensymCounter)
	env.gensymCounter++

	return &symbols.Symbol{
		Name: name,
	}
}

// CurrentScope returns the current scope
func (env *Environment) CurrentScope() scope.Scope {
	return env.scopes.Front().Value.(scope.Scope)
}

// PopScope pop a scope
func (env *Environment) PopScope() scope.Scope {
	return env.scopes.Remove(env.scopes.Front()).(scope.Scope)
}

// PushScope push a scope, if scope is nil create a new one
func (env *Environment) PushScope(s scope.Scope) scope.Scope {
	if s == nil {
		s = make(scope.Scope)
	}

	env.scopes.PushFront(s)

	return s
}

// CaptureScope captures the scope stack flattened except for the global scope
// if a symbol is shadowed, only capture the topmost binding
func (env *Environment) CaptureScope() scope.Scope {
	s := make(scope.Scope)

	for e := env.scopes.Front(); e.Next() != nil; e = e.Next() {
		otherScope := e.Value.(scope.Scope)
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
	env.scopes.Front().Value.(scope.Scope)[sym] = obj
}

// GetBinding get binding for symbol
func (env *Environment) GetBinding(sym *symbols.Symbol) types.Object {
	var obj types.Object

	// Go through scopes to find binding
	for e := env.scopes.Front(); e != nil; e = e.Next() {
		s := e.Value.(scope.Scope)
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
		s := e.Value.(scope.Scope)
		_, ok := s[sym]
		if ok {
			s[sym] = obj
			return nil
		}
	}

	return fmt.Errorf("unbound symbol %v", sym)
}

// FindSymbol returns a symbol or nil if not found
func (env *Environment) FindSymbol(name string) *symbols.Symbol {
	return env.currentNamespace.FindSymbol(name)
}

// FindExportedSymbolInNamespace find exported symbol
func (env *Environment) FindExportedSymbolInNamespace(name string, ns string) *symbols.Symbol {
	return env.namespaces[ns].FindSymbol(name)
}

// FindInternedSymbolInNamespace find interned symbol
func (env *Environment) FindInternedSymbolInNamespace(name string, ns string) *symbols.Symbol {
	return env.namespaces[ns].FindSymbol(name)
}

// DefineSymbol adds a new symbol to the environment or returns an
// existing symbol
func (env *Environment) DefineSymbol(name string, reserved bool, value types.Object) *symbols.Symbol {
	return env.currentNamespace.DefineSymbol(name, reserved, value, false)
}

// InternSymbol interns a symbol
func (env *Environment) InternSymbol(name string) *symbols.Symbol {
	return env.currentNamespace.Intern(name)
}

// InternKeyword adds a keyword
func (env *Environment) InternKeyword(name string) *symbols.Symbol {
	sym := env.keywordNamespace.Intern(name)

	// Is keyword
	sym.IsKeyword = true

	// Export the keyword
	env.keywordNamespace.Export(sym)

	// Add a global reserved binding
	sym.Reserved = true
	env.AddGlobalBinding(sym, sym)

	return sym
}

// Context returns the environment context
func (env *Environment) Context() map[string]interface{} {
	return env.context
}

// PushDepthContext push specific depth context type
func (env *Environment) PushDepthContext(d string) {
	ctx, ok := env.context[d]
	if ok {
		env.context[d] = ctx.(uint64) + 1
	} else {
		env.context[d] = uint64(1)
	}
}

// PopDepthContext pop specific depth context type
func (env *Environment) PopDepthContext(d string) {
	ctx, ok := env.context[d]
	if ok {
		env.context[d] = ctx.(uint64) - 1
	}
}

// HasDepthContext check if specific depth context type is set
func (env *Environment) HasDepthContext(d string) bool {
	ctx, ok := env.context[d]
	return ok && ctx.(uint64) > 0
}

// Eval evaluates an object with this environment
func (env *Environment) Eval(obj types.Object, context interface{}) (types.Object, error) {
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
		r, err := env.Eval(c.Car, context)
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
				col, serr := args.Map(func(obj types.Object, index interface{}) (types.Object, error) {
					return env.Eval(obj, context)
				})

				if serr != nil {
					return nil, serr
				}

				args = col.(*cons.Cons)
			}
		}

		// Evaluate function call
		result, err = fun.Eval(args, env, context)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
