package macros

import (
	"errors"
	"io"

	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/reader"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
)

// BackquoteMacro backquote reader,
func BackquoteMacro(rd *reader.Reader) (types.Object, error) {
	var obj types.Object
	var err error

	rd.BackquoteDepth++

	for {
		obj, err = rd.ReadObject()
		if err != nil {
			if err == io.EOF {
				return nil, errors.New("End of stream before end of backquote")
			}
			return nil, err
		}

		if obj != nil {
			break
		}
	}

	rd.BackquoteDepth--

	return &cons.Cons{
		Car: environment.BackquoteSymbol,
		Cdr: &cons.Cons{
			Car: obj,
			Cdr: types.NIL,
		},
	}, nil
}
