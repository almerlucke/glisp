package macros

import (
	"errors"
	"io"

	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/reader"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
)

// OpenParenthesisMacro is called when an open parenthesis is encountered
func OpenParenthesisMacro(reader *reader.Reader) (types.Object, error) {
	dotFound := false
	dottedObjCnt := 0
	builder := cons.ListBuilder{}

	reader.Depth++

	for {
		obj, err := reader.ReadObject()
		if err != nil {
			if err == io.EOF {
				return nil, errors.New("unmatched parenthesis")
			}

			return nil, err
		}

		if obj == environment.CloseParenthesisSymbol {
			if dotFound && dottedObjCnt != 1 {
				return nil, errors.New("expected one object after dot")
			}

			break
		}

		if obj == environment.DotSymbol {
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

	reader.Depth--

	if builder.Head == nil {
		return types.NIL, nil
	}

	return builder.Head, nil
}

// CloseParenthesisMacro is called when a closing parenthesis is encountered
func CloseParenthesisMacro(reader *reader.Reader) (types.Object, error) {
	if reader.Depth == 0 {
		return nil, errors.New("unmatched parenthesis")
	}

	return environment.CloseParenthesisSymbol, nil
}
