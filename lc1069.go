package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	st := make(map[*TreeNode]int, 0)
	q := make([]*TreeNode, 100005)
	hh, tt := 0, 0
	q[hh] = root
	st[root] = 1
	flag := true

	for {
		t := q[hh]
		hh++

		if st[t]%2 != t.Val%2 {
			flag = false
			break
		}

		lnode := t.Left
		rnode := t.Right
		if lnode != nil {
			tt++
			q[tt] = lnode
			st[lnode] = st[t] + 1
		}
		if rnode != nil {
			tt++
			q[tt] = rnode
			st[rnode] = st[t] + 1
		}

		if hh > tt {
			break
		}
	}

	return flag
}
