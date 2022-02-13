package main

import "fmt"

func repeatedStringMatch(a string, b string) int {

	res := make([]byte, 0)
	ans := 0
	i := 0
	for i = 0; i < len(b)-len(a); i++ {
		fmt.Printf("%s\n", b[i:i+len(a)])
		if b[i:i+len(a)] != a {
			res = append(res, b[i])
			i += len(a)
		} else {
			ans++
		}
	}
	for ; i < len(b); i++ {
		res = append(res, b[i])
	}
	fmt.Printf("%s\n", res)
	tmp := a + a
	idx := 0
	for i := 0; i < len(tmp)-len(res); i++ {
		if tmp[i:i+len(res)] == string(res) {
			idx = i
			break
		}
	}
	if idx >= len(a) {
		ans++
	}

	return ans
}
