package utils

import (
	"hash/fnv"
)

func NameID(s string) int {
	h := fnv.New64a() // FNV-1a 64-bit
	_, _ = h.Write([]byte(s))
	r := int(h.Sum64())
	if r < 0 {
		r = -r
	}
	return r
}
