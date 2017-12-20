package macros

import (
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
	"unicode"

	"github.com/almerlucke/glisp/reader"
	"github.com/almerlucke/glisp/types"
)

// DispatchMacro dispatches a macro based on the character after the sharpsign
func DispatchMacro(rd *reader.Reader) (types.Object, error) {
	// Get optional digits
	digits, err := rd.NextRunes(math.MaxInt32, func(c rune) (bool, error) {
		return !unicode.IsDigit(c), nil
	})

	if err != nil {
		if err == io.EOF {
			return nil, errors.New("end of stream reached before end of dispatch macro")
		}

		return nil, err
	}

	arg, _ := strconv.ParseInt(string(digits), 10, 64)

	r, c, err := rd.ReadChar()
	if err != nil {
		if err == io.EOF {
			return nil, errors.New("end of stream reached before end of dispatch macro")
		}

		return nil, err
	}

	macro := rd.DispatchMacroForCharacter(c)
	if macro == nil {
		return nil, fmt.Errorf("undefined dispatch macro character %c", r)
	}

	return macro(uint64(arg), rd)
}
