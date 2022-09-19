package main

import "fmt"

func main()  {
	root:=initNode(1)
	cur:=root
	cur.Next=initNode(2)
	cur=cur.Next
	cur.Next=initNode(3)
	cur=cur.Next
	cur.Next=initNode(4)
	printListNode(root)
	printListNode(swapPairs2(root))
}

type ListNode struct {
	Val int
	Next *ListNode
}

func initNode(val int)*ListNode{
	return &ListNode{Val: val}
}

func printListNode(l *ListNode)  {
	for l!=nil{
		fmt.Print(l.Val)
		fmt.Print(" ")
		l=l.Next
	}
	fmt.Println()
}

// 方法一
// 递归
// 时间复杂度O(N)，节点个数
// 空间复杂度O(N)，递归调用的栈的深度，也就是节点个数
func swapPairs(head *ListNode) *ListNode {
	if head==nil || head.Next==nil{
		return head
	}
	// 记录head的下一个值，也就是会成为head前置节点的值
	newHead:=head.Next
	// 两两交换，下一个起始点为第三个节点
	head.Next=swapPairs(newHead.Next)
	// 使第二个节点指向第一个节点
	newHead.Next=head
	// 返回新节点
	return newHead
}

// 方法二
// 迭代
// 时间复杂度O(N)，节点个数
// 空间复杂度O(1)
func swapPairs2(head *ListNode) *ListNode {
	temp:=&ListNode{0,head}
	pre:=temp
	for pre.Next!=nil&&pre.Next.Next!=nil{
		cur:=pre.Next
		next:=cur.Next
		pre.Next=next
		cur.Next=next.Next
		next.Next=cur
		// 最后和原来的第一个节点交换
		pre=cur
	}
	return temp.Next
}