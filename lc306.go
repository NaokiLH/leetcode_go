package main

import (
	"strconv"
)

func check(a1, a2 int, num string) bool {
	if len(num) == 0 {
		return true
	}
	a3 := a1 + a2
	str3 := strconv.Itoa(a3)
	if len(str3) > len(num) {
		return false
	}
	if str3 == num[0:len(str3)] {
		return check(a2, a3, num[len(str3):])
	}
	return false
}

func isAdditiveNumber(num string) bool {
	n := len(num)

	for i := 1; i <= n; i++ {
		a1, err := strconv.Atoi(num[0:i])
		if err != nil {
			return false
		}
		for j := 1; j <= n; j++ {
			if i+j >= n {
				continue
			}
			a2, err := strconv.Atoi(num[0+i : 0+i+j])
			if err != nil {
				return false
			}
			if check(a1, a2, num[0+i+j:]) {
				return true
			}
			if a2 == 0 {
				break
			}
		}
		if a1 == 0 {
			return false
		}
	}
	return false
}
