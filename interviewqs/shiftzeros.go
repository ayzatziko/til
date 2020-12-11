package interviewqs

// given array, needed to move all zeroes to tail with keeping original ordering of non zero values
// memory should be O(1)
// [1,2,2,0,0,7,43,0,2,0,6] => [1,2,2,7,43,2,6,0,0,0,0]
// O(n**2)
func shiftzeros(arr []int) {
	for {
		zero, nonzero := -1, -1
		for i, v := range arr {
			if v == 0 {
				zero = i
				break
			}
		}

		if zero == -1 {
			return
		}

		for i := zero; i < len(arr); i++ {
			if arr[i] != 0 {
				nonzero = i
				break
			}
		}

		if nonzero == -1 {
			return // done
		}

		arr[zero], arr[nonzero] = arr[nonzero], 0
	}
}

// O(n)
func shiftzerosOptimized(arr []int) {
	lzero, lnonzero := 0, 0
	for {
		zero := lzero + 1
		got := false
		for ; zero < len(arr); zero++ {
			if arr[zero] == 0 {
				got = true
				break
			}
		}

		if !got {
			return
		}

		nonzero := lnonzero + 1
		if zero > nonzero {
			nonzero = zero + 1
		}

		got = false
		for ; nonzero < len(arr); nonzero++ {
			if arr[nonzero] != 0 {
				got = true
				break
			}
		}

		if !got {
			return // done
		}

		arr[zero], arr[nonzero] = arr[nonzero], 0
		lzero, lnonzero = zero, nonzero
	}
}
