package main

// func checkWays(pairs [][]int) int {

// 	st := make(map[int]bool, 0)
// 	g := make([][]int, 505)

// 	for i := range g {
// 		g[i] = make([]int, 0)
// 	}

// 	for i := range pairs {
// 		a, b := pairs[i][0], pairs[i][1]
// 		st[a] = true
// 		st[b] = true
// 		g[a] = append(g[a], b)
// 		g[b] = append(g[b], a)
// 	}
// 	n := len(st)

// 	flag1, flag2 := false, false

// 	for i := range g {
// 		if len(g[i]) == n-1 {
// 			flag1 = true
// 		}
// 	}
// 	if !flag1 {
// 		return 0
// 	}

// 	return 0
// }
// func main() {
// 	pairs := [][]int{{1, 2}, {2, 3}}
// 	checkWays(pairs)
// }
