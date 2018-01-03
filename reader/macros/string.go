package macros

import (
	"container/list"
	"errors"
	"fmt"
	"math"
	"strconv"
	"unicode"

	"github.com/almerlucke/glisp/interfaces/reader"
	"github.com/almerlucke/glisp/reader/utils"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/strings"
)

func runeToHexValue(r rune) uint32 {
	var v uint32

	if r > 96 {
		v = uint32(r - 87)
	} else if r > 64 {
		v = uint32(r - 55)
	} else {
		v = uint32(r - 48)
	}

	return v
}

func isOctal(r rune) bool {
	return r >= 48 && r < 56
}

func isHexadecimal(r rune) bool {
	return unicode.Is(unicode.ASCII_Hex_Digit, r)
}

func octalEscape(startDigit rune, rd reader.Reader) (rune, error) {
	rs, err := rd.NextRunes(2, func(r rune) (bool, error) {
		return !isOctal(r), nil
	})

	o := uint(startDigit - 48)
	for _, r := range rs {
		o = (o << 3) + uint(r-48)
	}

	return rune(o), err
}

func hexadecimalEscape(rd reader.Reader) (rune, error) {
	rs, err := rd.NextRunes(math.MaxInt32, func(r rune) (bool, error) {
		return !isHexadecimal(r), nil
	})

	hex := uint32(0)
	for _, r := range rs {
		hex = (hex << 4) + runeToHexValue(r)
	}

	return rune(hex), err
}

func unicodeEscape(n int, escapeChar rune, rd reader.Reader) ([]rune, error) {
	rs, err := rd.NextRunes(n, func(r rune) (bool, error) {
		return !isHexadecimal(r), nil
	})

	if err != nil {
		return nil, err
	}

	if len(rs) != n {
		return nil, errors.New("unicode char literal expected %d ASCII chars")
	}

	str := fmt.Sprintf("'\\%c%s'", escapeChar, string(rs))

	str, err = strconv.Unquote(str)
	if err != nil {
		return nil, err
	}

	return []rune(str), err
}

func escapeSequence(rd reader.Reader) ([]rune, error) {
	c, _, err := rd.ReadChar()
	if err != nil {
		return nil, err
	}

	var s []rune

	switch c {
	case 'a':
		c = 0x07
	case 'b':
		c = 0x08
	case 'f':
		c = 0x0C
	case 'n':
		c = 0x0A
	case 'r':
		c = 0x0D
	case 't':
		c = 0x09
	case 'v':
		c = 0x0B
	case 'x':
		c, err = hexadecimalEscape(rd)
	case 'u':
		s, err = unicodeEscape(4, 'u', rd)
	case 'U':
		s, err = unicodeEscape(8, 'U', rd)
	default:
		if isOctal(c) {
			c, err = octalEscape(c, rd)
		}
	}

	// If we only have one rune, wrap it in a slice
	if s == nil {
		s = []rune{c}
	}

	return s, err
}

// StringMacro macro for parsing a string object from stream
func StringMacro(rd reader.Reader) (types.Object, error) {
	l := list.New()

	for true {
		c, _, err := rd.ReadChar()
		if err != nil {
			return nil, err
		}

		if c == '"' {
			break
		} else if rd.IsNewline(c) {
			return nil, errors.New("multiline string not allowed")
		} else if c == '\\' {
			s, err := escapeSequence(rd)
			if err != nil {
				return nil, err
			}

			for _, r := range s {
				l.PushBack(r)
			}
		} else {
			l.PushBack(c)
		}
	}

	return strings.String(utils.RuneListToSlice(l)), nil
}
