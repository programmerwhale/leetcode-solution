package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func initNode(val int) *TreeNode{
	return &TreeNode{Val: val}
}

func main()  {
	head:=initNode(5)
	cur:=head
	cur.Left=initNode(4)
	cur=head.Left
	cur.Left=initNode(11)
	cur=cur.Left
	cur.Left=initNode(7)
	cur.Right=initNode(2)
	head.Right=initNode(8)
	cur=head.Right
	cur.Left=initNode(13)
	cur.Right=initNode(4)
	cur=cur.Right
	cur.Left=initNode(5)
	cur.Right=initNode(1)

	fmt.Println(preorderTraversal(head))
	fmt.Println(hasPathSum(head,22))
}

func preorderTraversal(root *TreeNode)[]int{
	res:=[]int{}
	dfs(&res,root)
	return res
}

func dfs(res *[]int, root *TreeNode)  {
	if root==nil{
		return
	}
	*res=append(*res,root.Val)
	dfs(res,root.Left)
	dfs(res,root.Right)
}

func hasPathSum(root *TreeNode, targetSum int) (res [][]int) {
	if root==nil{
		return res
	}
	path:=[]int{}
	var dfs func(*TreeNode,int)
	dfs = func(root *TreeNode, targetSum int) {
		if root==nil{
			return
		}
		path=append(path,root.Val)
		targetSum-=root.Val
		if targetSum==0&&root.Left==nil&&root.Right==nil{
			// 不能直接append，需要copy
			// path共享，遍历时会被修改
			temp:=make([]int,len(path))
			copy(temp,path)
			res=append(res,temp)
		}
		dfs(root.Left,targetSum)
		dfs(root.Right,targetSum)
		targetSum+=root.Val
		path=path[:len(path)-1]
	}
	dfs(root,targetSum)
	return res
}
