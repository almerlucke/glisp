package macros

import (
	"github.com/almerlucke/glisp/reader"
	"github.com/almerlucke/glisp/types"
)

// CommentMacro skips all chars after ; until newline
func CommentMacro(rd *reader.Reader) (types.Object, error) {
	c, _, err := rd.ReadChar()

	for err == nil && c != reader.Newline {
		// Skip all chars
		c, _, err = rd.ReadChar()
	}

	return nil, err
}
