package main

// import "fmt"

// type pair struct {
// 	x, y int
// }

// func highestPeak(isWater [][]int) [][]int {
// 	n := len(isWater)
// 	m := len(isWater[0])
// 	dx := []int{1, 0, -1, 0}
// 	dy := []int{0, 1, 0, -1}
// 	q := make([]pair, 1e6+5)
// 	ans := make([][]int, len(isWater))
// 	for i := range ans {
// 		ans[i] = make([]int, len(isWater[0]))
// 	}

// 	hh, tt := 0, -1

// 	for i := range isWater {
// 		for j := range isWater[i] {
// 			if isWater[i][j] == 1 {
// 				tt++
// 				q[tt] = pair{i, j}
// 			}
// 		}
// 	}

// 	for {
// 		x, y := q[hh].x, q[hh].y
// 		hh++

// 		fmt.Printf("%d %d\n", x, y)

// 		for i := 0; i < 4; i++ {
// 			a, b := x+dx[i], y+dy[i]
// 			if a < 0 || b < 0 || a >= n || b >= m || isWater[a][b] == 1 {
// 				continue
// 			}
// 			tt++
// 			q[tt] = pair{a, b}
// 			isWater[a][b] = 1
// 			ans[a][b] = ans[x][y] + 1
// 		}

// 		if hh > tt {
// 			break
// 		}
// 	}

// 	return ans
// }
// func main() {
// 	t := [][]int{{0}, {0}, {0}, {0}, {1}, {0}, {0}, {1}, {1}}
// 	fmt.Printf("%d\n", highestPeak(t))
// }
