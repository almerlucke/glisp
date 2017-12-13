package tokenizer

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
	"io"
	"os"
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
	// Identifier token
	Identifier
	// String token
	String
	// Integer token
	Integer
	// Float token
	Float
)

// untilCondition for nextRunes
type untilCondition func(rune) (bool, error)

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

func runeListToSlice(l *list.List) []rune {
	rs := make([]rune, l.Len())
	rc := 0
	for e := l.Front(); e != nil; e = e.Next() {
		rs[rc] = e.Value.(rune)
		rc++
	}

	return rs
}

func isTerminatingChar(r rune) bool {
	return r == '(' || r == ')' || unicode.IsSpace(r) || r == '"' || r == '`' || r == ','
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

		// Check until condition for stop signal or error
		stop, err := until(c)
		if err != nil {
			return runeListToSlice(rl), err
		}

		if stop {
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
		if isStartOfIdentifier(c) {
			token, err = tokenizer.parseIdentifier(c)
		} else {
			token = &Token{Type: OpenParenthesis}
		}
	}

	return token, err
}
