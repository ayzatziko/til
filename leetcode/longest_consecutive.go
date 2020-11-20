package leetcode

func LongestConsecutive(nums []int) int {
	var (
		hash = map[int]int{}
		seq  = func(n int) int {
			ok := true
			for i := 0; ; i++ {
				if n, ok = hash[n]; !ok {
					return i
				}
			}
		}
	)

	for _, n := range nums {
		hash[n] = n - 1
	}

	max := 0
	for _, n := range nums {
		if t := seq(n); max < t {
			max = t
		}
	}
	return max
}
