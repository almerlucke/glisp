package utils

import (
	"container/list"
	"regexp"
	"strings"
)

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
