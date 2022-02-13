package main

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func increasingTriplet(nums []int) bool {
	m := 0
	q := make([]int, 5e5+5)
	q[0] = -2e9
	for i := range nums {
		l, r := 0, m
		for {
			mid := (l + r + 1) >> 1

			if q[mid] < nums[i] {
				l = mid
			} else {
				r = mid - 1
			}

			if l >= r {
				break
			}
		}
		m = max(m, r+1)
		q[r+1] = nums[i]
	}

	return m >= 3
}
