package main

func main()  {
	node:=&TreeNode{Val: 4}
	root:=node
	root.Left=initTreeNode(1)
	root.Right=initTreeNode(6)
	root=root.Right
	root.Right=initTreeNode(7)
	convertBST2(node)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func initTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}
// 深搜递归
func convertBST(root *TreeNode) *TreeNode {
	sum:=0
	var dfs func(*TreeNode)
	dfs=func(root *TreeNode){
		if root==nil{
			return
		}
		dfs(root.Right)
		root.Val+=sum
		sum=root.Val
		dfs(root.Left)
	}
	dfs(root)
	return root
}

// 深搜迭代
func convertBST2(node *TreeNode) *TreeNode {
	sum:=0
	stack:=[]*TreeNode{}
	root:=node
	for root!=nil||len(stack)>0{
		for root!=nil{
			stack=append(stack,root)
			root=root.Right
		}
		root=stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		root.Val+=sum
		sum=root.Val
		root=root.Left
	}
	return root
}

// todo
// Morris
func getSuceessor(node *TreeNode)*TreeNode{
	succ:=node.Right
	for succ.Left!=nil&&succ.Left!=node{
		succ=succ.Left
	}
	return succ
}
func convertBST3(root *TreeNode)*TreeNode{
	sum:=0
	node:=root
	for node!=nil{
		if node.Right==nil{
			sum+=node.Val
			node.Val=sum
			node=node.Left
		}else{
			succ:=getSuceessor(node)
			if succ.Left==nil{
				succ.Left=node
				node=node.Right
			}else {
				succ.Left=nil
				sum+=node.Val
				node.Val=sum
				node=node.Left
			}
		}
	}
	return root
}