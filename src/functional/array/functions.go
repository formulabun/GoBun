package array

type Possible = interface{}

func Map[T any, P any](array []T, function func(T)(P))([]P) {
  result := make([]P, len(array))
  for i, t := range array {
    result[i] = function(t)
  }
  return result
}

func FindFirst[T any](array []T, function func(T)(bool)) *T {
  for _, t := range array {
    if function(t) {
      return &t
    }
  }
  return nil
}

func Any[T any](array []T, function func(T)(bool)) bool {
  for _, t := range array {
    if function(t) {
      return true
    }
  }
  return false
}

func All[T any](array []T, function func(T)(bool)) bool {
  return ! Any(array, inverse(function))
}
