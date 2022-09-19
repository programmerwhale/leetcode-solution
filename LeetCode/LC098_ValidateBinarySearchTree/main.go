package main

import (
	"fmt"
	"math"
)

/*
98. Validate Binary Search Tree验证二叉搜索树
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。

示例 1：

输入：root = [2,1,3]
输出：true
示例 2：

输入：root = [5,1,4,null,null,3,6]
输出：false
解释：根节点的值是 5 ，但是右子节点的值是 4 。

提示：

树中节点数目范围在[1, 104] 内
-231 <= Node.val <= 231 - 1
*/
func main() {
	root := &TreeNode{Val: 5, Left: &TreeNode{4, initTreeNode(1), initTreeNode(7)}, Right: initTreeNode(8)}
	fmt.Println(isValidBST2(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func initTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

//  递归
//  根节点是左子树的上限，是右子树的下限
//  时间复杂度O(n)，空间复杂度O(n)
func isValidBST(root *TreeNode) bool {
	return recurse(root, math.MinInt64, math.MaxInt64)
}

func recurse(root *TreeNode, lower, higher int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= higher {
		return false
	}
	// 对左边树来说，根节点的值是左子树里面的最大值，所以更新higher
	// 对右边树来说，根节点的值是右子树里面的最小值，所以更新lower
	return recurse(root.Left, lower, root.Val) && recurse(root.Right, root.Val, higher)
}

// 中序遍历
// 二叉搜索树「中序遍历」得到的值构成的序列一定是升序的
//  时间复杂度O(n)，空间复杂度O(n)
func isValidBST2(root *TreeNode) bool {
	stack := []*TreeNode{}
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			//  左
			root = root.Left
		}
		// 从栈取值
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		//  如果比上一个节点小，则false
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		// 右
		// 不用加入栈中，因为会在循环一开始的时候就先加入栈
		root = root.Right
	}
	return true
}
