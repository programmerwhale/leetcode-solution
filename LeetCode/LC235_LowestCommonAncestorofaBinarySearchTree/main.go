package main

func main()  {

}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 迭代
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	for root!=nil{
		if root.Val>p.Val&&root.Val>q.Val{
			root=root.Left
		}else if root.Val<p.Val&&root.Val<q.Val{
			root=root.Right
		}else{
			return root
		}
	}
	return root
}

// 递归
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root==nil{
		return nil
	}
	if root.Val>p.Val&&root.Val>q.Val{
		return lowestCommonAncestor(root.Left,p,q)
	}else if root.Val<p.Val&&root.Val<q.Val{
		return lowestCommonAncestor(root.Right,p,q)
	}else{
		return root
	}
}