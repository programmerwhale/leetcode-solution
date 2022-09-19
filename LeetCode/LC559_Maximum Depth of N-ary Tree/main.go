package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

// 层次遍历
func maxDepth2(root *Node) (res int) {
	if root == nil {
		return
	}
	q := []*Node{root}
	for len(q) > 0 {
		p := []*Node{}
		for len(q) > 0 {
			node := q[0]
			q = q[1:]
			for _, v := range node.Children {
				p = append(p, v)
			}
		}
		q = p
		res++
	}
	return
}

// 递归
func maxDepth(root *Node) (res int) {
	if root == nil {
		return
	}
	tmp := 0
	for _, node := range root.Children {
		tmp = max(tmp, maxDepth(node))
	}
	res = tmp + 1
	return
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
