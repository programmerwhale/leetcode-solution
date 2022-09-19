package main

import (
	"fmt"
	"math"
)

/*todo
121. 二叉树中的最大路径和
路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其 最大路径和 。



示例 1：


输入：root = [1,2,3]
输出：6
解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
示例 2：


输入：root = [-10,9,20,null,null,15,7]
输出：42
解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42


提示：

树中节点数目范围是 [1, 3 * 104]
-1000 <= Node.val <= 1000
*/
func main() {
	head := initNode(-10)
	head.Left = initNode(9)
	head.Right = initNode(20)
	cur := head
	cur = cur.Right
	cur.Left = initNode(15)
	cur.Right = initNode(7)
	fmt.Println(maxPathSum(head))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func initNode(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt64
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)
		priceNewPath := node.Val + leftGain + rightGain
		maxSum = max(maxSum, priceNewPath)
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
