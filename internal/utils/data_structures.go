package utils

// WARINING: USE IF S IS SMALL AND IF YOU NOT FACE ALLOCATION PROBLEM
func SetSubSets[T comparable](s map[T]bool) (values []map[T]bool) {
	if len(s) == 0 {
		values = append(values, make(map[T]bool))
	} else {
		var e T
		for x := range s {
			e = x
			break
		}
		copyS := CopyMap(s)
		delete(copyS, e)
		copySSubsets := SetSubSets(copyS)
		for _, subset := range copySSubsets {
			values = append(values, CopyMap(subset))
		}
		for _, subset := range copySSubsets {
			c := CopyMap(subset)
			c[e] = true
			values = append(values, c)
		}
	}
	return
}

func SliceToSet[T comparable](s []T) map[T]bool {
	m := make(map[T]bool)
	for _, e := range s {
		m[e] = true
	}
	return m
}

func SliceToMap[K comparable, T, V any](s []T, key func(T) K, value func(T) V) map[K]V {
	m := make(map[K]V)
	for _, t := range s {
		m[key(t)] = value(t)
	}
	return m
}

func CopyMap[K comparable, V any](m map[K]V) map[K]V {
	m2 := make(map[K]V)
	for k, v := range m {
		m2[k] = v
	}
	return m2
}

// unstable
func Keys[K comparable, V any](m map[K]V) (res []K) {
	for k := range m {
		res = append(res, k)
	}
	return
}

// unstable
func Values[K comparable, V any](m map[K]V) (res []V) {
	for _, v := range m {
		res = append(res, v)
	}
	return
}

// unstable
func SetToSlice[T comparable](s map[T]bool) []T {
	m := make([]T, 0)
	for k := range s {
		m = append(m, k)
	}
	return m
}

func MapSlice[A, B any](in []A, f func(A) B) (out []B) {
	for _, e := range in {
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
		for v := range vs {
			PutPairToSetOfPairs(&m, v, k)
		}
	}

	return m
}

func AppendMap[K comparable, V any](values map[K]V, to *map[K]V) {
	for k, v := range values {
		(*to)[k] = v
	}
}

func AppendMapWithResolving[K comparable, V any](values map[K]V, to *map[K]V, resolver func(key K, value1, value2 V) V) {
	for k, v := range values {
		if v2, ok := (*to)[k]; ok {
			(*to)[k] = resolver(k, v, v2)
		} else {
			(*to)[k] = v
		}
	}
}

func SetIntersection[K comparable](s1, s2 map[K]bool) map[K]bool {
	result := make(map[K]bool)
	origin, target := s1, s2
	if len(s2) < len(s1) {
		origin, target = s2, s1
	}
	for key := range origin {
		if _, ok := target[key]; ok {
			result[key] = true
		}
	}
	return result
}

func SetUnion[K comparable](s1, s2 map[K]bool) map[K]bool {
	result := make(map[K]bool)
	for key := range s1 {
		result[key] = true
	}
	for key := range s2 {
		result[key] = true
	}
	return result
}
