package main

func main()  {
	
}
type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

// 递归
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root==nil{
		root=&TreeNode{Val:val}
	}
	if val<root.Val{
		root.Left=insertIntoBST(root.Left,val)
	}
	if val>root.Val{
		root.Right=insertIntoBST(root.Right,val)
	}
	return root
}

// 迭代
func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val:val}
	}
	node:=root
	var pre *TreeNode
	for node!=nil{
		pre=node
		// 不能两个if 因为cur值被修改了
		if val<node.Val{
			node=node.Left
		}else if val>node.Val{
			node=node.Right
		}
	}
	if val > pre.Val {
		pre.Right = &TreeNode{Val: val}
	} else {
		pre.Left = &TreeNode{Val: val}
	}
	return root
}