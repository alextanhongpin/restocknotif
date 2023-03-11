package slices

func Map[K, V any](s []K, fn func(i int) V) []V {
	res := make([]V, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = fn(i)
	}

	return res
}
