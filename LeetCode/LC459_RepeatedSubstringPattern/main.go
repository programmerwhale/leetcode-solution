package main

import "fmt"

func main() {
	fmt.Println(repeatedSubstringPattern("babbabbab"))
}
func repeatedSubstringPattern(s string) bool {
	return kmp(s+s, s)
}
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
func kmp(haystack string, needle string) bool {
	if len(needle) == 0 {
		return false
	}
	next := make([]int, len(needle))
	getNext(next, needle)
	j := 0 // 模式串的起始位置
	for i := 1; i < len(haystack)-1; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1] // 回退到j的前一位
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) { // j指向了模式串的末尾
			return true
		}
	}
	return false
}
