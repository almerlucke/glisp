package tokenizer

import (
	"bufio"
	"bytes"
	"container/list"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	_linefeed       = 0xA
	_carriageReturn = 0xD
)

const (
	// Unknown token
	Unknown = iota
	// Comment token
	Comment
	// OpenParenthesis token
	OpenParenthesis
	// CloseParenthesis token
	CloseParenthesis
	// BackTick token
	BackTick
	// Keyword token
	Keyword
	// String token
	String
	// Integer token
	Integer
	// Float token
	Float
)

// untilCondition for nextRunes
type untilCondition func(rune) bool

// Token type
type Token struct {
	Type    int
	Content string
}

// Tokenizer turns a stream of runes in tokens
type Tokenizer struct {
	f         *os.File
	r         io.RuneScanner
	lineCount int
	charCount int
}

// Error specific for the tokenizer
func (tokenizer *Tokenizer) Error(msg string) error {
	return fmt.Errorf("Error on line %d, char %d: %s",
		tokenizer.lineCount,
		tokenizer.charCount,
		msg,
	)
}

// CreateWithFile creates a new tokenizer from a file path
func CreateWithFile(file string) (*Tokenizer, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	t := CreateWithRuneScanner(bufio.NewReader(f))
	t.f = f

	return t, nil
}

// CreateWithString creates a new tokenizer with a string
func CreateWithString(s string) *Tokenizer {
	return CreateWithRuneScanner(strings.NewReader(s))
}

// CreateWithBytes creates a new tokenizer with a byte slice
func CreateWithBytes(b []byte) *Tokenizer {
	return CreateWithRuneScanner(bytes.NewBuffer(b))
}

// CreateWithRuneScanner creates a new tokenizer with a runescanner
func CreateWithRuneScanner(reader io.RuneScanner) *Tokenizer {
	return &Tokenizer{
		f: nil,
		r: reader,
	}
}

// Destroy a tokenizer
func (tokenizer *Tokenizer) Destroy() {
	if tokenizer.f != nil {
		tokenizer.f.Close()
	}
}

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

func runeListToSlice(l *list.List) []rune {
	rs := make([]rune, l.Len())
	rc := 0
	for e := l.Front(); e != nil; e = e.Next() {
		rs[rc] = e.Value.(rune)
		rc++
	}

	return rs
}

func (tokenizer *Tokenizer) read() (rune, error) {
	c, _, err := tokenizer.r.ReadRune()

	if err != nil {
		return c, err
	}

	if c == _linefeed {
		tokenizer.newLine()
	} else if c == _carriageReturn {
		tokenizer.newLine()

		// We are going to check for CR-LF sequence
		c, _, err = tokenizer.r.ReadRune()
		if err != nil {
			return 0, err
		}

		// If c is not a linefeed, unread
		if c != _linefeed {
			c = _linefeed
			err = tokenizer.unread()
			if err != nil {
				return 0, err
			}
		}
	} else {
		tokenizer.charCount++
	}

	return c, err
}

func (tokenizer *Tokenizer) unread() error {
	return tokenizer.r.UnreadRune()
}

func (tokenizer *Tokenizer) newLine() {
	tokenizer.lineCount++
	tokenizer.charCount = 0
}

func (tokenizer *Tokenizer) skipWhitespace() error {
	c, err := tokenizer.read()

	// Skip while space runes
	for err == nil && unicode.IsSpace(c) {
		c, err = tokenizer.read()
	}

	if err == nil {
		// Unread last rune
		err = tokenizer.unread()
	}

	return err
}

func (tokenizer *Tokenizer) nextRunes(n int, until untilCondition) ([]rune, error) {
	rl := list.New()

	for n > 0 {
		c, err := tokenizer.read()
		if err != nil {
			// return err and slice, err can be EOF
			return runeListToSlice(rl), err
		}

		if until(c) {
			// Unread the last rune that caused until to return true
			err = tokenizer.unread()
			if err != nil {
				return runeListToSlice(rl), nil
			}

			break
		}

		rl.PushBack(c)
		n--
	}

	return runeListToSlice(rl), nil
}

func (tokenizer *Tokenizer) octalEscape(startDigit rune) (rune, error) {
	rs, err := tokenizer.nextRunes(2, func(r rune) bool {
		return !isOctal(r)
	})

	o := uint(startDigit - 48)
	for _, r := range rs {
		o = (o << 3) + uint(r-48)
	}

	return rune(o), err
}

func (tokenizer *Tokenizer) hexadecimalEscape() (rune, error) {
	rs, err := tokenizer.nextRunes(math.MaxInt32, func(r rune) bool {
		return !isHexadecimal(r)
	})

	hex := uint32(0)
	for _, r := range rs {
		hex = (hex << 4) + runeToHexValue(r)
	}

	return rune(hex), err
}

func (tokenizer *Tokenizer) unicodeEscape(n int, escapeChar rune) ([]rune, error) {
	rs, err := tokenizer.nextRunes(n, func(r rune) bool {
		return !isHexadecimal(r)
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

// NextToken unless an error occurred, error can be io.EOF
func (tokenizer *Tokenizer) NextToken() (*Token, error) {
	err := tokenizer.skipWhitespace()
	if err != nil {
		return nil, tokenizer.Error(err.Error())
	}

	c, err := tokenizer.read()
	if err != nil {
		return nil, tokenizer.Error(err.Error())
	}

	var token *Token

	switch c {
	case '(':
		token = &Token{Type: OpenParenthesis}
	case ')':
		token = &Token{Type: CloseParenthesis}
	case '"':
		token, err = tokenizer.parseDoubleQuoteString()
		if err == io.EOF {
			return nil, tokenizer.Error("end of file reached before end of string")
		} else if err != nil {
			return nil, tokenizer.Error(err.Error())
		}
	case '/':
		token = &Token{Type: OpenParenthesis}
	default:
		token = &Token{Type: OpenParenthesis}
	}

	return token, err
}
