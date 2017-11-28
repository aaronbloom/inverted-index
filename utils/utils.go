package utils

import (
	"strings"
)

// PathSplit splits strings based upon common file delimiters
func PathSplit(r rune) bool {
	return r == '/' || r == '\\' || r == ' ' || r == '.'
}

// NeutralString is a basic whitespace tokeniser and lowercase filter on strings
func NeutralString(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	return s
}
