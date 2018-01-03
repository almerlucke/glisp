package functions

import (
	"fmt"

	globals "github.com/almerlucke/glisp/globals/symbols"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/symbols"
)

// AssignableFunctionImp assign function implementation
type AssignableFunctionImp func(*cons.Cons, types.Object, environment.Environment) (types.Object, error)

// BuildinFunctionImp build in function implementation
type BuildinFunctionImp func(*cons.Cons, environment.Environment) (types.Object, error)

// BuildinFunction object
type BuildinFunction struct {
	imp      BuildinFunctionImp
	numArgs  int
	evalArgs bool
}

// NewBuildinFunction creates a new build in function
func NewBuildinFunction(imp BuildinFunctionImp, numArgs int, evalArgs bool) *BuildinFunction {
	return &BuildinFunction{
		imp:      imp,
		numArgs:  numArgs,
		evalArgs: evalArgs,
	}
}

// NumArgs number of arguments
func (fun *BuildinFunction) NumArgs() int {
	return fun.numArgs
}

// EvalArgs evaluate arguments before calling eval
func (fun *BuildinFunction) EvalArgs() bool {
	return fun.evalArgs
}

// Type of Function
func (fun *BuildinFunction) Type() types.Type {
	return types.Function
}

// Eval evaluates a function
func (fun *BuildinFunction) Eval(args *cons.Cons, env environment.Environment) (types.Object, error) {
	return fun.imp(args, env)
}

// String for Stringer
func (fun *BuildinFunction) String() string {
	return fmt.Sprintf("function(%p)", fun)
}

// AssignableFunction function that implements Assignable
type AssignableFunction struct {
	*BuildinFunction
	assignImp AssignableFunctionImp
}

// NewAssignableFunction creates a new assignable function
func NewAssignableFunction(imp BuildinFunctionImp, assignImp AssignableFunctionImp, numArgs int, evalArgs bool) *AssignableFunction {
	return &AssignableFunction{
		BuildinFunction: NewBuildinFunction(imp, numArgs, evalArgs),
		assignImp:       assignImp,
	}
}

// Assign call
func (fun *AssignableFunction) Assign(args *cons.Cons, val types.Object, env environment.Environment) (types.Object, error) {
	return fun.assignImp(args, val, env)
}

// LambdaFunction anonymous function
type LambdaFunction struct {
	argList       []*symbols.Symbol
	capturedScope environment.Scope
	body          *cons.Cons
}

// NewLambdaFunction creates a new lambda function
func NewLambdaFunction(argList []*symbols.Symbol, capturedScope environment.Scope, body *cons.Cons) *LambdaFunction {
	return &LambdaFunction{
		argList:       argList,
		capturedScope: capturedScope,
		body:          body,
	}
}

// NumArgs number of arguments
func (fun *LambdaFunction) NumArgs() int {
	return len(fun.argList)
}

// EvalArgs evaluate args
func (fun *LambdaFunction) EvalArgs() bool {
	return true
}

// Type of Function
func (fun *LambdaFunction) Type() types.Type {
	return types.Function
}

// String for Stringer
func (fun *LambdaFunction) String() string {
	return fmt.Sprintf("lambda(%p)", fun)
}

// Eval lambda function evaluation
func (fun *LambdaFunction) Eval(args *cons.Cons, env environment.Environment) (types.Object, error) {
	// First push captured scope
	env.PushScope(fun.capturedScope)

	// Push local scope for bound input arguments
	env.PushScope(nil)

	// Pop both local and captured scopes, even when an error occurs
	defer func() {
		env.PopScope()
		env.PopScope()
	}()

	// Bind arguments
	for _, sym := range fun.argList {
		if sym.Reserved {
			return nil, fmt.Errorf("can't bind to reserved symbol %v", sym)
		}

		env.AddBinding(sym, args.Car)

		if args.Cdr.Type() == types.Cons {
			args = args.Cdr.(*cons.Cons)
		} else {
			args = nil
		}
	}

	// Add binding for &rest symbol with remaining args
	if args != nil {
		env.AddBinding(globals.AndRestSymbol, args)
	} else {
		env.AddBinding(globals.AndRestSymbol, types.NIL)
	}

	// Evaluate body
	var result types.Object = types.NIL
	var err error

	if fun.body != nil {
		err = fun.body.Iter(func(obj types.Object, index uint64) error {
			result, err = env.Eval(obj)
			return err
		})

		if err != nil {
			return nil, err
		}
	}

	// Return the result
	return result, nil
}
