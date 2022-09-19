package main

import "fmt"

func main() {
	a := initNode(5, nil)
	b := initNode(4, a)
	c := initNode(3, b)
	d := initNode(2, c)
	head := initNode(1, d)
	PrintListNode(head)
	reverseList2(head)
	fmt.Println("----------------------------")
	PrintListNode(a)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func initNode(val int, next *ListNode) *ListNode {
	return &ListNode{
		Val:  val,
		Next: next,
	}
}

func PrintListNode(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

// 方法一：迭代
// 反转需要交换前要先记录下一个节点，为了反转后归位
// 时间复杂度o(n),空间复杂度o(1)
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	// 注意return的是pre，此时cur为nil
	return prev
}

//  递归
/*	1. 确定终止条件，head==nil一定能写出来；
	2. 递归公式：返回值为下一个节点
	3. 记录下一个位置（head.Next.Next，此时发现终止条件应该加一个head.Next!=nil），为了归位
	4. head.Next=nil，否则可能会产生环*/
// 时间复杂度o(n),空间复杂度o(n)
func reverseList2(head *ListNode) *ListNode {
	// 一直找到倒数第一个元素再停止，让尾巴两个换位置
	if head == nil || head.Next == nil {
		return head
	}
	// 哑节点 一直是个nil
	newNode := reverseList2(head.Next)
	// 第一次到这里的时候只有head为第二个元素，newnode为另一侧的nil
	// 绕成一个圈
	head.Next.Next = head
	// 并断开之前的节点
	head.Next = nil
 	return newNode
}
