package array

func identity[T any](b T) T {
  return b
}

func AnyBool(array []bool) bool {
  return Any(array, identity[bool])
}

func AllBool(array []bool) bool {
  return All(array, identity[bool])
}
