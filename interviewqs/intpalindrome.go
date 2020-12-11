package interviewqs

// return int is palindrome without stronv.Itoa
// 12321 -> true
// 1221 -> true
// 1232 -> false
func isIntPalindrome(v int) bool {
	f, pow := first(v)
	if pow == 0 {
		return true
	}

	last := v % 10
	if f != last {
		return false
	}

	m := 1
	for i := 0; i < pow; i++ {
		m *= 10
	}

	return isIntPalindrome(v - last - f*m)
}

func first(v int) (int, int) {
	pow := 0
	for {
		nv := v % 10
		if nv == v {
			return pow, v
		}
		v = nv
		pow++
	}
}
