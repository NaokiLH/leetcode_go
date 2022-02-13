package main

import "fmt"

func slowestKey(releaseTimes []int, keysPressed string) byte {
	ans := 0
	res := 0
	last := 0
	for i := range releaseTimes {
		fmt.Printf("%d %d\n", releaseTimes[i], last)
		if releaseTimes[i]-last >= res {
			if releaseTimes[i]-last > res || keysPressed[i] > keysPressed[ans] {
				ans = i
			}
			res = releaseTimes[i] - last
			last = releaseTimes[i]
		}
	}

	return keysPressed[ans]
}
