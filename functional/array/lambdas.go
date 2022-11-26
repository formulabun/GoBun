package array

func identity[T any](b T) T {
	return b
}

func equals[T comparable](a T) (func (T) bool) {
  return func (b T) bool {
    return a == b
  }
}

func inverse[T any](function func(T) bool) func(T) bool {
	return func(t T) bool {
		return !function(t)
	}
}
