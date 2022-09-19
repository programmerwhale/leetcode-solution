package main
/*
143. 重排链表
给定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1：
输入：head = [1,2,3,4]
输出：[1,4,2,3]

示例 2：
输入：head = [1,2,3,4,5]
输出：[1,5,2,4,3]

提示：
链表的长度范围为 [1, 5 * 104]
1 <= node.val <= 1000
*/
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
	fmt.Println(findMid(head).Val)
	//l:=reverseList(head)
	//printListNode(l)
	reorderList(head)
	printListNode(head)
}

// 找到中间节点，用快慢指针
// 从中间节点之后的值开始反转
// 再合并
// 时间复杂度O(N)，空间复杂度O(1)
func reorderList(head *ListNode)  {
	if head==nil{
		return
	}
	mid:=findMid(head)
	l1:=head
	l2:=mid.Next
	mid.Next=nil
	l2=reverseList(l2)
	mergeList(l1,l2)
}
func findMid(head *ListNode)*ListNode{
	slow,fast:=head,head
	// 注意这里结束处理
	for fast.Next!=nil&&fast.Next.Next!=nil{
		slow=slow.Next
		fast=fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode)*ListNode{
	var pre, cur *ListNode = nil, head
	for cur!=nil{
		next:=cur.Next
		cur.Next=pre
		pre=cur
		cur=next
	}
	return pre
}
func mergeList(l1,l2  *ListNode)*ListNode {
	var dummy *ListNode
	dummy=initNode(0)
	cur:=dummy
	for l1!=nil&&l2!=nil{
		cur.Next=l1
		l1=l1.Next
		cur=cur.Next
		cur.Next=l2
		cur=cur.Next
		l2=l2.Next
	}
	if l1!=nil{
		cur.Next=l1
	}
	if l2!=nil{
		cur.Next=l2
	}
	return dummy.Next
}
