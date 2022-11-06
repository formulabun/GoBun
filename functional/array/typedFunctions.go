package array

func AnyBool(array []bool) bool {
	return Any(array, identity[bool])
}

func AllBool(array []bool) bool {
	return All(array, identity[bool])
}
