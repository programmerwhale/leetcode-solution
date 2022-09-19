package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代
func isSymmetric(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	u, v := root.Left, root.Right
	q := []*TreeNode{}
	q = append(q, u)
	q = append(q, v)
	for len(q) > 1 {
		u, v = q[0], q[1]
		q = q[2:]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}
		if u != nil && v != nil && u.Val != v.Val {
			return false
		}

		// 不判断一边是空的情况，空也要放进去
		q = append(q, u.Left)
		q = append(q, v.Right)
		q = append(q, u.Right)
		q = append(q, v.Left)
	}
	return true
}

//递归
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compare(root.Left, root.Right)
}
func compare(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left != nil && right == nil {
		return false
	} else if left == nil && right != nil {
		return false
	} else if left != nil && right != nil && left.Val != right.Val {
		return false
	}

	return compare(left.Left, right.Right) && compare(left.Right, right.Left)
}
