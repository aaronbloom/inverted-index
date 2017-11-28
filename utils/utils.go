package utils

import (
	"strings"
)

// FilePathTokenizer is a basic file delimiting tokeniser
func FilePathTokenizer(r rune) bool {
	return r == '/' || r == '\\' || r == ' ' || r == '.' || r == '-'
}

// LowerCaseFilter is a lowercase filter on strings
func LowerCaseFilter(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	return s
}
