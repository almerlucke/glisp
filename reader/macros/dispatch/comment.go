package dispatch

import (
	"errors"
	"io"

	"github.com/almerlucke/glisp/reader"
	"github.com/almerlucke/glisp/types"
)

func commentErr(err error) error {
	if err == io.EOF {
		return errors.New("End of stream reached before end of comment")
	}

	return err
}

// CommentDispatch multi line comment dispatch macro
func CommentDispatch(arg uint64, rd *reader.Reader) (types.Object, error) {
	nestedLevel := 1

	for {
		r, _, err := rd.ReadChar()
		if err != nil {
			return nil, commentErr(err)
		}

		if r == '#' {
			r, _, err = rd.ReadChar()
			if err != nil {
				return nil, commentErr(err)
			}

			if r == '|' {
				nestedLevel++
			}
		} else if r == '|' {
			r, _, err = rd.ReadChar()
			if err != nil {
				return nil, commentErr(err)
			}

			if r == '#' {
				nestedLevel--
				if nestedLevel == 0 {
					break
				}
			}
		}
	}

	return nil, nil
}
