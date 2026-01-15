package utils

import (
	"regexp"
	"strings"
)

var (
	camelingRegex = regexp.MustCompile(`[0-9A-Za-z]+`)
	allUpperRegex = regexp.MustCompile(`^[A-Z][A-Z0-9]+$`)
)

// TODO - investigate if this function is good
func CNameToCamelName(s string) string {
	chunks := camelingRegex.FindAllString(s, -1)
	for i, chunk := range chunks {
		if allUpperRegex.MatchString(chunk) { // TODO - why?
			chunks[i] = strings.ToUpper(chunk[:1]) + strings.ToLower(chunk[1:])
		} else {
			chunks[i] = ToUpperFirst(chunk)
		}
	}
	return strings.Join(chunks, "")
}

func ToUpperFirst(str string) string {
	for i := range str {
		if i != 0 {
			return strings.ToUpper(str[:i]) + str[i:]
		}
	}
	return strings.ToUpper(str) // zero or single rune
}

func ToLowerFirst(str string) string {
	for i := range str {
		if i != 0 {
			return strings.ToLower(str[:i]) + str[i:]
		}
	}
	return strings.ToLower(str) // zero or single rune
}
