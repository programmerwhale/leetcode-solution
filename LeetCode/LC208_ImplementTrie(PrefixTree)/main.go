package main

func main()  {

}

type Trie struct {
	children [26]*Trie
	isEnd bool
}


func Constructor() Trie {
	return Trie{}
}

// 1.插入字符串
// 子节点存在。沿着指针移动到子节点，继续处理下一个字符
// 子节点不存在。创建一个新的子节点，记录在 children 数组的对应位置上，然后沿着指针移动到子节点，继续搜索下一个字符。
// 最后标记end
func (this *Trie) Insert(word string)  {
	for _,ch:=range word{
		// 处理为下标
		ch-='a'
		if this.children[ch]==nil{
			// 占个位置
			this.children[ch]=&Trie{}
		}
		this=this.children[ch]
	}
	this.isEnd=true
}


func (this *Trie) Search(word string) bool {
	node:=this.SearchPrefix(word)
	return node!=nil&&node.isEnd
}

func (this *Trie) SearchPrefix(word string) *Trie {
	for _,ch:=range word{
		ch-='a'
		if this.children[ch]==nil{
			return nil
		}
		this=this.children[ch]
	}
	return this
}


func (this *Trie) StartsWith(prefix string) bool {
	return this.SearchPrefix(prefix) != nil
}
