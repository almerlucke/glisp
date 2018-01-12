package utils

import (
	"container/list"
	"regexp"
	"strings"
	"unicode"
)

// IsSmallUnicodeLiteral check for small unicode literal
func IsSmallUnicodeLiteral(token string) bool {
	smallUnicode := regexp.MustCompile(`^u[0-9a-fA-F]{4}$`)

	return smallUnicode.MatchString(token)
}

// IsLargeUnicodeLiteral check for large unicode literal
func IsLargeUnicodeLiteral(token string) bool {
	largeUnicode := regexp.MustCompile(`^U[0-9a-fA-F]{8}$`)

	return largeUnicode.MatchString(token)
}

// RuneToHexValue rune to hex value
func RuneToHexValue(r rune) uint32 {
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

// IsOctal check if rune is octal
func IsOctal(r rune) bool {
	return r >= 48 && r < 56
}

// IsHexadecimal check if rune is hexadecimal
func IsHexadecimal(r rune) bool {
	return unicode.Is(unicode.ASCII_Hex_Digit, r)
}

// RuneListToSlice converts a list of runes to a slice of runes
func RuneListToSlice(l *list.List) []rune {
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

// IsKeyword check if symbol name is keyword
func IsKeyword(s string) bool {
	return strings.HasPrefix(s, ":")
}

// SplitNamespace splits a symbol by namespace and returns the namespace,
// the symbol and a bool indicating if we need to look at all symbols or
// only exported ones
func SplitNamespace(s string) (string, string, bool) {
	components := strings.Split(s, "::")

	if len(components) > 1 {
		return components[0], components[1], true
	}

	components = strings.Split(s, ":")

	if len(components) > 1 {
		return components[0], components[1], false
	}

	return "", s, false
}
