package macros

import (
	"errors"
	"io"

	"github.com/almerlucke/glisp/globals/symbols"
	"github.com/almerlucke/glisp/interfaces/reader"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
)

type listContext struct {
	Depth int
}

// OpenParenthesisMacro is called when an open parenthesis is encountered
func OpenParenthesisMacro(rd reader.Reader) (types.Object, error) {
	dotFound := false
	dottedObjCnt := 0
	builder := cons.ListBuilder{}

	var lctx *listContext
	ctx, ok := rd.Context()["listContext"]
	if !ok {
		lctx = &listContext{}
		rd.Context()["listContext"] = lctx
	} else {
		lctx = ctx.(*listContext)
	}

	lctx.Depth++

	for {
		obj, err := rd.ReadObject()
		if err != nil {
			if err == io.EOF {
				return nil, errors.New("unmatched parenthesis")
			}

			return nil, err
		}

		if obj == symbols.CloseParenthesisSymbol {
			if dotFound && dottedObjCnt != 1 {
				return nil, errors.New("expected one object after dot")
			}

			break
		}

		if obj == symbols.DotSymbol {
			dotFound = true
		} else if obj != nil {
			if dotFound {
				if builder.Tail == nil {
					return nil, errors.New("expected at least one object before dot")
				}
				dottedObjCnt++
				builder.Tail.Cdr = obj
			} else {
				builder.PushBack(&cons.Cons{
					Car: obj,
					Cdr: types.NIL,
				})
			}
		}
	}

	lctx.Depth--

	if builder.Head == nil {
		return types.NIL, nil
	}

	return builder.Head, nil
}

// CloseParenthesisMacro is called when a closing parenthesis is encountered
func CloseParenthesisMacro(rd reader.Reader) (types.Object, error) {
	ctx, ok := rd.Context()["listContext"]
	if !ok || ctx.(*listContext).Depth == 0 {
		return nil, errors.New("unmatched parenthesis")
	}

	return symbols.CloseParenthesisSymbol, nil
}
