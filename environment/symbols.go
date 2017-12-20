package environment

import (
	"github.com/almerlucke/glisp/types"
	"github.com/almerlucke/glisp/types/symbols"
)

// DotSymbol is used for dotted lists in the reader
var DotSymbol = &symbols.Symbol{
	Name:     ".",
	Reserved: true,
}

// CloseParenthesisSymbol is used to signal a closing parenthesis
// in the OpenParenthesisMacro
var CloseParenthesisSymbol = &symbols.Symbol{
	Name:     ")",
	Reserved: true,
}

// NILSymbol always references NIL instead of the symbol
var NILSymbol = &symbols.Symbol{
	Name:     "NIL",
	Reserved: true,
	Value:    types.NIL,
}

// TSymbol always references T instead of the symbol
var TSymbol = &symbols.Symbol{
	Name:     "T",
	Reserved: true,
	Value:    types.T,
}

// QuoteSymbol is used for quoted objects
var QuoteSymbol = &symbols.Symbol{
	Name:     "SYSTEM::QUOTE",
	Reserved: true,
}

// BackquoteSymbol is used for backquoted objects
var BackquoteSymbol = &symbols.Symbol{
	Name:     "SYSTEM::BACKQUOTE",
	Reserved: true,
}

// UnquoteSymbol is used for unquoting
var UnquoteSymbol = &symbols.Symbol{
	Name:     "SYSTEM::UNQUOTE",
	Reserved: true,
}

// SpliceSymbol is used for splicing
var SpliceSymbol = &symbols.Symbol{
	Name:     "SYSTEM::SPLICE",
	Reserved: true,
}

// AndRestSymbol is used to bind extra arguments in lambda function calls
var AndRestSymbol = &symbols.Symbol{
	Name:     "&REST",
	Reserved: true,
}
