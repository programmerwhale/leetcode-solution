package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深搜递归，中序遍历
func findMode(root *TreeNode) (res []int) {
	maxTimes := 0
	times := 0
	val := math.MinInt64
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		//处理
		if root.Val == val {
			times++
		} else if root.Val > val {
			times, val = 1, root.Val
		}
		if times == maxTimes {
			res = append(res, val)
		} else if times > maxTimes {
			maxTimes = times
			res = []int{val}
		}

		dfs(root.Right)
	}
	dfs(root)
	return
}
