package main

import "sort"

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {

	ans := make([][]int, 0)
	n := len(nums1)
	m := len(nums2)
	for i := 0; i < min(n, k); i++ {
		for j := 0; j < min(m, k); j++ {
			res := []int{nums1[i], nums2[j]}
			ans = append(ans, res)
		}
	}

	sort.Slice(ans, func(i, j int) bool {
		sum1 := ans[i][0] + ans[i][1]
		sum2 := ans[j][0] + ans[j][1]
		return sum1 < sum2
	})

	ans = ans[0:k]

	return ans
}
