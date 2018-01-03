package functions

import (
	"fmt"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/interfaces/function"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/symbols"

	globals "github.com/almerlucke/glisp/globals/symbols"
)

// LambdaEvalContext holds lambda and macro call depth
type LambdaEvalContext struct {
	Depth uint64
}

// LambdaFunction anonymous function
type LambdaFunction struct {
	argList       []*symbols.Symbol
	capturedScope environment.Scope
	body          *cons.Cons
}

// getLambdaEvalContext gets lambda call context
func getLambdaEvalContext(env environment.Environment) *LambdaEvalContext {
	ctx, ok := env.Context()["LambdaEvalContext"]
	if ok {
		return ctx.(*LambdaEvalContext)
	}

	ctx = &LambdaEvalContext{}
	env.Context()["LambdaEvalContext"] = ctx

	return ctx.(*LambdaEvalContext)
}

// PushLambdaEvalCall push
func PushLambdaEvalCall(env environment.Environment) {
	ctx := getLambdaEvalContext(env)
	ctx.Depth++
}

// PopLambdaEvalCall pop
func PopLambdaEvalCall(env environment.Environment) {
	ctx := getLambdaEvalContext(env)
	ctx.Depth--
}

// IsInsideLambdaEval return true if lambda eval context depth is not 0
func IsInsideLambdaEval(env environment.Environment) bool {
	ctx := getLambdaEvalContext(env)
	return ctx.Depth > 0
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
func (fun *LambdaFunction) Eval(args *cons.Cons, env environment.Environment, context interface{}) (result types.Object, err error) {
	// First push captured scope
	env.PushScope(fun.capturedScope)

	// Push local scope for bound input arguments
	env.PushScope(nil)

	// Push call
	PushLambdaEvalCall(env)

	// Pop both local and captured scopes, even when an error occurs
	defer func() {
		// Pop scopes
		env.PopScope()
		env.PopScope()

		// Pop call
		PopLambdaEvalCall(env)

		if r := recover(); r != nil {
			// Return value
			returnContext, ok := r.(*function.ReturnContext)
			if ok {
				result = returnContext.Object
			} else {
				// Continue to panic
				panic(r)
			}
		}
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

	if fun.body != nil {
		err = fun.body.Iter(func(obj types.Object, index uint64) error {
			result, err = env.Eval(obj, context)
			return err
		})

		if err != nil {
			return nil, err
		}
	}

	// Return the result
	return result, nil
}
