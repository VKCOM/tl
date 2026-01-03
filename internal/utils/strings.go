package utils

import "strings"

func UpperFirst(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func LowerFirst(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func ShiftAll(s []string, shift string) []string {
	r := make([]string, len(s))
	for i, e := range s {
		r[i] = shift + e
	}
	return r
}

func Append(res []string, values ...string) []string {
	resCopy := make([]string, len(res))
	copy(resCopy, res)
	resCopy = append(resCopy, values...)
	return resCopy
}
