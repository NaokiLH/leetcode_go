package main

// import (
// 	"sort"
// )

// type Pair struct {
// 	l int
// 	r int
// }

// func check(mid int, houses []int, heaters []int) bool {
// 	h := make([]Pair, 0)
// 	for i, v := range heaters {
// 		if i == 0 {
// 			h = append(h, Pair{v - mid, v + mid})
// 		} else {
// 			last := &h[len(h)-1]
// 			if v-mid <= last.r {
// 				last.r = v + mid
// 			} else {
// 				h = append(h, Pair{v - mid, v + mid})
// 			}
// 		}
// 	}

// 	for _, v := range houses {
// 		res := sort.Search(len(h), func(i int) bool { return h[i].r >= v })
// 		if res < 0 || res >= len(h) {
// 			return false
// 		}
// 		if h[res].l > v {
// 			return false
// 		}
// 	}

// 	return true
// }
// func findRadius(houses []int, heaters []int) int {
// 	sort.Slice(houses, func(i, j int) bool { return houses[i] < houses[j] })
// 	sort.Slice(heaters, func(i, j int) bool {
// 		return heaters[i] < heaters[j]
// 	})
// 	var ans int = 0
// 	l, r := 0, 1000000000
// 	for {
// 		mid := (l + r) >> 1
// 		if check(mid, houses, heaters) {
// 			ans = mid
// 			r = mid
// 		} else {
// 			l = mid + 1
// 		}
// 		if l >= r {
// 			break
// 		}
// 	}

// 	return ans
// }
