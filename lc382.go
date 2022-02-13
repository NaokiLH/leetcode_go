package main

import (
	"math"
	"sort"
)

func parse(s string) int {
	h := (s[0]-'0')*10 + (s[1] - '0')
	m := (s[3]-'0')*10 + (s[4] - '0')
	x := int(h)*60 + int(m)
	return x
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func findMinDifference(timePoints []string) int {
	list := make([]int, 2e4+5)

	for _, v := range timePoints {
		list = append(list, parse(v))
	}
	sort.Ints(list)
	ans := 0x3f3f3f3f
	for i := range list {
		if i == 0 {
			continue
		}
		ans = min(ans, int(math.Abs(float64(list[i]-list[i-1]))))
	}

	return ans
}
