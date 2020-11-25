package crackint

// Recursion and Dynamic Programming

// Problem 8.1 Triple Step
func waysToRunUp(stairsN int) int {
	if stairsN <= 0 {
		return 0
	} else if stairsN == 1 {
		return 1
	} else if stairsN == 2 {
		return 2
	} else if stairsN == 3 {
		return 4
	}

	n := waysToRunUp(stairsN - 1)
	n += waysToRunUp(stairsN - 2)
	n += waysToRunUp(stairsN - 3)
	return n
}

func waysToRunUpMemo(stairsN int) int {
	memoN := 4
	if stairsN+1 > memoN {
		memoN = stairsN + 1
	}
	memo := make([]int, memoN)
	memo[1] = 1
	memo[2] = 2
	memo[3] = 4
	return waysToRunUpMemoRecur(stairsN, memo)
}

func waysToRunUpMemoRecur(stairsN int, memo []int) int {
	if stairsN <= 0 {
		return 0
	}

	if n := memo[stairsN]; n > 0 {
		return n
	}

	n := waysToRunUpMemoRecur(stairsN-1, memo)
	n += waysToRunUpMemoRecur(stairsN-2, memo)
	n += waysToRunUpMemoRecur(stairsN-3, memo)
	memo[stairsN] = n
	return n
}
