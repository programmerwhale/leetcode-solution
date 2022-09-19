/*
LC884. 两句话中的不常见单词
句子 是一串由空格分隔的单词。每个 单词 仅由小写字母组成。

如果某个单词在其中一个句子中恰好出现一次，在另一个句子中却 没有出现 ，那么这个单词就是 不常见的 。

给你两个 句子 s1 和 s2 ，返回所有 不常用单词 的列表。返回列表中单词可以按 任意顺序 组织。

 

示例 1：

输入：s1 = "this apple is sweet", s2 = "this apple is sour"
输出：["sweet","sour"]
示例 2：

输入：s1 = "apple apple", s2 = "banana"
输出：["banana"]
 

提示：

1 <= s1.length, s2.length <= 200
s1 和 s2 由小写英文字母和空格组成
s1 和 s2 都不含前导或尾随空格
s1 和 s2 中的所有单词间均由单个空格分隔

分类：
字符串/哈希/

思路：
一般验证是都存在是否重复的题都优先考虑用哈希表做，可以快速定位
仔细读题，
该题为返回两个字符串中所有不重复的单词，可以遍历两个数组并记录出现的次数放进map；结束后遍历之出现一次的次数
还有题目要输出s2中不在s1里的单词，思路大体一致；遍历s1中的单词放入map，再用s2中的单词去看是否在map中。
*/
package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(uncommonFromSentences("this apple is sweet","this apple is sour"))
}

func uncommonFromSentences(s1 string, s2 string) (res []string) {
	ans1:=strings.Fields(s1)
	ans2:=strings.Fields(s2)
	mp:=make(map[string]int,len(ans1)+len(ans2))
	for _,v:=range ans1{
		mp[v]++
	}
	for _,v:=range ans2{
		mp[v]++
	}
	for i,v:=range mp{
		if v==1{
			res=append(res,i)
		}
	}
	return res
}