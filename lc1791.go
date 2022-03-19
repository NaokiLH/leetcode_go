package main

func findCenter(edges [][]int) int {
	mp := make(map[int]int, 0)
	ans := 0

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	for i := range edges {
		a, b := edges[i][0], edges[i][1]
		mp[a]++
		mp[b]++
		ans = max(ans, max(mp[a], mp[b]))
	}
	return ans
}
