package macros

import (
	"errors"
	"io"

	globals "github.com/almerlucke/glisp/globals/symbols"

	"github.com/almerlucke/glisp/interfaces/reader"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
	"github.com/almerlucke/glisp/types/symbols"
)

// UnquoteMacro used for , and ,@ forms in a backquote form
func UnquoteMacro(rd reader.Reader) (types.Object, error) {
	ctx, ok := rd.Context()["backquoteContext"]
	if !ok || ctx.(*BackquoteContext).Depth == 0 {
		return nil, errors.New("unquote and splice can only be used in a backquote form")
	}

	splice := false

	r, _, err := rd.ReadChar()
	if err != nil {
		if err == io.EOF {
			return nil, errors.New("end of stream reached before end of unquote")
		}

		return nil, err
	}

	// check for splice ,@
	if r == '@' {
		splice = true
	} else {
		err = rd.UnreadChar()
		if err != nil {
			return nil, err
		}
	}

	var obj types.Object

	for {
		obj, err = rd.ReadObject()
		if err != nil {
			if err == io.EOF {
				if splice {
					return nil, errors.New("end of stream reached before end of splice")
				}

				return nil, errors.New("end of stream reached before end of unquote")
			}

			return nil, err
		}

		if obj != nil {
			break
		}
	}

	var sym *symbols.Symbol
	if splice {
		sym = globals.SpliceSymbol
	} else {
		sym = globals.UnquoteSymbol
	}

	return &cons.Cons{
		Car: sym,
		Cdr: &cons.Cons{
			Car: obj,
			Cdr: types.NIL,
		},
	}, nil
}
