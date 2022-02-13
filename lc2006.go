package main

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
func countKDifference(nums []int, k int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if abs(nums[i]-nums[j]) == k {
				ans++
			}
		}
	}
	return ans
}
