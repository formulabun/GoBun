package strings

// NullTerminatedString returns a string from a null-terminated byte slice,
// removing any bytes over 127.
func SafeNullTerminated(data []byte) string {
	newBytes := make([]byte, 0)

	// Filter out all bytes above 127 and cut at a null terminator
	for i := 0; i < len(data); i++ {
		if data[i] >= 128 {
			continue
		} else if data[i] == 0 {
			break
		}
		newBytes = append(newBytes, data[i])
	}
	return string(newBytes)
}

// NullTerminatedString returns a slice of bytes from a null-terminated byte slice.
func NullTerminated(data []byte) []byte {
	newBytes := make([]byte, 0)

	// Cut at a null terminator
	for i := 0; i < len(data); i++ {
		if data[i] == 0 {
			break
		}
		newBytes = append(newBytes, data[i])
	}
	return newBytes
}
