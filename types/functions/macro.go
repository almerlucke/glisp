package functions

import (
	"fmt"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/interfaces/function"
	"github.com/almerlucke/glisp/scope"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/symbols"

	globals "github.com/almerlucke/glisp/globals/symbols"
)

// MacroFunction is like a lambda function except args are
// not evaluated and the result of the initial evaluation (expansion)
// is evaluated again
type MacroFunction struct {
	argList       []*symbols.Symbol
	capturedScope scope.Scope
	body          *cons.Cons
}

// NewMacroFunction creates a new macro function
func NewMacroFunction(argList []*symbols.Symbol, capturedScope scope.Scope, body *cons.Cons) *MacroFunction {
	return &MacroFunction{
		argList:       argList,
		capturedScope: capturedScope,
		body:          body,
	}
}

// EvalArgs evaluate args
func (fun *MacroFunction) EvalArgs() bool {
	return false
}

// NumArgs number of arguments
func (fun *MacroFunction) NumArgs() int {
	return len(fun.argList)
}

// Type of Function
func (fun *MacroFunction) Type() types.Type {
	return types.Function
}

// String for Stringer
func (fun *MacroFunction) String() string {
	return fmt.Sprintf("macro(%p)", fun)
}

// Macro expansion
func (fun *MacroFunction) expansion(args *cons.Cons, env environment.Environment, context interface{}) (result types.Object, err error) {
	// First push captured scope
	env.PushScope(fun.capturedScope)

	// Push local scope for bound input arguments
	env.PushScope(nil)

	// Push call
	env.PushDepthContext("CallDepth")

	defer func() {
		// Pop scopes
		env.PopScope()
		env.PopScope()

		// Pop call
		env.PopDepthContext("CallDepth")

		if r := recover(); r != nil {
			// Return value
			returnContext, ok := r.(*function.ReturnContext)
			if ok {
				// Return the evaluation of the macro expansion
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

	// Expand macro body
	if fun.body != nil {
		err = fun.body.Iter(func(obj types.Object, index interface{}) (bool, error) {
			result, err = env.Eval(obj, context)
			return false, err
		})

		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Eval lambda function evaluation
func (fun *MacroFunction) Eval(args *cons.Cons, env environment.Environment, context interface{}) (result types.Object, err error) {
	result, err = fun.expansion(args, env, context)
	if err != nil {
		return nil, err
	}

	// Return the evaluation of the macro expansion
	return env.Eval(result, context)
}
