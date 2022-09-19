package main
/*
105. 从前序与中序遍历序列构造二叉树
给定两个整数数组 preorder 和 inorder ，其中 preorder 是二叉树的先序遍历， inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
示例 1:
输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出: [3,9,20,null,null,15,7]

示例 2:
输入: preorder = [-1], inorder = [-1]
输出: [-1]
*/
func main()  {
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 迭代
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder)==0||len(inorder)==0{
		return nil
	}
	return travelsal(preorder,inorder)
}
func travelsal(preorder,inorder []int)*TreeNode{
	if len(preorder)==0||len(inorder)==0{
		return nil
	}
	rootval:=preorder[0]
	root:=&TreeNode{Val:rootval}
	index:=-1
	for i,v:=range inorder{
		if v==rootval{
			index=i
			break
		}
	}
	leftpreorder:=preorder[1:index+1]
	rightpreorder:=preorder[index+1:]
	leftinorder:=inorder[:index]
	rightinorder:=inorder[index+1:]
	root.Left=travelsal(leftpreorder,leftinorder)
	root.Right=travelsal(rightpreorder,rightinorder)
	return root
}