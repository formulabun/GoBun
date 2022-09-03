package strings

type Stringer struct {
	S string
}

func (s Stringer) String() string {
	return s.S
}

func Join(strings []string, seperator string) (result string) {
	var sep = ""
	for _, s := range strings {
		result += sep + s
		sep = seperator
	}
	return result
}
