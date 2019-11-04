package compare

// FindMissingElement compares two slices of different length,
// and returns the extra element between the two
func FindMissingElement(s1, s2 []int) int {
	s := append(s1, s2...)
	missing := 0
	for _, v := range s {
		missing = missing ^ v // Bitwise XOR operator between two same integers reverts 0
	}
	return missing
}
