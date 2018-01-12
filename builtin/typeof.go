package builtin

import (
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
	"github.com/almerlucke/glisp/types/symbols"
)

// TypeOf builtin function
func TypeOf(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	var typeSym *symbols.Symbol

	switch args.Car.Type() {
	case types.Array:
		typeSym = env.InternKeyword("ARRAY")
	case types.Null:
		typeSym = env.InternKeyword("NULL")
	case types.Boolean:
		typeSym = env.InternKeyword("BOOLEAN")
	case types.Cons:
		typeSym = env.InternKeyword("CONS")
	case types.Dictionary:
		typeSym = env.InternKeyword("DICTIONARY")
	case types.Function:
		typeSym = env.InternKeyword("FUNCTION")
	case types.Namespace:
		typeSym = env.InternKeyword("NAMESPACE")
	case types.Number:
		typeSym = env.InternKeyword("NUMBER")
	case types.String:
		typeSym = env.InternKeyword("STRING")
	case types.Symbol:
		typeSym = env.InternKeyword("SYMBOL")
	}

	return typeSym, nil
}

// CreateBuiltinTypeOf creates a builtin function object
func CreateBuiltinTypeOf() *functions.BuiltinFunction {
	return functions.NewBuiltinFunction(TypeOf, 1, true)
}
