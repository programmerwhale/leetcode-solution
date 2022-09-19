package main
/*
106. 从中序与后序遍历序列构造二叉树
给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。

输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
输出：[3,9,20,null,null,15,7]

输入：inorder = [-1], postorder = [-1]
输出：[-1]
*/
func main(){

}

type TreeNode struct {
	    Val int
	    Left *TreeNode
	    Right *TreeNode
	 }
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder)==0||len(postorder)==0{
		return nil
	}
	return traversal(inorder,postorder)
}

func traversal(inorder []int, postorder []int) *TreeNode {
	if len(inorder)==0||len(postorder)==0{
		return nil
	}
	rootval:=postorder[len(postorder)-1]
	root:=&TreeNode{}
	root.Val=rootval
	if len(postorder)==1{
		return root
	}
	index:=0
	for i:=0;i<len(inorder);i++{
		if inorder[i]==root.Val{
			index=i
			break
		}
	}
	leftinorder:=inorder[:index]
	rightinorder:=inorder[index+1:]
	leftpostorder:=postorder[:len(leftinorder)]
	rightpostorder:=postorder[len(leftinorder):len(postorder)-1]
	root.Left=traversal(leftinorder,leftpostorder)
	root.Right=traversal(rightinorder,rightpostorder)
	return root
}