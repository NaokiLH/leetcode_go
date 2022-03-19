package main

func longestWord(words []string) string {
	max_s := ""
	mp := make(map[string]bool, 0)

	for _, i := range words {
		mp[i] = true
	}
	mp[""] = true
	check := func(a string) bool {
		for i := range a {
			if !mp[a[0:i]] {
				return false
			}
		}
		return true
	}
	for _, i := range words {
		if check(i) {
			if len(max_s) < len(i) {
				max_s = i
			} else if len(max_s) == len(i) && max_s > i {
				max_s = i
			}

		}
	}

	return max_s
}

func main() {
	words := []string{"w", "wo", "wor", "worl", "world"}
	longestWord(words)
}
