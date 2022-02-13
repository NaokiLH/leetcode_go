package main

import "math"

func numFriendRequests(ages []int) int {
	pre := make([]int, 123)
	ans := 0

	for _, v := range ages {
		pre[v]++
	}
	for i := range pre {
		if i != 0 {
			pre[i] += pre[i-1]
		}
	}

	for _, v := range ages {
		r := v
		l := int(math.Floor(0.5*float64(v) + 7))
		if l > r {
			continue
		}

		ans += pre[r] - pre[l]
		if v <= r && v > l {
			ans--
		}
	}

	return ans
}
