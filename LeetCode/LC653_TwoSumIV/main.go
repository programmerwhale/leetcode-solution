package main

func main() {

}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 深搜
func findTarget(root *TreeNode, k int) bool {
	mp := map[int]int{}
	res := 0
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if _, ok := mp[k-root.Val]; ok {
			res++
		}
		mp[root.Val] = 1
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return res > 0
}
func findTarget2(root *TreeNode, k int) bool {
	mp := map[int]bool{}
	var dfs func(*TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		if _, ok := mp[k-root.Val]; ok {
			return true
		}
		mp[root.Val] = true
		return dfs(root.Left) || dfs(root.Right)
	}
	return dfs(root)

}
