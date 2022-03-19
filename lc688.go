package main

func knightProbability(n int, k int, row int, column int) float64 {
	type pair struct {
		x, y, path int
	}

	ans := 0.0

	dx := []int{1, 2, -1, -2, 1, 2, -1, -2}
	dy := []int{2, 1, 2, 1, -2, -1, -2, -1}

	dp := make([][][]int, 30)

	for i := range dp {
		dp[i] = make([][]int, 30)
		for j := range dp[i] {
			dp[i][j] = make([]int, 105)
		}
	}

	q := make([]pair, 100005)
	hh, tt := 0, 0
	q[0] = pair{row, column, 0}
	dp[row][column][0] = 1
	for {
		t := q[hh]
		hh++
		if t.path > k {
			break
		}

		for i := 0; i < 8; i++ {
			a, b := t.x+dx[i], t.y+dy[i]
			if a < 0 || a >= n || b < 0 || b >= n {
				continue
			}
		}

		if hh > tt {
			break
		}
	}

	return ans
}
