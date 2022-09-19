package main

import "fmt"

/*
99. Recover Binary Search Tree恢复二叉搜索树
给你二叉搜索树的根节点 root ，该树中的 恰好 两个节点的值被错误地交换。请在不改变其结构的情况下，恢复这棵树 。

示例 1：

输入：root = [1,3,null,null,2]
输出：[3,1,null,null,2]
解释：3 不能是 1 的左孩子，因为 3 > 1 。交换 1 和 3 使二叉搜索树有效。
示例 2：


输入：root = [3,1,4,null,null,2]
输出：[2,1,4,null,null,3]
解释：2 不能在 3 的右子树中，因为 2 < 3 。交换 2 和 3 使二叉搜索树有效。

提示：
树上节点的数目在范围 [2, 1000] 内
-231 <= Node.val <= 231 - 1

进阶：使用 O(n) 空间复杂度的解法很容易实现。你能想出一个只使用 O(1) 空间的解决方案吗？
*/
func main() {
	root := &TreeNode{Val: 3, Left: &TreeNode{5, nil, initTreeNode(2)}, Right: &TreeNode{1, initTreeNode(4), nil}}
	fmt.Println(preorderTraversal(root))
	recoverTree2(root)
	fmt.Println(preorderTraversal(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func initTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	dfs(&res, root)
	return res
}

func dfs(res *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	dfs(res, root.Left)
	dfs(res, root.Right)
}

//笨方法：显式中序遍历
// 时间复杂度O(N),空间复杂度O(N)
func recoverTree(root *TreeNode) {
	ans := []*TreeNode{}
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		ans = append(ans, root)
		dfs(root.Right)
	}
	dfs(root)
	l, r := 0, 0
	for i := 0; i < len(ans)-1; i++ {
		if ans[i].Val > ans[i+1].Val {
			l = i
			break
		}
	}
	for i := len(ans) - 1; i > 0; i-- {
		if ans[i-1].Val > ans[i].Val {
			r = i
			break
		}
	}
	node1 := ans[l]
	node2 := ans[r]
	node1.Val, node2.Val = node2.Val, node1.Val
}

//隐式中序遍历
// 时间复杂度O(N),空间复杂度O(H)
func recoverTree2(root *TreeNode){
	stack:=[]*TreeNode{}
	var x,y,pre *TreeNode
	for len(stack)>0|| root != nil{
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre!=nil&&root.Val<pre.Val{
			//2
			y=root
			if x==nil{
				//3
				x=pre
			}else {
				break
			}
		}
		//1
		pre=root


		root=root.Right
	}
	x.Val,y.Val=y.Val,x.Val
}

//Morris
func recoverTree3(root *TreeNode)  {
	var x, y, pred, predecessor *TreeNode

	for root != nil {
		if root.Left != nil {
			// predecessor 节点就是当前 root 节点向左走一步，然后一直向右走至无法走为止
			predecessor = root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}

			// 让 predecessor 的右指针指向 root，继续遍历左子树
			if predecessor.Right == nil {
				predecessor.Right = root
				root = root.Left
			} else { // 说明左子树已经访问完了，我们需要断开链接
				if pred != nil && root.Val < pred.Val {
					y = root
					if x == nil {
						x = pred
					}
				}
				pred = root
				predecessor.Right = nil
				root = root.Right
			}
		} else { // 如果没有左孩子，则直接访问右孩子
			if pred != nil && root.Val < pred.Val {
				y = root
				if x == nil {
					x = pred
				}
			}
			pred = root
			root = root.Right
		}
	}
	x.Val, y.Val = y.Val, x.Val
}