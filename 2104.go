package main

func subArrayRanges(nums []int) int64 {
	pre := make([]int, 10005)
	for i := range nums {
		if i == 0 {
			pre[i] = nums[i]
			continue
		}
		pre[i] = nums[i] + pre[i-1]
	}
	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}
	max := func(a, b int64) int64 {
		if a > b {
			return a
		} else {
			return b
		}
	}
	minlist := make([]int, 10005)

	for i := range pre {
		if i == 0 {
			minlist[i] = 0
			continue
		}
		minlist[i] = min(pre[i], minlist[i-1])
	}

	ans := int64(0)

	for i := range pre {
		tmp := int64(pre[i]) - int64(minlist[i])
		ans = max(ans, int64(tmp))
	}

	return int64(ans)
}
