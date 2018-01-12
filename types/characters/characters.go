package characters

import (
	"fmt"
	"strconv"
	"unicode/utf8"

	"github.com/almerlucke/glisp/interfaces/reader"
	"github.com/almerlucke/glisp/reader/utils"
	"github.com/almerlucke/glisp/types"
)

// Character type
type Character rune

// NewWithToken create a new character from a token
func NewWithToken(token string) (Character, error) {
	switch token {
	case "Backspace":
		return Character(reader.Backspace), nil
	case "Newline":
		return Character(reader.Newline), nil
	case "Tab":
		return Character(reader.Tab), nil
	case "Rubout":
		return Character(reader.Delete), nil
	case "Page":
		return Character(reader.Page), nil
	case "Return":
		return Character(reader.Return), nil
	case "Space":
		return Character(reader.Space), nil
	}

	if utils.IsSmallUnicodeLiteral(token) || utils.IsLargeUnicodeLiteral(token) {
		literal := fmt.Sprintf("'\\%s'", token)
		literal, err := strconv.Unquote(literal)
		if err != nil {
			return 0, err
		}

		return Character([]rune(literal)[0]), nil
	}

	len := utf8.RuneCountInString(token)

	if len > 1 {
		return 0, fmt.Errorf("illegal character name %v", token)
	}

	return Character([]rune(token)[0]), nil
}

// Type Character for Object interface
func (c Character) Type() types.Type {
	return types.Character
}

// String for stringer interface
func (c Character) String() string {
	if c == reader.Backspace {
		return fmt.Sprintf("#\\Backspace")
	} else if c == reader.Newline {
		return fmt.Sprintf("#\\Newline")
	} else if c == reader.Tab {
		return fmt.Sprintf("#\\Tab")
	} else if c == reader.Delete {
		return fmt.Sprintf("#\\Rubout")
	} else if c == reader.Page {
		return fmt.Sprintf("#\\Page")
	} else if c == reader.Return {
		return fmt.Sprintf("#\\Return")
	} else if c == reader.Space {
		return fmt.Sprintf("#\\Space")
	}

	return fmt.Sprintf("#\\%c", rune(c))
}

// Eql check
func (c Character) Eql(obj types.Object) bool {
	return c == obj
}

// Equal check
func (c Character) Equal(obj types.Object) bool {
	return c.Eql(obj)
}
