package main

import "fmt"

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

func main()  {

	head:=initNode(1)
	cur:=head
	cur.Next=initNode(2)
	cur=cur.Next
	cur.Next=initNode(3)
	cur=cur.Next
	cur.Next=initNode(4)
	cur=cur.Next
	cur.Next=initNode(5)
	printListNode(head)
	l:=rotateRight(head,2)
	printListNode(l)
}

// k小于链表长度的情况
func rotateRight(head *ListNode, k int) *ListNode {
	if k==0 ||head==nil||head.Next==nil{
		return  head
	}
	// 求链表长度
	n := 1
	iter := head
	for iter.Next != nil {
		iter = iter.Next
		n++
	}
	k%=n
	node1,node2:=head,head
	pre:=head
	if k==0{
		return head
	}
	// 快慢指针
	for k>1{
		node1=node1.Next
		k--
	}
	for node1!=nil&&node1.Next!=nil{
		node1=node1.Next
		pre=node2
		node2=node2.Next
	}
	// 前置节点的next断开
	pre.Next=nil
	// 尾部连上头
	node1.Next=head
	return node2
}