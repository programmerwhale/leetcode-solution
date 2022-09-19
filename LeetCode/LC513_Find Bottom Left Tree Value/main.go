package main

func main() {

}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 回溯
var maxDeep int // 全局变量 深度
var value int   //全局变量 最终值
func findBottomLeftValue(root *TreeNode) int {
	if root.Left == nil && root.Right == nil { //需要提前判断一下（不要这个if的话提交结果会出错，但执行代码不会。防止这种情况出现，故先判断是否只有一个节点）
		return root.Val
	}
	findLeftValue(root, maxDeep)
	return value
}

func findLeftValue(root *TreeNode, deep int) {
	//最左边的值在左边
	if root.Left == nil && root.Right == nil {
		if deep > maxDeep {
			value = root.Val
			maxDeep = deep
		}
	}
	//递归
	if root.Left != nil {
		deep++
		findLeftValue(root.Left, deep)
		deep-- //回溯
		//findLeftValue(root.Left, deep+1)
	}
	if root.Right != nil {
		deep++
		findLeftValue(root.Right, deep)
		deep-- //回溯
		//findLeftValue(root.Right, deep+1)
	}
}

// 迭代,层序遍历
// O(N),O(N)
func findBottomLeftValue2(root *TreeNode) (res int) {
	q := []*TreeNode{root}
	for len(q) > 0 {
		p := []*TreeNode{}
		for i, node := range q {
			if i == 0 {
				res = node.Val
			}
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return
}
