package main

func main()  {
	
}
type TreeNode struct {
	Val int
	Left,Right *TreeNode
}

// 递归
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root==nil{
		return nil
	}
	if root.Val<low{
		right:=trimBST(root.Right,low,high)
		return right
	}
	if root.Val>high{
		left:=trimBST(root.Left,low,high)
		return left
	}
	root.Left=trimBST(root.Left,low,high)
	root.Right=trimBST(root.Right,low,high)
	return root
}


// 迭代
func trimBST2(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	// 处理 root，让 root 移动到[low, high] 范围内，注意是左闭右闭
	for root != nil && (root.Val<low||root.Val>high){
		if root.Val < low{
			root = root.Right
		}else{
			root = root.Left
		}
	}
	cur := root
	// 此时 root 已经在[low, high] 范围内，处理左孩子元素小于 low 的情况（左节点是一定小于 root.Val，因此天然小于 high）
	for cur != nil{
		// 里面也要用for循环，因为可能整个子树下面都不符合条件，不能只去掉一个
		for cur.Left!=nil && cur.Left.Val < low{
			cur.Left = cur.Left.Right
		}
		cur = cur.Left
	}
	// 重置一下！！！
	cur = root
	// 此时 root 已经在[low, high] 范围内，处理右孩子大于 high 的情况
	for cur != nil{
		for cur.Right!=nil && cur.Right.Val > high{
			cur.Right = cur.Right.Left
		}
		cur = cur.Right
	}
	return root
}