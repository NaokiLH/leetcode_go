package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	list := strings.Split(path, "/")

	for _, v := range list {
		fmt.Printf("%s\n", v)
	}

	q := make([]string, 3000)
	hh, tt := 0, -1

	for _, p := range list {
		if p == ".." {
			if hh <= tt {
				tt--
			}
			continue
		}
		if p == "." {
			continue
		}
		if p == "" {
			continue
		}
		tt++
		q[tt] = p
	}

	if hh > tt {
		return "/"
	} else {
		ans := ""
		for i := hh; i <= tt; i++ {
			ans += "/"
			ans += q[i]
		}
		return ans
	}
}
