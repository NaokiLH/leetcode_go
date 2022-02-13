package main

import "sort"

func minimumDifference(nums []int, k int) int {
	ans := 0
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}

	for i := 0; i < len(nums)-k; i++ {
		ans = min(ans, nums[i+k-1]-nums[i])
	}

	return ans
}
