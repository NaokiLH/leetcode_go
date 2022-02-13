package main

func modifyString(s string) string {

	for i := range s {
		if s[i] == '?' {
			l, r := byte('.'), byte('.')
			if i-1 >= 0 {
				l = s[i-1]
			}
			if i+1 <= len(s)-1 {
				r = s[i+1]
			}
			for j := byte('a'); j <= byte('z'); j++ {
				if j != l && j != r {
					break
				}
			}
		}
	}
	return s
}
