package main

import (
	"sort"
)

func isNStraightHand(hand []int, groupSize int) bool {
	n := len(hand)
	if n%groupSize != 0 {
		return false
	}

	sort.Slice(hand, func(i, j int) bool {
		return hand[i] < hand[j]
	})
	mp := make(map[int]int, 0)
	for _, v := range hand {
		mp[v]++
	}

	for _, v := range hand {
		if mp[v] == 0 {
			continue
		}
		num := mp[v]

		for i := v; i < v+groupSize; i++ {
			if mp[i] >= num {
				mp[i] -= num
			} else {
				return false
			}
		}
	}

	return true
}
