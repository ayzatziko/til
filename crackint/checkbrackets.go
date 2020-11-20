package crackint

func checkBrackets(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	var opBrackets = map[byte]struct{}{'{': {}, '[': {}, '(': {}}
	var opp = map[byte]byte{']': '[', '}': '{', ')': '('}

	sN := len(s) / 2
	stack := make([]byte, 0, sN)

	for _, b := range []byte(s) {
		if _, ok := opBrackets[b]; ok {
			if len(stack) >= sN {
				return false
			}

			stack = append(stack, b)
			continue
		}

		lastN := len(stack) - 1
		if lastN == -1 {
			return false
		}

		if stack[lastN] != opp[b] {
			return false
		}

		stack = stack[:lastN]
	}

	return len(stack) == 0
}
