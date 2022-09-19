package main

import "fmt"

/*
104.Maximum Depth of Binary Tree二叉树的最大深度
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
*/
func main() {
	root := &TreeNode{Val: 3, Left: &TreeNode{9, nil, nil},
		Right: &TreeNode{20, initTreeNode(12), initTreeNode(17)}}
	fmt.Println(maxDepth2(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func initTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

// dfs
// 时间复杂度O(N),空间复杂度O(height)，取决于树的高度
func maxDepth(root *TreeNode) (depth int) {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// bfs
// 时间复杂度O(N),空间复杂度O(n)
func maxDepth2(root *TreeNode) (depth int) {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		p := []*TreeNode{}
		depth++
		for len(q) > 0 {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return depth
}
