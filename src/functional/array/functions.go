package array

func Map[T any, P any](array []T, function func(T)(P))([]P) {
  result := make([]P, len(array))
  for i, t := range array {
    result[i] = function(t)
  }
  return result
}
