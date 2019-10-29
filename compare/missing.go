package compare

// FindMissingElement compares two slices of different length,
// and returns the extra element between the two
func FindMissingElement(s1, s2 []int) int {
	s := append(s1,s2...)
	for len(s) > 1 {
		// value := s[0]
		for i, v := range s[:] {
			if i == 0 {
				continue
			}
			if s[0] == v { // Found duplicate
				s[i], s[len(s)-1] = s[len(s)-1], s[i] // Put duplicate value in the end
				s = s[:len(s)-1] // Re-slice to exlude the last element (duplicates))
			}
		}
	}
	return s[0]
	}