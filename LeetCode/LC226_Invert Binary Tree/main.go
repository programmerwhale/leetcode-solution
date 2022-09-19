package main

func main() {

}

/*
给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。

输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]

输入：root = [2,1,3]
输出：[2,3,1]

输入：root = []
输出：[]
*/
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 方法一：深搜迭代 前序或后序
func invertTree(node *TreeNode) *TreeNode {
	root := node
	q := []*TreeNode{}
	for root != nil || len(q) > 0 {
		for root != nil {
			q = append(q, root)
			root.Left, root.Right = root.Right, root.Left
			root = root.Left
		}
		root = q[0]
		q = q[1:]
		root = root.Right
	}
	return node
}

// 方法二：深搜递归
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

// 方法三：广搜
func invertTree3(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	return root
}
