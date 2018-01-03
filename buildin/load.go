package buildin

import (
	"bufio"
	"errors"
	"io"
	"os"

	defaultReader "github.com/almerlucke/glisp/reader"

	"github.com/almerlucke/glisp/globals/tables"
	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/functions"
	"github.com/almerlucke/glisp/types/strings"
)

// Load buildin function
func Load(args *cons.Cons, env environment.Environment, context interface{}) (types.Object, error) {
	if args.Car.Type() != types.String {
		return nil, errors.New("load expected a path string as first argument")
	}

	path := args.Car.(strings.String)

	f, err := os.Open(string(path))
	if err != nil {
		return nil, err
	}

	defer f.Close()

	readTable := tables.DefaultReadTable
	dispatchTable := tables.DefaultDispatchTable

	rd := defaultReader.New(bufio.NewReader(f), readTable, dispatchTable, env)

	obj, err := rd.ReadObject()
	var result types.Object

	for err == nil {
		result, err = env.Eval(obj, context)
		if err != nil {
			return nil, rd.ErrorWithError(err)
		}

		obj, err = rd.ReadObject()
	}

	if err != nil && err != io.EOF {
		return nil, rd.ErrorWithError(err)
	}

	if result != nil {
		return result, nil
	}

	return types.NIL, nil
}

// CreateBuildinLoad creates a buildin function object
func CreateBuildinLoad() *functions.BuildinFunction {
	return functions.NewBuildinFunction(Load, 0, true)
}
