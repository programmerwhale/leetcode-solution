package main

import (
	"fmt"
	"sort"
)

func main()  {
	words:=[]string{"yo","ew","fc","zrc","yodn","fcm","qm","qmo","fcmz","z","ewq","yod","ewqz","y"}
	fmt.Println(longestWord2(words))
}

// 方法一：排序+hash
// 时间复杂度O(nlgn)，空间复杂度O(n)
func longestWord(words []string) string {
	sort.Slice(words,func(i,j int)bool{
		s,t:=words[i],words[j]
		return len(s) < len(t) || len(s) == len(t) && s > t
	})
	// 排序后[z y yo qm fc ew zrc yod qmo fcm ewq yodn fcmz ewqz]
	mp:=map[string]bool{}
	// 需要给个初始值
	mp[""]=true
	res:=""
	for _,v:=range words{
		if _,ok:=mp[v[:len(v)-1]];ok{
			res=v
			mp[v]=true
		}
	}
	return res
}

// todo
// 方法二：前缀树
// 每一个前缀都在words中

func longestWord2(words []string) (longest string) {
	t:=&Trie{}
	for _,v:=range words{
		t.Insert(v)
	}
	for _,v:=range words{
		if t.Search(v)&& (len(v) > len(longest) || len(v) == len(longest) && v < longest) {
			longest = v
		}
	}
	return longest
}

type Trie struct {
	children [26]*Trie
	isEnd bool
}
func Construct() Trie{
	return Trie{}
}

// 1.插入字符串
// 子节点存在。沿着指针移动到子节点，继续处理下一个字符
// 子节点不存在。创建一个新的子节点，记录在 children 数组的对应位置上，然后沿着指针移动到子节点，继续搜索下一个字符。
// 最后标记end
func(t *Trie)Insert(word string){
	node:=t
	for _,ch:=range word{
		// 处理为下标
		ch-='a'
		if node.children[ch]==nil{
			// 占个位置
			node.children[ch]=&Trie{}
		}
		node=node.children[ch]
	}
	node.isEnd=true
}

// 2.查找
func(t *Trie)Search(word string)bool{
	for _,ch:=range word{
		ch-='a'
		// 到这就结束了 或者其中一步没有子串存在
		if t.children[ch]==nil||!t.children[ch].isEnd{
			return false
		}
		t=t.children[ch]
	}
	return true
}