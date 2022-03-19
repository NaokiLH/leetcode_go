package main

import "strconv"

func dfs(u *TreeNode, ans *string) {
	*ans += strconv.Itoa(u.Val)
	// if u.Right == nil && u.Left == nil{
	//     return
	// }
	if u.Left != nil {
		*ans += "("
		dfs(u.Left, ans)
		*ans += ")"
	}
	if u.Left == nil && u.Right != nil {
		*ans += "()"
	}
	if u.Right != nil {
		*ans += "("
		dfs(u.Right, ans)
		*ans += ")"
	}
}

func tree2str(root *TreeNode) string {
	ans := ""
	dfs(root, &ans)
	return ans
}
