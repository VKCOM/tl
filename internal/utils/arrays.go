package utils

func IntRange(from, to, step int) []int {
	r := make([]int, 0)
	for i := from; i < to; i += step {
		r = append(r, i)
	}
	return r
}
