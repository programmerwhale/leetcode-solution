package main

import (
	"fmt"
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
	head := initNode(1)
	head.Left = initNode(2)
	head.Right = initNode(3)
	fmt.Println(sumNumbers2(head))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func initNode(val int) *TreeNode {
	return &TreeNode{val, nil, nil}
}

//dfs
//时间复杂度o(n)
//空间复杂度o(n)
func sumNumbers(root *TreeNode) (sum int) {
	path := 0
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		path = path*10 + root.Val
		if root.Left == nil && root.Right == nil {
			//tmp := path
			sum += path
		}
		dfs(root.Left)
		dfs(root.Right)
		path /= 10
	}
	dfs(root)
	return
}

//bfs
//时间复杂度o(n)
//空间复杂度o(n)
// 广搜 记录当前结点的路径值
type pairs struct {
	node *TreeNode
	num int
}
func sumNumbers2(root *TreeNode) (sum int) {
	if root==nil{
		return
	}
	// 注意初始值
	queue:=[]pairs{pairs{root,root.Val}}
	for len(queue)>0{
		pair:=queue[0]
		queue=queue[1:]
		left,right:=pair.node.Left,pair.node.Right
		num:=pair.num
		// 终止条件：到叶子节点了直接加
		if left==nil&&right==nil{
			sum+=num
		}else{
			if left!=nil{
				queue=append(queue,pairs{left,num*10+left.Val})
			}
			if right!=nil{
				queue=append(queue,pairs{right,num*10+right.Val})
			}
		}
	}
	return
}











