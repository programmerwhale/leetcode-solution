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
	head := initNode(1)
	cur := head
	cur.Left = initNode(2)
	cur = head.Left
	cur.Left = initNode(3)
	cur.Right = initNode(4)
	head.Right = initNode(5)
	cur = head.Right
	cur.Left = initNode(6)

	fmt.Println(preorderTraversal(head))
	flatten(head)
	fmt.Println(preorderTraversal(head))
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
func flatten(root *TreeNode) {
	arr := []*TreeNode{}
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		arr = append(arr, root)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)

	// 后面这个负值，不用新定义变量
	for i := 1; i < len(arr); i++ {
		prev, curr := arr[i-1], arr[i]
		prev.Left, prev.Right = nil, curr
	}
}

// 原地交换
func flatten2(root *TreeNode)  {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			predecessor := next
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			predecessor.Right = curr.Right
			curr.Left, curr.Right = nil, next
		}
		curr = curr.Right
	}
}