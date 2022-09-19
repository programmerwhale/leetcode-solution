package main

func main()  {
	
}

type Node struct {
	Val int
    Children []*Node
}

// 递归
func preorder(root *Node) (res []int) {
	var dfs func(*Node)
	dfs=func(root *Node){
		if root==nil{
			return
		}
		res=append(res,root.Val)
		for _,node:=range root.Children{
			dfs(node)
		}
	}
	dfs(root)
	return
}

// 迭代
func preorder2(root *Node) (res []int) {
	if root==nil{
		return res
	}
	stack:=[]*Node{root}
	for len(stack)>0{
		node:=stack[len(stack)-1]
		stack=stack[:len(stack)-1]
		res=append(res,node.Val)
		// 加入栈的时候需要倒序往里面放
		for i:=len(node.Children)-1;i>=0;i--{
			stack=append(stack,node.Children[i])
		}
	}
	return
}