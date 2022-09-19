package main

import "fmt"

func main() {
	head := initNode(3)
	cur := head
	cur.Left = initNode(1)
	cur.Right = initNode(2)
	fmt.Println(preorderTraversal2(head))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func initNode(val int) *TreeNode {
	return &TreeNode{Val: val}
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

func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		res = append(res, node.Val)
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left != nil {
			q = append(q, node.Left)
		}
	}
	return res
}
func preorderTraversal3(root *TreeNode) []int {

	res := []int{}
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
	}
	return res
}
