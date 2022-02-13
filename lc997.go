package main

func findJudge(n int, trust [][]int) int {
	var cnt, st [10005]int

	for i := 0; i < len(trust); i++ {
		cnt[trust[i][1]]++
		st[trust[i][0]] = 1
	}
	flag := 0
	var ans = -1
	for i := 1; i <= n; i++ {
		if st[i] == 0 && cnt[i] == n-1 {
			flag++
			ans = i
		}
	}
	if flag == 1 {
		return ans
	} else {
		return -1
	}
}
