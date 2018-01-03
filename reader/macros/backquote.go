package macros

import (
	"errors"
	"io"

	"github.com/almerlucke/glisp/globals/symbols"
	"github.com/almerlucke/glisp/interfaces/reader"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
)

// BackquoteContext contains backquote depth
type BackquoteContext struct {
	Depth int
}

// BackquoteMacro backquote reader,
func BackquoteMacro(rd reader.Reader) (types.Object, error) {
	var obj types.Object
	var err error

	var bctx *BackquoteContext
	ctx, ok := rd.Context()["backquoteContext"]
	if !ok {
		bctx = &BackquoteContext{}
		rd.Context()["backquoteContext"] = bctx
	} else {
		bctx = ctx.(*BackquoteContext)
	}

	bctx.Depth++

	for {
		obj, err = rd.ReadObject()
		if err != nil {
			if err == io.EOF {
				return nil, errors.New("end of stream before end of backquote")
			}
			return nil, err
		}

		if obj != nil {
			break
		}
	}

	bctx.Depth--

	return &cons.Cons{
		Car: symbols.BackquoteSymbol,
		Cdr: &cons.Cons{
			Car: obj,
			Cdr: types.NIL,
		},
	}, nil
}
