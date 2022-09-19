package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func initNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func main() {
	head := initNode(5)
	cur := head
	cur.Left = initNode(4)
	cur = head.Left
	cur.Left = initNode(11)
	cur = cur.Left
	cur.Left = initNode(7)
	cur.Right = initNode(2)
	head.Right = initNode(8)
	cur = head.Right
	cur.Left = initNode(13)
	cur.Right = initNode(4)
	cur = cur.Right
	cur.Right = initNode(1)

	fmt.Println(preorderTraversal(head))
	fmt.Println(hasPathSum(head, 22))
}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	dfs(&res, root)
	return res
}

func dfs(res *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	dfs(res, root.Left)
	dfs(res, root.Right)
}

// 更精简的递归
func hasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 左右节点皆为空，即这里为叶子结点时直接判断值
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// 初版递归
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return traversal(root, targetSum-root.Val)
}

func traversal(root *TreeNode, cnt int) bool {
	if root.Left == nil && root.Right == nil && cnt == 0 {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return false
	}
	if root.Left != nil {
		if traversal(root.Left, cnt-root.Left.Val) {
			return true
		}
	}
	if root.Right != nil {
		if traversal(root.Right, cnt-root.Right.Val) {
			return true
		}
	}
	return false
}
