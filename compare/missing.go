package compare

// FindMissingElement compares two slices of different length,
// and returns the extra element between the two
func FindMissingElement(s1, s2 []int) int {
	// Find sum of elements in first slice
	sum1 := 0
	for _, v := range s1 {
		sum1 += v
	}

	// Find sum of elements in second slice
	sum2 := 0
	for _, v := range s2 {
		sum2 += v
	}

	// Taking into consideration the case where the extra number is < 0
	// The slice with greater lenght has the extra number
	if len(s1) > len(s2) {
		return sum1 - sum2
	}
	return sum2 - sum1

}
