package main

import (
	"fmt"
)

/*
103.Binary Tree Zigzag Level Order Traversal二叉树的锯齿形层序遍历
给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

 

示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：[[3],[20,9],[15,7]]
示例 2：

输入：root = [1]
输出：[[1]]
示例 3：

输入：root = []
输出：[]
 

提示：

树中节点数目在范围 [0, 2000] 内
-100 <= Node.val <= 100
*/
func main() {
	root := &TreeNode{Val: 3, Left: &TreeNode{9, nil,nil},
		Right: &TreeNode{20, initTreeNode(15), initTreeNode(17)}}
	fmt.Println(zigzagLevelOrder(root))
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
func zigzagLevelOrder(root *TreeNode) (res [][]int) {
	if root==nil{
		return
	}
	q:=[]*TreeNode{root}
	level:=0
	for len(q)>0{
		res=append(res,[]int{})
		p:=[]*TreeNode{}
		for len(q)>0{
			node:=q[0]
			q=q[1:]
			res[level]=append(res[level],node.Val)
			if node.Left!=nil{
				p=append(p,node.Left)
			}
			if node.Right!=nil{
				p=append(p,node.Right)
			}
		}
		level++
		q=p
	}
	for i:=1;i<len(res);i=i+2{
		reverse(res[i])
	}

	return
}
func reverse(ans []int){
	for i:=0;i<len(ans)/2;i++{
		ans[i],ans[len(ans)-i-1]=ans[len(ans)-i-1],ans[i]
	}
}
