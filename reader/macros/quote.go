package macros

import (
	"errors"
	"io"

	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/reader"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
)

// QuoteMacro quote an object
func QuoteMacro(rd *reader.Reader) (types.Object, error) {
	var obj types.Object
	var err error

	for {
		obj, err = rd.ReadObject()
		if err != nil {
			if err == io.EOF {
				return nil, errors.New("end of stream reached before end of quote")
			}
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
