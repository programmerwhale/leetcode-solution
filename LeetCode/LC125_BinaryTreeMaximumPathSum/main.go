package main

import (
	"fmt"
	"math"
)

func main()  {
	head:=initTreeNode(-10)
	/*head.Left=initTreeNode(9)
	head.Right=initTreeNode(20)
	cur:=head.Right
	cur.Left=initTreeNode(15)
	cur.Right=initTreeNode(7)*/
	fmt.Println(maxPathSum(head))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func initTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

// 时间空间复杂度都是0(N)
func maxPathSum(root *TreeNode) (maxx int) {
	// 初始化
	maxx=math.MinInt64
	var findsum func(*TreeNode)int
	findsum=func(root *TreeNode)int  {
		if root==nil{
			return 0
		}
		leftsum:=max(findsum(root.Left),0)
		rightsum:=max(findsum(root.Right),0)
		// 记录最大值但返回当前结点的最大贡献值
		maxx=max(maxx,root.Val+leftsum+rightsum)
		return root.Val+max(leftsum,rightsum)
	}
	findsum(root)
	return maxx
}

func max(a,b int) int {
	if a>b{
		return a
	}
	return b
}