package main

import (
	"sort"
)

func search(s string, mp map[string]bool) bool {

	for i := 1; i <= len(s); i++ {
		if mp[s[:i]] {
			mp[s] = search(s[i:], mp)
		}
	}
	return mp[s]
}

func findAllConcatenatedWordsInADict(words []string) []string {
	ans := make([]string, 0)

	mp := make(map[string]bool, 0)

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	for _, v := range words {
		mp[v] = true
	}

	for _, v := range words {
		if search(v, mp) {
			ans = append(ans, v)
		}
	}

	return ans
}
