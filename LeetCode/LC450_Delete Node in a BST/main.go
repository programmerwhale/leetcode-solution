package main
type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}
func main()  {
	
}
// 递归法
// 五种情况
func deleteNode(root *TreeNode, key int) *TreeNode {
	// 情况一：根节点为空
	if root==nil{
		return root
	}
	if root.Val==key{
		// 第二种情况：左右孩子都为空（叶子节点），直接删除节点， 返回NULL为根节点
		//(第二种情况可以合并进第三种里面)
/*		if root.Left==nil&&root.Right==nil{
			return nil
		}*/
		// 第三种情况：其左孩子为空，右孩子不为空，删除节点，右孩子补位 ，返回右孩子为根节点
		if root.Left==nil{
			return root.Right
		}
		// 第四种情况：其右孩子为空，左孩子不为空，删除节点，左孩子补位，返回左孩子为根节点
		if root.Right==nil{
			return root.Left
		}
		// 第五种情况：左右孩子节点都不为空，则将删除节点的左子树放到删除节点的右子树的最左面节点的左孩子的位置
		// 并返回删除节点右孩子为新的根节点。
		cur:=root.Right
		for cur!=nil&&cur.Left!=nil{
			cur=cur.Left
		}
		cur.Left=root.Left
		// 不能直接用root=cur，因为cur不知道是在第几层
		root=root.Right
	}else if root.Val>key{
		root.Left=deleteNode(root.Left,key)
	}else{
		root.Right=deleteNode(root.Right,key)
	}
	return root
}

// 通用递归
func deleteNode2(root *TreeNode, key int) *TreeNode {
	if root==nil{
		return root
	}
	if root.Val==key{
		if root.Right==nil{
			return root.Left
		}
		cur:=root.Right
		for cur!=nil&&cur.Left!=nil{
			cur=cur.Left
		}
		root.Val,cur.Val=cur.Val,root.Val
	}
	root.Left=deleteNode2(root.Left,key)
	root.Right=deleteNode2(root.Right,key)
	return root
}

// 迭代法
func deleteNode3(root *TreeNode, key int) *TreeNode {
	if root==nil{
		return root
	}
	cur:=root
	// 前驱节点
	var pre *TreeNode
	// 为了找到要删除的节点
	for cur!=nil{
		if cur.Val==key{
			break
		}
		pre=cur
		if cur.Val>key{
			cur=cur.Left
		}else{
			cur=cur.Right
		}
	}
	if pre==nil{
		return deleteOneNode(cur)
	}
	if pre.Left!=nil&&pre.Left.Val==key{
		pre.Left=deleteOneNode(cur)
	}
	if pre.Right!=nil&&pre.Right.Val==key{
		pre.Right=deleteOneNode(cur)
	}
	return root
}

// 辅助函数，删除
func deleteOneNode(root *TreeNode)*TreeNode{
	if root.Left==nil&&root.Right==nil{
		return nil
	}
	if root.Left==nil{
		return root.Right
	}
	if root.Right==nil{
		return root.Left
	}
	cur:=root.Right
	for cur!=nil&&cur.Left!=nil{
		cur=cur.Left
	}
	cur.Left=root.Left
	return root.Right
}
