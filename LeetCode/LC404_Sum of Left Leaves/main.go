package main

func main() {

}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 递归法
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	value := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		value += root.Left.Val
	}
	return value + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

// 迭代法
func sumOfLeftLeaves2(root *TreeNode) (sum int) {
	if root == nil {
		return 0
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			sum += node.Left.Val
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return
}
