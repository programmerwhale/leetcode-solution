package main

import (
	"fmt"
)

func main() {
	haystack := "helphello"
	needle := "ll"
	fmt.Println(strStr2(haystack, needle))
}

/*
本题是经典的字符串单模匹配的模型，因此可以使用字符串匹配算法解决，常见的字符串匹配算法包括暴力匹配、
Knuth-Morris-Pratt 算法、
Boyer-Moore 算法、
Sunday 算法等，本文将讲解
Knuth-Morris-Pratt 算法。

因为哈希方法可能出现哈希值相等但是字符串不相等的情况，而strStr函数要求匹配结果必定正确，
因此本文不介绍哈希方法，有兴趣的读者可以自行了解滚动哈希的实现（如Rabin-Karp 算法）。
*/

// 方法一：暴力法
// 时间复杂度：o(n*m), 空间复杂度o(1)
func strStr(haystack string, needle string) int {
	hlen, nlen := len(haystack), len(needle)
	// 这里可以等于 有两个串长度相等的情况
	for i := 0; i+nlen <= hlen; i++ {
		if haystack[i:i+nlen] == needle {
			return i
		}
	}
	return -1
}

// todo
// 方法二：kmp
// 构造前缀表
func getNext(next []int, s string) {
	j := 0 // j表示 最长相等前后缀长度
	next[0] = j

	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1] // 回退前一位
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j // next[i]是i（包括i）之前的最长相等前后缀长度
	}
}
func strStr2(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := make([]int, len(needle))
	getNext(next, needle)
	j := 0 // 模式串的起始位置
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1] // 回退到j的前一位
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) { // j指向了模式串的末尾
			return i - len(needle) + 1
		}
	}
	return -1
}

// todo
// 方法三：Sunday
// 每次匹配都会从 目标字符串中 提取 待匹配字符串与 模式串 进行匹配：
//
//若匹配，则返回当前 idx
//不匹配，则查看 待匹配字符串 的后一位字符 c：
//若c存在于Pattern中，则 idx = idx + 偏移表[c]
//否则，idx = idx + len(pattern)
//Repeat Loop 直到 idx + len(pattern) > len(String)

func strStr3(haystack string, needle string) int {
	hsize := len(haystack)
	nsize := len(needle)
	dic := map[byte]int{}
	for i, v := range needle {
		dic[byte(v)] = nsize - i
	}

	i := 0
	for i <= hsize-nsize {

	}
	return -1
}
