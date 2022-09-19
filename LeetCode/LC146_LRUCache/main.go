package main

import "fmt"

/*
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

示例：

输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4

1 <= capacity <= 3000
0 <= key <= 10000
0 <= value <= 105
最多调用 2 * 105 次 get 和 put
*/

/*
思路:
方法一：用数组来存储
	缺点：每一次都需要遍历数组修改标记
	时间复杂度：O(n)
方法二：单链表
	缺点：顺序访问链表
	时间复杂度：O(n)
方法三：双向链表+哈希表
	快速定位位置，
	Get：在map里获取Node O(1)
	Put：在map里看是否已存在，再去双链表操作
	时间复杂度：O(1)

	为什么用双链表？
	因为需要删除操作，需要获取都前驱节点位置
*/
type Node struct {
	Val int
	Key int
	Next *Node
	Prev *Node
}

func initNode(key,value int)*Node{
	return &Node{Key: key,Val: value,Next: nil,Prev: nil}
}

type DoubleLinkedList struct {
	size int
	head,tail *Node
}

func initDoubleLinkedList() *DoubleLinkedList{
	dll:=&DoubleLinkedList{
		size: 0,
		head: initNode(0,0),
		tail: initNode(0,0),
	}
	//注意这里处理头尾节点的关系
	dll.head.Next=dll.tail
	dll.tail.Prev=dll.head
	return dll
}

func (this *DoubleLinkedList)AddLast(x *Node){
	x.Prev=this.tail.Prev
	x.Next=this.tail
	// 顺序不能反
	this.tail.Prev.Next=x
	this.tail.Prev=x
	// size++
	this.size++
}

func (this *DoubleLinkedList)Remove(x *Node){
	x.Prev.Next=x.Next
	x.Next.Prev=x.Prev
	x.Prev=nil
	x.Next=nil
	this.size--
}

func (this *DoubleLinkedList)RemoveFirst() *Node {
	//如果首尾相邻
	if this.head==this.tail.Prev{
		return nil
	}
	node:=this.head.Next
	this.Remove(node)
	return node
}

type LRUCache struct {
	capacity int
	knmp map[int]*Node
	cache *DoubleLinkedList
}

func Construct(capacity int)*LRUCache{
	return &LRUCache{
		capacity:capacity,
		knmp: map[int]*Node{},
		cache: initDoubleLinkedList(),
	}
}

func (this *LRUCache)Get(key int) int {
	node,exist:=this.knmp[key]
	if !exist{
		fmt.Println(-1)
		return -1
	}else{
		this.cache.Remove(node)
		this.cache.AddLast(node)
		fmt.Println(node.Val)
		return node.Val
	}
}

func (this *LRUCache)Put(key int,value int){
	node,exist:=this.knmp[key]
	newNode:=initNode(key,value)
	if exist{
		this.cache.Remove(node)
		delete(this.knmp,key)
	}else{
		if this.capacity==this.cache.size{
			first:=this.cache.RemoveFirst()
			delete(this.knmp,first.Key)
		}
	}
	this.cache.AddLast(newNode)
	this.knmp[key]=newNode
}

func main()  {
	lRUCache :=Construct(2)
	lRUCache.Put(1,1)
	lRUCache.Get(1);    // 返回 1
	lRUCache.Put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	lRUCache.Get(2);    // 返回 -1 (未找到)
	lRUCache.Put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	lRUCache.Get(1);    // 返回 -1 (未找到)
	lRUCache.Get(3);    // 返回 3
	lRUCache.Get(4);    // 返回 4
}
