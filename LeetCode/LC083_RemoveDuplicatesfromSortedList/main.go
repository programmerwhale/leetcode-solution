package main
/*
83. 删除排序链表中的重复元素 Remove Duplicates from Sorted List
给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。

给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
输入：head = [1,1,2]
输出：[1,2]
输入：head = [1,1,2,3,3]
输出：[1,2,3]
 

提示：

链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序 排列
*/
import (
	"fmt"
)

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
	cur.Next=initNode(1)
	cur=cur.Next
	cur.Next=initNode(2)
	cur=cur.Next
	cur.Next=initNode(2)
	cur=cur.Next
	cur.Next=initNode(2)
	printListNode(head)
	l:=deleteDuplicates(head)
	printListNode(l)
}

// 遍历
// 时间复杂度：O(n)，其中 nn 是链表的长度。
//空间复杂度：O(1)。
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			// 核心，直接跳过中间的
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}