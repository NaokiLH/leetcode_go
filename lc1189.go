package main

func maxNumberOfBalloons(text string) int {
	b, a, l, o, n := 0, 0, 0, 0, 0

	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}

	for _, v := range text {
		switch v {
		case 'a':
			a++
		case 'b':
			b++
		case 'l':
			l++
		case 'o':
			o++
		case 'n':
			n++
		}
	}

	return min(n, min(min(a, b), min(l/2, o/2)))
}
