package simpledi

// sourandset
// check if set a contains every item of b or not, givent algorithms run in O(n^2)
func sourandset(a []input, b []input) bool {
	if len(b) > len(a) {
		return false
	}

	for _, ai := range a {
		found := false
		for _, bi := range b {
			if ai.typ == bi.typ {
				found = true
			}
		}

		if !found {
			return false
		}
	}

	return true
}
