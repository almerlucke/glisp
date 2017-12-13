package reader

import (
	"container/list"
	"errors"
	"fmt"
	"io"
	"unicode"

	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/types"
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
	return reader.scanner.UnreadRune()
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
func (reader *Reader) SkipWhitespace() error {
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

func (reader *Reader) parseSymbol() (types.Object, error) {
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

	return reader.env.DefineSymbol(string(runeListToSlice(cl)), false), nil
}

func (reader *Reader) parseNonTerminatingMacro() (types.Object, error) {
	return nil, nil
}

func (reader *Reader) Read() (types.Object, error) {
	// First skip whitespace
	err := reader.SkipWhitespace()
	if err != nil {
		return nil, reader.ErrorWithError(err)
	}

	c, ci, err := reader.ReadChar()
	if err != nil {
		return nil, reader.ErrorWithError(err)
	}

	if ci == nil {
		return nil, reader.Error(fmt.Sprintf("Illegal character %c found", c))
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
			return nil, reader.ErrorWithError(err)
		}

		obj, err = reader.parseSymbol()
		if err != nil && err != io.EOF {
			return nil, reader.ErrorWithError(err)
		}

	case TerminatingMacro:
		if ci.Macro != nil {
			obj, err = ci.Macro(reader)
			if err != nil {
				return nil, reader.ErrorWithError(err)
			}
		} else {
			return nil, reader.Error(fmt.Sprintf("No macro function attached to macro char %c", c))
		}

	case NonTerminatingMacro:
		obj, err = reader.parseNonTerminatingMacro()
		if err != nil {
			return nil, reader.ErrorWithError(err)
		}
	}

	return obj, nil
}
