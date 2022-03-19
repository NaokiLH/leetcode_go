package main

func luckyNumbers(matrix [][]int) []int {

	ans := make([]int, 0)
	for i := range matrix {
		res := 0x3f3f3f3f
		resj := 0
		for j := range matrix[i] {
			if res > matrix[i][j] {
				res = matrix[i][j]
				resj = j
			}
		}
		flag := true
		for j := range matrix {
			if matrix[j][resj] > res {
				flag = false
			}
		}

		if flag {
			ans = append(ans, res)
		}
	}
	return ans
}
