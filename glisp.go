package glisp

import (
	"github.com/almerlucke/glisp/environment"
	"github.com/almerlucke/glisp/reader"
	"github.com/almerlucke/glisp/reader/macros"
	"github.com/almerlucke/glisp/reader/macros/dispatch"
	"github.com/almerlucke/glisp/types/functions/buildin"
)

// DefaultReadTable contains the default reader characters and syntax types
var DefaultReadTable = generateDefaultReadTable()

// DefaultDispatchTable contains the default reader dispatch table
var DefaultDispatchTable = generateDefaultDispatchTable()

// CreateDefaultEnvironment creates a default GLisp environment
func CreateDefaultEnvironment() *environment.Environment {
	env := environment.New()

	env.AddBinding(environment.QuoteSymbol, buildin.CreateBuildinQuote())
	env.AddBinding(environment.BackquoteSymbol, buildin.CreateBuildinBackquote())

	return env
}

func generateDefaultDispatchTable() reader.DispatchTable {
	table := map[rune]reader.DispatchMacroFunction{
		'|': dispatch.CommentDispatch,
	}

	return table
}

// GenerateDefaultReadTable returns the default read table
func generateDefaultReadTable() reader.ReadTable {
	table := map[rune]*reader.Character{
		reader.Backspace: &reader.Character{
			SyntaxType: reader.Constituent,
			Char:       reader.Backspace,
		},
		reader.Tab: &reader.Character{
			SyntaxType: reader.Whitespace,
			Char:       reader.Tab,
		},
		reader.Newline: &reader.Character{
			SyntaxType: reader.Whitespace,
			Char:       reader.Newline,
		},
		reader.Page: &reader.Character{
			SyntaxType: reader.Whitespace,
			Char:       reader.Page,
		},
		reader.Return: &reader.Character{
			SyntaxType: reader.Whitespace,
			Char:       reader.Return,
		},
		reader.Space: &reader.Character{
			SyntaxType: reader.Whitespace,
			Char:       reader.Space,
		},
		'!': &reader.Character{
			SyntaxType:      reader.Constituent,
			Char:            '!',
			ReservedForUser: true,
		},
		'"': &reader.Character{
			SyntaxType: reader.TerminatingMacro,
			Char:       '"',
			Macro:      macros.StringMacro,
		},
		'#': &reader.Character{
			SyntaxType: reader.NonTerminatingMacro,
			Char:       '#',
			Macro:      macros.DispatchMacro,
		},
		'$': &reader.Character{
			SyntaxType: reader.Constituent,
			Char:       '$',
		},
		'%': &reader.Character{
			SyntaxType: reader.Constituent,
			Char:       '%',
		},
		'&': &reader.Character{
			SyntaxType: reader.Constituent,
			Char:       '&',
		},
		'\'': &reader.Character{
			SyntaxType: reader.TerminatingMacro,
			Char:       '\'',
			Macro:      macros.QuoteMacro,
		},
		'(': &reader.Character{
			SyntaxType: reader.TerminatingMacro,
			Char:       '(',
			Macro:      macros.OpenParenthesisMacro,
		},
		')': &reader.Character{
			SyntaxType: reader.TerminatingMacro,
			Char:       ')',
			Macro:      macros.CloseParenthesisMacro,
		},
		'*': &reader.Character{
			SyntaxType: reader.Constituent,
			Char:       '*',
		},
		'+': &reader.Character{
			SyntaxType: reader.Constituent,
			Char:       '+',
		},
		',': &reader.Character{
			SyntaxType: reader.TerminatingMacro,
			Char:       ',',
			Macro:      macros.UnquoteMacro,
		},
		'-': &reader.Character{
			SyntaxType: reader.Constituent,
			Char:       '-',
		},
		'.': &reader.Character{
			SyntaxType: reader.Constituent,
			Char:       '.',
		},
		'/': &reader.Character{
			SyntaxType: reader.Constituent,
			Char:       '/',
		},
		':': &reader.Character{
			SyntaxType: reader.Constituent,
			Char:       ':',
		},
		';': &reader.Character{
			SyntaxType: reader.TerminatingMacro,
			Char:       ';',
			Macro:      macros.CommentMacro,
		},
		'<': &reader.Character{
			SyntaxType: reader.Constituent,
			Char:       '<',
		},
		'=': &reader.Character{
			SyntaxType: reader.Constituent,
			Char:       '=',
		},
		'>': &reader.Character{
			SyntaxType: reader.Constituent,
			Char:       '>',
		},
		'?': &reader.Character{
			SyntaxType:      reader.Constituent,
			Char:            '?',
			ReservedForUser: true,
		},
		'@': &reader.Character{
			SyntaxType: reader.Constituent,
			Char:       '@',
		},
		'[': &reader.Character{
			SyntaxType:      reader.Constituent,
			Char:            '[',
			ReservedForUser: true,
		},
		'\\': &reader.Character{
			SyntaxType: reader.SingleEscape,
			Char:       '\\',
		},
		']': &reader.Character{
			SyntaxType:      reader.Constituent,
			Char:            ']',
			ReservedForUser: true,
		},
		'^': &reader.Character{
			SyntaxType: reader.Constituent,
			Char:       '^',
		},
		'_': &reader.Character{
			SyntaxType: reader.Constituent,
			Char:       '_',
		},
		'`': &reader.Character{
			SyntaxType: reader.TerminatingMacro,
			Char:       '`',
			Macro:      macros.BackquoteMacro,
		},
		'{': &reader.Character{
			SyntaxType:      reader.Constituent,
			Char:            '{',
			ReservedForUser: true,
		},
		'|': &reader.Character{
			SyntaxType: reader.MultipleEscape,
			Char:       '|',
		},
		'}': &reader.Character{
			SyntaxType:      reader.Constituent,
			Char:            '}',
			ReservedForUser: true,
		},
		'~': &reader.Character{
			SyntaxType: reader.Constituent,
			Char:       '~',
		},
		reader.Delete: &reader.Character{
			SyntaxType: reader.Constituent,
			Char:       reader.Delete,
		},
	}

	// Add digits as constituents
	for i := 0; i < 10; i++ {
		c := rune(i + 48)
		table[c] = &reader.Character{
			SyntaxType: reader.Constituent,
			Char:       c,
		}
	}

	// Add uppercase and lowercase letters as constituents
	for i := 0; i < 26; i++ {

		// Uppercase
		c := rune(i + 65)
		table[c] = &reader.Character{
			SyntaxType: reader.Constituent,
			Char:       c,
		}

		// Lowercase
		c = rune(i + 97)
		table[c] = &reader.Character{
			SyntaxType: reader.Constituent,
			Char:       c,
		}
	}

	return table
}
