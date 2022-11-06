package array

func identity[T any](b T) T {
	return b
}

func inverse[T any](function func(T) bool) func(T) bool {
	return func(t T) bool {
		return !function(t)
	}
}
