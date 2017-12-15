package reader

import (
	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/cons"
)

func quoteMacro(reader *Reader) (types.Object, error) {
	var obj types.Object
	var err error

	for {
		obj, err = reader.Read()
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
