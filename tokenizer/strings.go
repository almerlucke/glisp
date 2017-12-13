package tokenizer

import (
	"container/list"
	"errors"
	"fmt"
	"math"
	"strconv"
	"unicode"
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

func (tokenizer *Tokenizer) octalEscape(startDigit rune) (rune, error) {
	rs, err := tokenizer.nextRunes(2, func(r rune) (bool, error) {
		return !isOctal(r), nil
	})

	o := uint(startDigit - 48)
	for _, r := range rs {
		o = (o << 3) + uint(r-48)
	}

	return rune(o), err
}

func (tokenizer *Tokenizer) hexadecimalEscape() (rune, error) {
	rs, err := tokenizer.nextRunes(math.MaxInt32, func(r rune) (bool, error) {
		return !isHexadecimal(r), nil
	})

	hex := uint32(0)
	for _, r := range rs {
		hex = (hex << 4) + runeToHexValue(r)
	}

	return rune(hex), err
}

func (tokenizer *Tokenizer) unicodeEscape(n int, escapeChar rune) ([]rune, error) {
	rs, err := tokenizer.nextRunes(n, func(r rune) (bool, error) {
		return !isHexadecimal(r), nil
	})

	if err != nil {
		return nil, err
	}

	if len(rs) != n {
		return nil, errors.New("Unicode char literal expected %d ASCII chars")
	}

	str := fmt.Sprintf("'\\%c%s'", escapeChar, string(rs))

	str, err = strconv.Unquote(str)
	if err != nil {
		return nil, err
	}

	return []rune(str), err
}

func (tokenizer *Tokenizer) escapeSequence() ([]rune, error) {
	c, err := tokenizer.read()
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
		c, err = tokenizer.hexadecimalEscape()
	case 'u':
		s, err = tokenizer.unicodeEscape(4, 'u')
	case 'U':
		s, err = tokenizer.unicodeEscape(8, 'U')
	default:
		if isOctal(c) {
			c, err = tokenizer.octalEscape(c)
		}
	}

	// If we only have one rune, wrap it in a slice
	if s == nil {
		s = []rune{c}
	}

	return s, err
}

func (tokenizer *Tokenizer) parseDoubleQuoteString() (*Token, error) {
	l := list.New()

	for true {
		c, err := tokenizer.read()
		if err != nil {
			return nil, err
		}

		if c == '"' {
			break
		} else if c == _linefeed {
			return nil, errors.New("Multiline string not allowed")
		} else if c == '\\' {
			s, err := tokenizer.escapeSequence()
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

	return &Token{
		Type:    String,
		Content: string(runeListToSlice(l)),
	}, nil
}
