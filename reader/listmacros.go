package reader

import (
	"errors"
	"io"

	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
)

// OpenParenthesisMacro is called when an open parenthesis is encountered
func openParenthesisMacro(reader *Reader) (types.Object, error) {
	dotFound := false
	dottedObjCnt := 0
	builder := cons.ListBuilder{}

	reader.depth++

	for {
		obj, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				return nil, errors.New("Unmatched parenthesis")
			}

			return nil, err
		}

		if obj == environment.CloseParenthesisSymbol {
			if dotFound && dottedObjCnt != 1 {
				return nil, errors.New("Expected one object after dot")
			}

			break
		}

		if obj == environment.DotSymbol {
			dotFound = true
		} else if obj != nil {
			if dotFound {
				if builder.Tail == nil {
					return nil, errors.New("Expected at least one object before dot")
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

	reader.depth--

	if builder.Head == nil {
		return types.NIL, nil
	}

	return builder.Head, nil
}

// CloseParenthesisMacro is called when a closing parenthesis is encountered
func closeParenthesisMacro(reader *Reader) (types.Object, error) {
	if reader.depth == 0 {
		return nil, errors.New("Unmatched parenthesis")
	}

	return environment.CloseParenthesisSymbol, nil
}
