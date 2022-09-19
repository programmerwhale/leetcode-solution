package main

import "fmt"

func main()  {
	head:=initTreeNode(1)
	head.Right=initTreeNode(1)
	/*head.Left=initTreeNode(9)
	head.Right=initTreeNode(20)
	cur:=head.Right
	cur.Left=initTreeNode(15)
	cur.Right=initTreeNode(7)*/
	fmt.Println(isSubtree(head,initTreeNode(1)))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func initTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}
// 方法一：递归
// 时间复杂度O(MN)
// 空间复杂度O(max(m,n))
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root==nil{
		return false
	}
	// 主函数，先看两个root构成的数是否一致，或者往两边找
	return check(root,subRoot)||isSubtree(root.Left,subRoot)||isSubtree(root.Right,subRoot)
}
// 子函数：看两棵树是否为空；root是否一致，一致的话递归往下找
func check(root *TreeNode, subRoot *TreeNode) bool {
	// 都为空 true
	if root==nil&&subRoot==nil{
		return true
	}
	// 有的空 有的不空
	if root==nil||subRoot==nil{
		return false
	}
	// 头相等 ，继续找
	if root.Val==subRoot.Val{
		return check(root.Left,subRoot.Left)&&check(root.Right,subRoot.Right)
	}
	return false
}

// todo
// 方法二：kmp
// 时间复杂度O(MN)
// 空间复杂度O(max(m,n))
func isSubtree2(root *TreeNode, subRoot *TreeNode) bool {
	return false
}

// todo
// 方法三：前序遍历
// 时间复杂度O(M+N)
// 空间复杂度O(1)
func isSubtree3(root *TreeNode, subRoot *TreeNode) bool {
	return false
}
func preorderTraserval(root TreeNode){

}