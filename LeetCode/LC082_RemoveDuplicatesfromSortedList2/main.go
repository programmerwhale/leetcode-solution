package main
/*
82. 删除排序链表中的重复元素 Remove Duplicates from Sorted List
给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。

 

示例 1：


输入：head = [1,2,3,3,4,4,5]
输出：[1,2,5]
示例 2：


输入：head = [1,1,1,2,3]
输出：[2,3]
 

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
	cur.Next=initNode(2)
	cur=cur.Next
	cur.Next=initNode(2)
	cur=cur.Next
	cur.Next=initNode(2)
	cur=cur.Next
	cur.Next=initNode(3)/**/
	printListNode(head)
	l:=deleteDuplicates2(head)
	printListNode(l)
}

// 方法一：mp记录重复数组
// 时间复杂度：O(n)，其中 nn 是链表的长度。
// 空间复杂度：O(n)。
func deleteDuplicates(head *ListNode) *ListNode {
	if head==nil{
		return nil
	}
	mp:=map[int]bool{}
	cur:=head
	for cur.Next!=nil{
		next:=cur.Next
		if next.Val==cur.Val{
			mp[cur.Val]=true
		}
		cur=next
	}

	pre:=&ListNode{0,nil}
	cur=pre
	temp:=head
	for temp!=nil{
		if mp[temp.Val]{
			temp=temp.Next
		}else{
			cur.Next=temp
			temp=temp.Next
			cur=cur.Next
			//断开链接
			cur.Next=nil
		}
	}
	return pre.Next
}

// todo
// 方法二：哑结点
// 时间复杂度：O(n)，其中 nn 是链表的长度。
// 空间复杂度：O(1)。
func deleteDuplicates2(head *ListNode) *ListNode{
	if head == nil {
		return nil
	}
	pre:=&ListNode{0,head}
	cur:=pre
	for cur.Next!=nil&&cur.Next.Next!=nil{
		if cur.Next.Val==cur.Next.Next.Val{
			tmp:=cur.Next.Val
			// 注意cur.Next!=nil
			for cur.Next!=nil&&cur.Next.Val==tmp{
				cur.Next=cur.Next.Next
			}
		}else{
			cur=cur.Next
		}
	}
	return pre.Next
}