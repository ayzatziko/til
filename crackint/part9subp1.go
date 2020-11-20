package crackint

// 1.1 is unique
func isUniqueChars(s string) bool {
	u := map[byte]struct{}{}

	for _, c := range []byte(s) {
		if _, ok := u[c]; ok {
			return false
		}

		u[c] = struct{}{}
	}

	return true
}

// 1.2 check permutations
func checkPermutation(a, b string) bool {
	chars := map[byte]int{}

	for _, c := range []byte(a) {
		chars[c]++
	}

	for _, c := range []byte(b) {
		chars[c]--
		if chars[c] < 0 {
			return false
		}
	}

	for _, diff := range chars {
		if diff != 0 {
			return false
		}
	}

	return true
}

// 1.3 urlify
func urlify(s string) string {
	var b []byte

	const space = byte(' ')
	rs := []byte("%20")

	for _, c := range []byte(s) {
		if c == space {
			b = append(b, rs...)
			continue
		}

		b = append(b, c)
	}

	return string(b)
}
