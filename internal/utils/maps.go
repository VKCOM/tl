package utils

func SliceToSet[T comparable](s *[]T) map[T]bool {
	m := make(map[T]bool)
	for _, e := range *s {
		m[e] = true
	}
	return m
}

// unstable
func Keys[K comparable, V any](m *map[K]V) (res []K) {
	for k, _ := range *m {
		res = append(res, k)
	}
	return
}

// unstable
func SetToSlice[T comparable](s *map[T]bool) []T {
	m := make([]T, 0)
	for k, _ := range *s {
		m = append(m, k)
	}
	return m
}

func MapSlice[A, B any](in *[]A, f func(A) B) (out []B) {
	for _, e := range *in {
		out = append(out, f(e))
	}
	return
}

func FilterSlice[A any](in []A, f func(A) bool) (out []A) {
	for _, e := range in {
		if f(e) {
			out = append(out, e)
		}
	}
	return
}

func PutPairToSetOfPairs[K, V comparable](in *map[K]map[V]bool, k K, v V) {
	if _, ok := (*in)[k]; !ok {
		(*in)[k] = make(map[V]bool)
	}
	(*in)[k][v] = true
}

func ReverseSetOfPairs[K, V comparable](in map[K]map[V]bool) map[V]map[K]bool {
	m := make(map[V]map[K]bool)

	for k, vs := range in {
		for v, _ := range vs {
			PutPairToSetOfPairs(&m, v, k)
		}
	}

	return m
}
