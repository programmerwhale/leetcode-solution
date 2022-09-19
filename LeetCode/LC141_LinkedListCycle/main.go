package main

import "fmt"
/*
141. 环形链表
给你一个链表的头节点 head ，判断链表中是否有环。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。

如果链表中存在环 ，则返回 true 。 否则，返回 false 。

 

示例 1：



输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。
示例 2：



输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。
示例 3：



输入：head = [1], pos = -1
输出：false
解释：链表中没有环。
 

提示：

链表中节点的数目范围是 [0, 104]
-105 <= Node.val <= 105
pos 为 -1 或者链表中的一个 有效索引 。
 

进阶：你能用 O(1)（即，常量）内存解决此问题吗？
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func initNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func printListNode(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val)
		fmt.Print(" ")
		l = l.Next
	}
	fmt.Println()
}

func main() {

	head := initNode(3)
	cur := head
	cur.Next = initNode(2)
	cur = cur.Next
	cur.Next = initNode(0)
	cur = cur.Next
	cur.Next = initNode(-4)
	printListNode(head)
	l := hasCycle2(head)
	fmt.Println(l)
}

// 方法一：哈希表
// 时间复杂度o(n),空间复杂度o(n)
func hasCycle(head *ListNode) bool {
	vis := map[*ListNode]int{}
	for head != nil {
		if _, ok := vis[head]; ok {
			return true
		}
		vis[head] = 1
		head = head.Next
	}
	return false
}

// 方法二：快慢指针
// 时间复杂度o(N),空间复杂度o(1)
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head      // 慢指针
	fast := head.Next // 快指针
	// 不重合执行循环
	for slow != fast {
		if fast == nil || fast.Next == nil { // 到链表尾部无环
			return false
		}
		slow = slow.Next      // 慢指针走一步
		fast = fast.Next.Next // 快指针走两步
	}
	// 重合返回true
	return true
}
