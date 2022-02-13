package main

// import "sort"

// type pair struct {
// 	a, b int
// }

// func containsNearbyDuplicate(nums []int, k int) bool {
// 	list := make([]pair, 0)
// 	for i := range nums {
// 		list = append(list, pair{nums[i], i})
// 	}

// 	sort.Slice(list, func(i, j int) bool {
// 		if list[i].a < list[j].a {
// 			return true
// 		} else if list[i].a == list[j].a {
// 			return list[i].b < list[j].b
// 		} else {
// 			return false
// 		}
// 	})

// 	for i := range list {
// 		if i == 0 {
// 			continue
// 		}
// 		if list[i].a == list[i-1].a && list[i].b-list[i-1].b <= k {
// 			return true
// 		}
// 	}
// 	return false
// }
