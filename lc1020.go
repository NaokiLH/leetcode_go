package main

import "fmt"

type pair struct {
	x, y int
}

func numEnclaves(grid [][]int) int {
	ans, cnt := 0, 0
	n, m := len(grid), len(grid[0])
	dx := []int{0, -1, 0, 1}
	dy := []int{1, 0, -1, 0}

	q := make([]pair, 250005)
	st := make(map[pair]bool, 250005)
	hh, tt := 0, -1

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				if i == 0 || i == n-1 || j == 0 || j == m-1 {
					tt++
					q[tt] = pair{i, j}
					st[pair{i, j}] = true
				}
				cnt++
			}
		}
	}

	for {
		t := q[hh]
		hh++
		ans++

		for i := 0; i < 4; i++ {
			v := pair{t.x + dx[i], t.y + dy[i]}
			if v.x < 0 || v.y < 0 || v.x > n || v.y > m || grid[v.x][v.y] != 1 || st[v] {
				continue
			}
			tt++
			q[tt] = v
			st[v] = true
		}

		if hh > tt {
			break
		}
	}

	return cnt - ans
}
func main() {
	grid := [][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}
	fmt.Printf("%d\n", numEnclaves(grid))
}
