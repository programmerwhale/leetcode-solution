package main

func main() {

}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 方法一：层序遍历
func countNodes(root *TreeNode) (res int) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		res++
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	return
}

// 方法二：深搜递归法
func countNodes2(root *TreeNode) (res int) {
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		res++
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return
}

// 方法三：深搜迭代法
func countNodes3(root *TreeNode) (res int) {
	q := []*TreeNode{}
	for root != nil || len(q) > 0 {
		for root != nil {
			q = append(q, root)
			root = root.Left
			res++
		}
		root = q[0]
		q = q[1:]
		res++
		root = root.Right
	}
	return
}

//方法四：递归
func countNodes4(root *TreeNode) (res int) {
	if root == nil {
		return
	}
	return countNodes4(root.Left) + countNodes4(root.Right) + 1
}
