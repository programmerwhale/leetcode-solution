package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
// o(n) o(n)
func tree2str(root *TreeNode) string {
	switch {
	case root == nil:
		return ""
	case root.Left == nil && root.Right == nil:
		return strconv.Itoa(root.Val)
	case root.Right == nil:
		return fmt.Sprintf("%d(%s)", root.Val, tree2str(root.Left))
	default:
		return fmt.Sprintf("%d(%s)(%s)", root.Val, tree2str(root.Left), tree2str(root.Right))
	}
}

//迭代
// o(n) o(n)
func tree2str2(root *TreeNode) string {
	ans := &strings.Builder{}
	stack := []*TreeNode{root}
	vis := map[*TreeNode]bool{}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if vis[node] {
			if node != root {
				ans.WriteByte(')')
			}
			stack = stack[:len(stack)-1]
		} else {
			vis[node] = true
			// 对根节点特殊处理
			if node != root {
				ans.WriteByte('(')
			}
			ans.WriteString(strconv.Itoa(node.Val))
			if node.Left == nil && node.Right != nil {
				ans.WriteString("()")
			}
			// 前序遍历：先右再左
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		}
	}
	return ans.String()
}
