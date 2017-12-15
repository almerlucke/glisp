package reader

import (
	"container/list"
	"errors"
	"fmt"
	"io"
	"reflect"
	"regexp"
	"strconv"
	"unicode"

	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/numbers"
)

const (
	// Backspace character
	Backspace = 0x8
	// Tab character
	Tab = 0x9
	// Newline character
	Newline = 0xA
	// Page character
	Page = 0xC
	// Return character
	Return = 0xD
	// Space character
	Space = 0x20
	// Delete character
	Delete = 0x7F
)

const (
	// Constituent syntax type
	Constituent = iota
	// Whitespace syntax type
	Whitespace
	// TerminatingMacro syntax type
	TerminatingMacro
	// NonTerminatingMacro syntax type
	NonTerminatingMacro
	// SingleEscape syntax type
	SingleEscape
	// MultipleEscape syntax type
	MultipleEscape
)

// untilCondition for nextRunes
type untilCondition func(rune) (bool, error)

// MacroFunction defines a function being called when a macro character is
// encountered
type MacroFunction func(reader *Reader) (types.Object, error)

// Character holds reader character info
type Character struct {
	SyntaxType      int
	Char            rune
	ReservedForUser bool
	Macro           MacroFunction
}

// Reader implements the Lisp reader
type Reader struct {
	scanner   io.RuneScanner
	readTable ReadTable
	lineCount int
	charCount int
	env       *environment.Environment
	depth     int
}

// Error specific for the reader
func (reader *Reader) Error(msg string) error {
	return fmt.Errorf("Error on line %d, char %d: %s",
		reader.lineCount+1,
		reader.charCount+1,
		msg,
	)
}

// ErrorWithError specific for the reader
func (reader *Reader) ErrorWithError(err error) error {
	if err == io.EOF {
		return err
	}

	return reader.Error(err.Error())
}

// New creates a new reader with the default read table
func New(scanner io.RuneScanner, env *environment.Environment) *Reader {
	return &Reader{
		scanner:   scanner,
		readTable: DefaultReadTable,
		env:       env,
	}
}

// ReadChar reads a single character from the stream
func (reader *Reader) ReadChar() (rune, *Character, error) {
	c, _, err := reader.scanner.ReadRune()

	if err != nil {
		return c, nil, err
	}

	if c == Newline {
		reader.newLine()
	} else if c == Return {
		reader.newLine()

		// We are going to check for CR-LF sequence
		c, _, err = reader.scanner.ReadRune()
		if err != nil {
			return 0, nil, err
		}

		// If c is not a linefeed, unread
		if c != Newline {
			c = Return
			err = reader.scanner.UnreadRune()
			if err != nil {
				return 0, nil, err
			}
		}
	} else {
		reader.charCount++
	}

	ci := reader.readTable[c]

	return c, ci, err
}

func (reader *Reader) newLine() {
	reader.lineCount++
	reader.charCount = 0
}

// UnreadChar unreads a single character from the stream
func (reader *Reader) UnreadChar() error {
	reader.charCount--
	return reader.scanner.UnreadRune()
}

func (reader *Reader) nextRunes(n int, until untilCondition) ([]rune, error) {
	rl := list.New()

	for n > 0 {
		c, _, err := reader.ReadChar()
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
			err = reader.UnreadChar()
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

// IsWhitespace check if char is whitespace
func (reader *Reader) IsWhitespace(c rune) bool {
	ci := reader.readTable[c]

	return ci != nil && ci.SyntaxType == Whitespace
}

// IsNewline check if char is considered a newline
func (reader *Reader) IsNewline(c rune) bool {
	return c == Newline || c == Return
}

// SkipWhitespace skip whitespace
func (reader *Reader) skipWhitespace() error {
	c, _, err := reader.ReadChar()

	// Skip while space runes
	for err == nil && reader.IsWhitespace(c) {
		c, _, err = reader.ReadChar()
	}

	if err == nil {
		// Unread last rune
		err = reader.UnreadChar()
	}

	return err
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

// IsInteger returns true if string represents an integer
func IsInteger(s string) bool {
	reg := regexp.MustCompile(`^[+-]?[0-9]+$`)

	return reg.MatchString(s)
}

// IsFloat returns true if string represents a float
func IsFloat(s string) bool {
	reg := regexp.MustCompile(`^[-+]?[0-9]*\.?[0-9]+([eE][-+]?[0-9]+)?$`)

	return reg.MatchString(s)
}

func (reader *Reader) tokenToObject(token string) types.Object {
	if IsInteger(token) {
		i, err := strconv.ParseInt(token, 10, 64)
		if err != nil {
			i = 0
		}

		return &numbers.Number{
			Kind:  reflect.Int64,
			Value: i,
		}
	} else if IsFloat(token) {
		f, err := strconv.ParseFloat(token, 64)
		if err != nil {
			f = 0
		}

		return &numbers.Number{
			Kind:  reflect.Float64,
			Value: f,
		}
	}

	sym := reader.env.DefineSymbol(token, false, nil)
	if sym.Value != nil {
		// Symbol has a self referencing value, return value instead of symbol
		return sym.Value
	}

	return sym
}

func (reader *Reader) parseToken() (types.Object, error) {
	cl := list.New()
	singleEscapeActive := false
	multipleEscapeActive := false
	c, ci, err := reader.ReadChar()

	for {
		if err != nil {
			if err == io.EOF {
				if singleEscapeActive || multipleEscapeActive {
					return nil, errors.New("EOF reached before end of escape")
				}

				break
			}

			return nil, reader.ErrorWithError(err)
		}

		if ci == nil {
			return nil, fmt.Errorf("Illegal character %c found", c)
		}

		if singleEscapeActive {
			cl.PushBack(c)
			singleEscapeActive = false
		} else if multipleEscapeActive {
			if ci.SyntaxType == MultipleEscape {
				multipleEscapeActive = false
			} else {
				cl.PushBack(c)
			}
		} else {
			if ci.SyntaxType == Constituent {
				cl.PushBack(unicode.ToUpper(c))
			} else if ci.SyntaxType == SingleEscape {
				singleEscapeActive = true
			} else if ci.SyntaxType == MultipleEscape {
				multipleEscapeActive = true
			} else if ci.SyntaxType == NonTerminatingMacro {
				cl.PushBack(c)
			} else if ci.SyntaxType == TerminatingMacro || ci.SyntaxType == Whitespace {
				err = reader.UnreadChar()
				if err != nil {
					return nil, err
				}
				break
			}
		}

		c, ci, err = reader.ReadChar()
	}

	return reader.tokenToObject(string(runeListToSlice(cl))), nil
}

func (reader *Reader) parseNonTerminatingMacro() (types.Object, error) {
	return nil, nil
}

func (reader *Reader) Read() (types.Object, error) {
	// First skip whitespace
	err := reader.skipWhitespace()
	if err != nil {
		return nil, err
	}

	c, ci, err := reader.ReadChar()
	if err != nil {
		return nil, err
	}

	if ci == nil {
		return nil, fmt.Errorf("Illegal character %c found", c)
	}

	var obj types.Object

	switch ci.SyntaxType {
	case SingleEscape:
		fallthrough

	case MultipleEscape:
		fallthrough

	case Constituent:
		err = reader.UnreadChar()
		if err != nil {
			return nil, err
		}

		obj, err = reader.parseToken()
		if err != nil && err != io.EOF {
			return nil, err
		}

	case TerminatingMacro:
		if ci.Macro != nil {
			obj, err = ci.Macro(reader)
			if err != nil {
				return nil, err
			}

			if obj == nil {
				return reader.Read()
			}
		} else {
			return nil, fmt.Errorf("No macro function attached to macro char %c", c)
		}

	case NonTerminatingMacro:
		obj, err = reader.parseNonTerminatingMacro()
		if err != nil {
			return nil, err
		}
	}

	return obj, nil
}
