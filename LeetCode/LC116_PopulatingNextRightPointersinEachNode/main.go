package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func initNode(val int) *Node {
	return &Node{Val: val}
}

func main() {
	head := initNode(1)
	cur := head
	cur.Left = initNode(2)
	cur = head.Left
	cur.Left = initNode(4)
	cur.Right = initNode(5)
	head.Right = initNode(3)
	cur = head.Right
	cur.Left = initNode(6)
	cur.Right = initNode(7)

	fmt.Println(preorderTraversal(head))
	connect(head)
	fmt.Println(preorderTraversal(head))
}

func preorderTraversal(root *Node) []int {
	res := []int{}
	dfs(&res, root)
	return res
}

func dfs(res *[]int, root *Node) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	dfs(res, root.Left)
	dfs(res, root.Right)
}

// bfs：不改变原有结构，只是加了个nil
// 时间复杂度：o(n)
// 空间复杂度：o(n)
// 这是一棵完美二叉树，它的最后一个层级包含N/2 个节点。广度优先遍历的复杂度取决于一个层级上的最大元素数量。这种情况下空间复杂度为O(N)。

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	// 外层循环的是层数
	for len(q) > 0 {
		tmp := q
		q = nil
		for i, node := range tmp {
			//  链接下一节点，注意这里加了一个判断条件i+1<len(tmp)
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}
			// 把左右节点加入队列
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	res := strings.HasPrefix("leetcode", "leet")
	fmt.Println(res)
	return root
}
