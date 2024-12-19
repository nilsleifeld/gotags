package h

func If[T any](condition bool, e T) T {
	if condition {
		return e
	}
	var zero T
	return zero
}

func ForEach[T any](values []T, f func(T) HTML) []HTML {
	var result []HTML
	for _, v := range values {
		result = append(result, f(v))
	}
	return result
}
