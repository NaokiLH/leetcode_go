package main

func catMouseGame(graph [][]int) int {
	n := len(graph)
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n*2)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	return 0
}
