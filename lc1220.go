package main

const MOD int = 1e9 + 7

func countVowelPermutation(n int) int {
	//a e i o u
	//0 1 2 3 4
	var dp [2e4 + 5][6]int
	for i := range dp[0] {
		dp[0][i] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 0; j < 5; j++ {
			switch {
			case j == 0:
				dp[i][j] = (dp[i][j] + dp[i][1]) % MOD
				dp[i][j] = (dp[i][j] + dp[i][2]) % MOD
				dp[i][j] = (dp[i][j] + dp[i][4]) % MOD
			case j == 1:
				dp[i][j] = (dp[i][j] + dp[i][0]) % MOD
				dp[i][j] = (dp[i][j] + dp[i][2]) % MOD
			case j == 2:
				dp[i][j] = (dp[i][j] + dp[i][1]) % MOD
				dp[i][j] = (dp[i][j] + dp[i][3]) % MOD
			case j == 3:
				dp[i][j] = (dp[i][j] + dp[i][2]) % MOD
			case j == 4:
				dp[i][j] = (dp[i][j] + dp[i][2]) % MOD
				dp[i][j] = (dp[i][j] + dp[i][3]) % MOD
			}
		}
	}

	ans := 0
	for i := range dp[n] {
		ans = (ans + dp[n][i]) % MOD
	}

	return ans
}
