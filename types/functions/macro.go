package functions

import (
	"fmt"

	"github.com/almerlucke/glisp/interfaces/environment"
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
	capturedScope environment.Scope
	body          *cons.Cons
}

// NewMacroFunction creates a new macro function
func NewMacroFunction(argList []*symbols.Symbol, capturedScope environment.Scope, body *cons.Cons) *MacroFunction {
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

// Eval lambda function evaluation
func (fun *MacroFunction) Eval(args *cons.Cons, env environment.Environment) (types.Object, error) {
	restore := func() {
		env.PopScope()
		env.PopScope()
	}

	// First push captured scope
	env.PushScope(fun.capturedScope)

	// Push local scope for bound input arguments
	env.PushScope(nil)

	// Bind arguments
	for _, sym := range fun.argList {
		if sym.Reserved {
			restore()
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
			restore()
			return nil, err
		}
	}

	// Restore environment scope
	restore()

	// Return the evaluation of the macro expansion
	return env.Eval(result)
}
