package main

import "math"

func main() {

}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 递归法
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = min(minDepth(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(minDepth(root.Right), minD)
	}
	return minD + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 层序遍历
func minDepth2(root *TreeNode) (res int) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		res++
		p := []*TreeNode{}
		for len(q) > 0 {
			node := q[0]
			q = q[1:]
			if node.Left == nil && node.Right == nil {
				return
			}
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return -1
}
