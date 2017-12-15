package reader

// ReadTable is a map from rune to character info
type ReadTable map[rune]*Character

// DefaultReadTable contains the default reader characters and syntax types
var DefaultReadTable = generateDefaultReadTable()

// GenerateDefaultReadTable returns the default read table
func generateDefaultReadTable() ReadTable {
	table := map[rune]*Character{
		Backspace: &Character{
			SyntaxType: Constituent,
			Char:       Backspace,
		},
		Tab: &Character{
			SyntaxType: Whitespace,
			Char:       Tab,
		},
		Newline: &Character{
			SyntaxType: Whitespace,
			Char:       Newline,
		},
		Page: &Character{
			SyntaxType: Whitespace,
			Char:       Page,
		},
		Return: &Character{
			SyntaxType: Whitespace,
			Char:       Return,
		},
		Space: &Character{
			SyntaxType: Whitespace,
			Char:       Space,
		},
		'!': &Character{
			SyntaxType:      Constituent,
			Char:            '!',
			ReservedForUser: true,
		},
		'"': &Character{
			SyntaxType: TerminatingMacro,
			Char:       '"',
			Macro:      stringMacro,
		},
		'#': &Character{
			SyntaxType: NonTerminatingMacro,
			Char:       '#',
		},
		'$': &Character{
			SyntaxType: Constituent,
			Char:       '$',
		},
		'%': &Character{
			SyntaxType: Constituent,
			Char:       '%',
		},
		'&': &Character{
			SyntaxType: Constituent,
			Char:       '&',
		},
		'\'': &Character{
			SyntaxType: TerminatingMacro,
			Char:       '\'',
			Macro:      quoteMacro,
		},
		'(': &Character{
			SyntaxType: TerminatingMacro,
			Char:       '(',
			Macro:      openParenthesisMacro,
		},
		')': &Character{
			SyntaxType: TerminatingMacro,
			Char:       ')',
			Macro:      closeParenthesisMacro,
		},
		'*': &Character{
			SyntaxType: Constituent,
			Char:       '*',
		},
		'+': &Character{
			SyntaxType: Constituent,
			Char:       '+',
		},
		',': &Character{
			SyntaxType: TerminatingMacro,
			Char:       ',',
		},
		'-': &Character{
			SyntaxType: Constituent,
			Char:       '-',
		},
		'.': &Character{
			SyntaxType: Constituent,
			Char:       '.',
		},
		'/': &Character{
			SyntaxType: Constituent,
			Char:       '/',
		},
		':': &Character{
			SyntaxType: Constituent,
			Char:       ':',
		},
		';': &Character{
			SyntaxType: TerminatingMacro,
			Char:       ';',
		},
		'<': &Character{
			SyntaxType: Constituent,
			Char:       '<',
		},
		'=': &Character{
			SyntaxType: Constituent,
			Char:       '=',
		},
		'>': &Character{
			SyntaxType: Constituent,
			Char:       '>',
		},
		'?': &Character{
			SyntaxType:      Constituent,
			Char:            '?',
			ReservedForUser: true,
		},
		'@': &Character{
			SyntaxType: Constituent,
			Char:       '@',
		},
		'[': &Character{
			SyntaxType:      Constituent,
			Char:            '[',
			ReservedForUser: true,
		},
		'\\': &Character{
			SyntaxType: SingleEscape,
			Char:       '\\',
		},
		']': &Character{
			SyntaxType:      Constituent,
			Char:            ']',
			ReservedForUser: true,
		},
		'^': &Character{
			SyntaxType: Constituent,
			Char:       '^',
		},
		'_': &Character{
			SyntaxType: Constituent,
			Char:       '_',
		},
		'`': &Character{
			SyntaxType: TerminatingMacro,
			Char:       '`',
		},
		'{': &Character{
			SyntaxType:      Constituent,
			Char:            '{',
			ReservedForUser: true,
		},
		'|': &Character{
			SyntaxType: MultipleEscape,
			Char:       '|',
		},
		'}': &Character{
			SyntaxType:      Constituent,
			Char:            '}',
			ReservedForUser: true,
		},
		'~': &Character{
			SyntaxType: Constituent,
			Char:       '~',
		},
		Delete: &Character{
			SyntaxType: Constituent,
			Char:       Delete,
		},
	}

	// Add digits as constituents
	for i := 0; i < 10; i++ {
		c := rune(i + 48)
		table[c] = &Character{
			SyntaxType: Constituent,
			Char:       c,
		}
	}

	// Add uppercase and lowercase letters as constituents
	for i := 0; i < 26; i++ {

		// Uppercase
		c := rune(i + 65)
		table[c] = &Character{
			SyntaxType: Constituent,
			Char:       c,
		}

		// Lowercase
		c = rune(i + 97)
		table[c] = &Character{
			SyntaxType: Constituent,
			Char:       c,
		}
	}

	return table
}
