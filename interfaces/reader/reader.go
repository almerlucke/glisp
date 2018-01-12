package reader

import (
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

// UntilCondition for nextRunes
type UntilCondition func(rune) (bool, error)

// MacroFunction defines a function being called when a macro character is
// encountered
type MacroFunction func(Reader) (types.Object, error)

// DispatchMacroFunction defines a function being called when a dispatch macro
// character is encountered. The standard dispatch char is #, after that an
// optional sequence of digits is read, after the sequence the next character is
// a selector for the macro function called
type DispatchMacroFunction func(uint64, Reader) (types.Object, error)

// Character holds reader character info
type Character struct {
	SyntaxType      int
	Char            rune
	ReservedForUser bool
	Macro           MacroFunction
}

// ReadTable is a map from rune to character info
type ReadTable map[rune]*Character

// DispatchTable is a map from rune to a dispatch macro function
type DispatchTable map[rune]DispatchMacroFunction

// Reader interface
type Reader interface {
	Error(msg string) error
	ErrorWithError(err error) error
	ReadChar() (rune, *Character, error)
	UnreadChar() error
	NextRunes(n int, until UntilCondition) ([]rune, error)
	IsWhitespace(c rune) bool
	IsNewline(c rune) bool
	DispatchMacroForCharacter(c *Character) DispatchMacroFunction
	ParseToken(bool) (string, error)
	ReadObject() (types.Object, error)
	Context() map[string]interface{}
}
