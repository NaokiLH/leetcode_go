package main

import "strconv"

func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	} else {
		return a
	}
}

func simplifiedFractions(n int) []string {
	ans := make([]string, 0)

	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			if gcd(i, j) > 1 {
				continue
			}
			a := strconv.Itoa(i)
			b := strconv.Itoa(j)
			res := a + "/" + b
			ans = append(ans, res)
		}
	}

	return ans
}
