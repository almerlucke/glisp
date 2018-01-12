package reader

import (
	"container/list"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"github.com/almerlucke/glisp/interfaces/environment"
	"github.com/almerlucke/glisp/interfaces/reader"
	"github.com/almerlucke/glisp/reader/utils"
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/numbers"
	"github.com/almerlucke/glisp/types/symbols"
)

// Reader implements the Lisp reader interface
type Reader struct {
	scanner        io.RuneScanner
	readTable      reader.ReadTable
	dispatchTable  reader.DispatchTable
	lineCount      int
	charCount      int
	env            environment.Environment
	context        map[string]interface{}
	listDepth      int
	backquoteDepth int
}

// Context returns the reader context
func (rd *Reader) Context() map[string]interface{} {
	return rd.context
}

// Error specific for the reader
func (rd *Reader) Error(msg string) error {
	return fmt.Errorf("error on line %d, char %d: %s",
		rd.lineCount+1,
		rd.charCount+1,
		msg,
	)
}

// ErrorWithError specific for the reader
func (rd *Reader) ErrorWithError(err error) error {
	if err == io.EOF {
		return err
	}

	return rd.Error(err.Error())
}

// New creates a new reader
func New(scanner io.RuneScanner, readTable reader.ReadTable, dispatchTable reader.DispatchTable, env environment.Environment) *Reader {
	return &Reader{
		scanner:       scanner,
		readTable:     readTable,
		dispatchTable: dispatchTable,
		env:           env,
		context:       map[string]interface{}{},
	}
}

// ReadChar reads a single character from the stream
func (rd *Reader) ReadChar() (rune, *reader.Character, error) {
	c, _, err := rd.scanner.ReadRune()

	if err != nil {
		return c, nil, err
	}

	if c == reader.Newline {
		rd.newLine()
	} else if c == reader.Return {
		rd.newLine()

		// We are going to check for CR-LF sequence
		c, _, err = rd.scanner.ReadRune()
		if err != nil {
			return 0, nil, err
		}

		// If c is not a linefeed, unread
		if c != reader.Newline {
			// Return is always converted to newline
			c = reader.Newline
			err = rd.scanner.UnreadRune()
			if err != nil {
				return 0, nil, err
			}
		}
	} else {
		rd.charCount++
	}

	ci := rd.readTable[c]

	return c, ci, err
}

func (rd *Reader) newLine() {
	rd.lineCount++
	rd.charCount = 0
}

// UnreadChar unreads a single character from the stream
func (rd *Reader) UnreadChar() error {
	rd.charCount--
	return rd.scanner.UnreadRune()
}

// NextRunes gets the next n runes from stream until condition
func (rd *Reader) NextRunes(n int, until reader.UntilCondition) ([]rune, error) {
	rl := list.New()

	for n > 0 {
		c, _, err := rd.ReadChar()
		if err != nil {
			// return err and slice, err can be EOF
			return utils.RuneListToSlice(rl), err
		}

		// Check until condition for stop signal or error
		stop, err := until(c)
		if err != nil {
			return utils.RuneListToSlice(rl), err
		}

		if stop {
			// Unread the last rune that caused until to return true
			err = rd.UnreadChar()
			if err != nil {
				return utils.RuneListToSlice(rl), nil
			}

			break
		}

		rl.PushBack(c)
		n--
	}

	return utils.RuneListToSlice(rl), nil
}

// IsWhitespace check if char is whitespace
func (rd *Reader) IsWhitespace(c rune) bool {
	ci := rd.readTable[c]

	return ci != nil && ci.SyntaxType == reader.Whitespace
}

// IsNewline check if char is considered a newline
func (rd *Reader) IsNewline(c rune) bool {
	return c == reader.Newline || c == reader.Return
}

// SkipWhitespace skip whitespace
func (rd *Reader) skipWhitespace() error {
	c, _, err := rd.ReadChar()

	// Skip while space runes
	for err == nil && rd.IsWhitespace(c) {
		c, _, err = rd.ReadChar()
	}

	if err == nil {
		// Unread last rune
		err = rd.UnreadChar()
	}

	return err
}

func (rd *Reader) tokenToObject(token string) (types.Object, error) {
	if utils.IsInteger(token) {
		i, err := strconv.ParseInt(token, 10, 64)
		if err != nil {
			i = 0
		}

		return &numbers.Number{
			Kind:  reflect.Int64,
			Value: numbers.Int64(i),
		}, nil
	} else if utils.IsFloat(token) {
		f, err := strconv.ParseFloat(token, 64)
		if err != nil {
			f = 0
		}

		return &numbers.Number{
			Kind:  reflect.Float64,
			Value: numbers.Float64(f),
		}, nil
	}

	var sym *symbols.Symbol

	if utils.IsKeyword(token) {
		sym = rd.env.InternKeyword(strings.TrimPrefix(token, ":"))
	} else {
		ns, name, interned := utils.SplitNamespace(token)
		if ns != "" {
			if interned {
				sym = rd.env.FindInternedSymbolInNamespace(name, ns)
			} else {
				sym = rd.env.FindExportedSymbolInNamespace(name, ns)
			}

			if sym == nil {
				return nil, fmt.Errorf("unknown symbol %v in namespace %v", name, ns)
			}
		} else {
			sym = rd.env.InternSymbol(token)
		}
	}

	if sym.Value != nil {
		// Symbol has a self referencing value, return value instead of symbol
		return sym.Value, nil
	}

	return sym, nil
}

// ParseToken parse token
func (rd *Reader) ParseToken(extended bool) (string, error) {
	cl := list.New()
	singleEscapeActive := false
	multipleEscapeActive := false
	c, ci, err := rd.ReadChar()

	for {
		if err != nil {
			if err == io.EOF {
				if singleEscapeActive || multipleEscapeActive {
					return "", errors.New("end of stream reached before end of escape")
				}

				break
			}

			return "", rd.ErrorWithError(err)
		}

		if ci == nil {
			return "", fmt.Errorf("illegal character %c found", c)
		}

		if singleEscapeActive {
			cl.PushBack(c)
			singleEscapeActive = false
		} else if multipleEscapeActive {
			if ci.SyntaxType == reader.MultipleEscape {
				multipleEscapeActive = false
			} else {
				cl.PushBack(c)
			}
		} else {
			if ci.SyntaxType == reader.Constituent {
				if extended {
					cl.PushBack(c)
				} else {
					cl.PushBack(unicode.ToUpper(c))
				}
			} else if ci.SyntaxType == reader.SingleEscape {
				singleEscapeActive = true
			} else if ci.SyntaxType == reader.MultipleEscape {
				multipleEscapeActive = true
			} else if ci.SyntaxType == reader.NonTerminatingMacro {
				cl.PushBack(c)
			} else if ci.SyntaxType == reader.TerminatingMacro || ci.SyntaxType == reader.Whitespace {
				err = rd.UnreadChar()
				if err != nil {
					return "", err
				}
				break
			}
		}

		c, ci, err = rd.ReadChar()
	}

	return string(utils.RuneListToSlice(cl)), nil
}

// DispatchMacroForCharacter returns a dispatch macro for a character
// can return nil if no dispatch macro is defined
func (rd *Reader) DispatchMacroForCharacter(c *reader.Character) reader.DispatchMacroFunction {
	return rd.dispatchTable[unicode.ToLower(c.Char)]
}

// ReadObject reads an object from the stream
func (rd *Reader) ReadObject() (types.Object, error) {
	// First skip whitespace
	err := rd.skipWhitespace()
	if err != nil {
		return nil, err
	}

	c, ci, err := rd.ReadChar()
	if err != nil {
		return nil, err
	}

	if ci == nil {
		return nil, fmt.Errorf("illegal character %c found", c)
	}

	var obj types.Object

	switch ci.SyntaxType {
	case reader.SingleEscape:
		fallthrough

	case reader.MultipleEscape:
		fallthrough

	case reader.Constituent:
		err = rd.UnreadChar()
		if err != nil {
			return nil, err
		}

		var token string
		token, err = rd.ParseToken(false)
		if err != nil && err != io.EOF {
			return nil, err
		}

		obj, err = rd.tokenToObject(token)
		if err != nil {
			return nil, err
		}

	case reader.NonTerminatingMacro:
		fallthrough
	case reader.TerminatingMacro:
		if ci.Macro != nil {
			obj, err = ci.Macro(rd)
			if err != nil {
				return nil, err
			}

			if obj == nil {
				return rd.ReadObject()
			}
		} else {
			return nil, fmt.Errorf("no macro function attached to macro char %c", c)
		}
	}

	return obj, nil
}
