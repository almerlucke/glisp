package macros

import (
	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/reader"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
)

// QuoteMacro quote an object
func QuoteMacro(reader *reader.Reader) (types.Object, error) {
	var obj types.Object
	var err error

	for {
		obj, err = reader.ReadObject()
		if err != nil {
			return nil, err
		}

		if obj != nil {
			break
		}
	}

	return &cons.Cons{
		Car: environment.QuoteSymbol,
		Cdr: &cons.Cons{
			Car: obj,
			Cdr: types.NIL,
		},
	}, nil
}
