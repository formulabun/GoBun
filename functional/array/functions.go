package array

func Map[T any, P any](array []T, function func(T) P) []P {
	result := make([]P, len(array))
	for i, t := range array {
		result[i] = function(t)
	}
	return result
}

func FindFirstIndex[T any](array []T, function func(T) bool) int {
	for i, t := range array {
		if function(t) {
			return i
		}
	}
	return -1
}

func FindFirstIndexMatching[T comparable](array []T, toSearch T) int {
  return FindFirstIndex(array, equals(toSearch))
}

func FindFirst[T any](array []T, function func(T) bool) *T {
  i := FindFirstIndex(array, function)
  if i >= 0 && i < len(array) {
    return &array[i]
  }
  return nil
}

func Any[T any](array []T, function func(T) bool) bool {
	for _, t := range array {
		if function(t) {
			return true
		}
	}
	return false
}

func All[T any](array []T, function func(T) bool) bool {
	return !Any(array, inverse(function))
}

func Filter[T any](array []T, function func(T) bool) []T {
  var i = 0
	result := make([]T, len(array))
  for _, t := range array {
    if function(t) {
      result[i] = t
      i++
    }
  }
  return result[:i]
}
