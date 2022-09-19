package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords2("Let's take LeetCode contest"))
}

// 方法一：利用内置函数
// 时间复杂度o(n), 空间复杂度o(n)
func reverseWords(s string) string {
	str := []byte(s)
	//先从头到尾转
	for i := 0; i < len(str)/2; i++ {
		str[i], str[len(str)-i-1] = str[len(str)-i-1], str[i]
	}
	ans := strings.Split(string(str), " ")
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}
	res := strings.Join(ans, " ")
	return res
}

// 方法二：手写
//
func reverseWords2(s string) string {
	length := len(s)
	ret := []byte{}
	// i在逻辑里面控制
	for i := 0; i < length; {
		// 记录开始位置
		start := i
		// 找到结束位置
		for i < length && s[i] != ' ' {
			i++
		}
		// 得到单词起始位置，从后面往ret里面加，i不动
		for p := start; p < i; p++ {
			// start+i-1 是结尾下标
			// -p去掉当前位置
			ret = append(ret, s[start+i-1-p])
		}
		// 加空格
		// i < length是为了结尾不加
		// 用for是为了有几个空格就加几个
		if i < length && s[i] == ' ' {
			i++
			ret = append(ret, ' ')
		}
	}
	return string(ret)
}
