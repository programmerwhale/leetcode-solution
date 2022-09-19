package main

import "container/heap"

type ListNode struct {
	Val int
	Next *ListNode
}
func main()  {
	lists:=[]*ListNode{}
	l1:=&ListNode{0,nil}
	lists=append(lists,l1)
	l2:=&ListNode{1,nil}
	lists=append(lists,l2)
	l3:=&ListNode{3,nil}
	lists=append(lists,l3)
	mergeKLists(lists)
}
// 方法一：分支合并，两两排序
func mergeKLists(lists []*ListNode) *ListNode {
	n:=len(lists)
	if n==0{
		return nil
	}
	if n==1{
		return lists[0]
	}
	mid:=n/2
	l1:=mergeKLists(lists[:mid])
	l2:=mergeKLists(lists[mid:])
	return merge2list(l1,l2)
}

func merge(l1,l2 *ListNode)*ListNode{

	head:=&ListNode{}
	cur:=head
	for l1!=nil&&l2!=nil{
		if l1.Val<=l2.Val{
			cur.Next=l1
			l1=l1.Next
		}else{
			cur.Next=l2
			l2=l2.Next
		}
		cur=cur.Next
	}
	if l1!=nil{
		cur.Next=l1
	}
	if l2!=nil{
		cur.Next=l2
	}
	return head.Next
}


// 合并两个有序链表
// 时间复杂度：每一组是O(2N),合并logK次
func merge2list(list1, list2 *ListNode) *ListNode {
	head := &ListNode{}
	ans := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			ans.Next = list1
			list1 = list1.Next
		} else {
			ans.Next = list2
			list2 = list2.Next
		}
		ans = ans.Next
	}
	if list1 != nil {
		ans.Next = list1
	}
	if list2 != nil {
		ans.Next = list2
	}
	return head.Next
}

// 优先级队列
type Heap struct {
	*_heap
}

func (h *Heap) Push(v interface{}) {
	heap.Push(h._heap, v)
}

func (h *Heap) Pop() interface{} {
	return heap.Pop(h._heap)
}

func (h *Heap) Fix(i int) {
	heap.Fix(h._heap, i)
}

type _heap struct {
	slice []*ListNode
	len   int
}

func CreateHeap() Heap {
	h := &_heap{}
	heap.Init(h)
	return Heap{h}
}

func (h *_heap) Len() int {
	return h.len
}

func (h *_heap) Less(i, j int) bool {
	// 最小堆
	return (h.slice)[i].Val < (h.slice)[j].Val
}

func (h *_heap) Swap(i, j int) {
	(h.slice)[i], (h.slice)[j] = (h.slice)[j], (h.slice)[i]
}

func (h *_heap) Push(x interface{}) {
	h.len++
	h.slice = append(h.slice, x.(*ListNode))
}

func (h *_heap) Pop() (ret interface{}) {
	if h.len != 0 {
		ret = (h.slice)[h.Len()-1]
		h.slice = (h.slice)[:h.Len()-1]
		h.len--
		return
	}
	return
}

func mergeKLists2(lists []*ListNode) *ListNode {
	head := &ListNode{}
	ans := head
	hp := CreateHeap()
	// 建堆
	for _, list := range lists {
		if list != nil {
			hp.Push(list)
		}
	}
	// 依次取出
	for hp.Len() != 0 {
		v := hp.Pop().(*ListNode)
		ans.Next = v
		ans = ans.Next
		// 如果下一个不为nil,入队
		if v.Next != nil {
			hp.Push(v.Next)
		}
	}
	return head.Next
}
