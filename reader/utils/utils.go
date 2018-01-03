package utils

import (
	"container/list"
	"regexp"
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
