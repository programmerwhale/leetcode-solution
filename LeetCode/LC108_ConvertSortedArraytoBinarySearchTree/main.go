package main

import "fmt"

/*
108.Convert Sorted Array to Binary Search Tree将有序数组转换为二叉搜索树
给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。
高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。

示例 1：
输入：nums = [-10,-3,0,5,9]
输出：[0,-3,9,-10,null,5]
解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：

示例 2：
输入：nums = [1,3]
输出：[3,1]
解释：[1,3] 和 [3,1] 都是高度平衡二叉搜索树。

提示：
1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums 按 严格递增 顺序排列
*/
func main() {
	nums := []int{-10, -3, 0, 5, 9}
	root := sortedArrayToBST(nums)
	fmt.Println(preorderTraversal(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

// 递归法
// 时间复杂度o(n), 空间复杂度o(logn)
func sortedArrayToBST(nums []int) *TreeNode {
	n:=len(nums)
	return construct(nums,0,n-1)

}
func construct(nums []int,left,right int)*TreeNode{
	if left>right{
		return nil
	}
	mid:=(left+right)/2
	// 选择任意一个中间位置数字作为根节点
	//mid := (left + right + rand.Intn(2)) / 2
	root:=&TreeNode{Val:nums[mid]}
	root.Left=construct(nums,left,mid-1)
	root.Right=construct(nums,mid+1,right)
	return root
}
