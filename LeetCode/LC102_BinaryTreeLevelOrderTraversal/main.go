package main

import (
	"fmt"
)

/*
102.Binary Tree Level Order Traversal二叉树的层序遍历
给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

示例 1：

输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]
示例 2：

输入：root = [1]
输出：[[1]]
示例 3：

输入：root = []
输出：[]

提示：

树中节点数目在范围 [0, 2000] 内
-1000 <= Node.val <= 1000
*/
func main() {
	root := &TreeNode{Val: 3, Left: &TreeNode{9, initTreeNode(10), initTreeNode(11)},
		Right: &TreeNode{20, initTreeNode(12), initTreeNode(17)}}
	fmt.Println(levelOrder(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func initTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

// bfs,双队列
// 时间复杂度o(n), 空间复杂度o(n)
func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	// 为了好找到位置加了level作为下标
	level := 0
	for len(q) > 0 {
		res = append(res, []int{})
		p := []*TreeNode{}
		for len(q) > 0 {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
			res[level] = append(res[level], node.Val)
		}
		level++
		q = p
	}
	return
}

func levelOrder2(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		ans := []int{}
		p := []*TreeNode{}
		for len(q) > 0 {
			node := q[0]
			ans = append(ans, node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
			q = q[1:]
		}
		if len(q) == 0 {
			res = append(res, ans)
		}
		q = p
	}
	return
}
