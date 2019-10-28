package compare

// FindMissingElement compares two slices of different length,
// and returns the extra element between the two
func FindMissingElement(s1, s2 []int) int {
	// Find which slice is longer
	var long, short []int
	if len(s1) > len(s2) {
		long, short = s1, s2
	} else {
		long, short = s2, s1
	}
	var missing int

	for _, v1 := range long {
		for _, v2 := range short {
			if v1 == v2 {
				// found = true
				break
			}
			missing = v1
		}
	}
	return missing

}
