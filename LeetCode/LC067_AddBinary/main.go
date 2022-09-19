package main

import (
	"fmt"
	"github.com/imroc/biu"
	"strconv"
)

/*
67. 二进制求和
给你两个二进制字符串，返回它们的和（用二进制表示）。

输入为 非空 字符串且只包含数字 1 和 0。



示例 1:

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"


提示：

每个字符串仅由字符 '0' 或 '1' 组成。
1 <= a.length, b.length <= 10^4
字符串如果不是 "0" ，就都不含前导零。
*/
func main() {
	a := "111"
	b := "11"
	fmt.Println(addBinary2(a, b))
}

// 方法一：模拟进位
// 时间复杂度O(n)
// 空间复杂度O(1)
func addBinary(a string, b string) string {
	res := ""
	carry := 0
	for len(a) > 0 || len(b) > 0 {
		if len(a) > 0 {
			carry += int(a[len(a)-1]) - '0'
			a = a[:len(a)-1]
		}
		if len(b) > 0 {
			carry += int(b[len(b)-1]) - '0'
			b = b[:len(b)-1]
		}
		temp := strconv.Itoa(carry % 2)
		res = temp + res
		carry /= 2
	}
	if carry == 1 {
		res = "1" + res
	}
	return res
}

// 方法二：位运算
// 时间复杂度O(n)
// 空间复杂度O(1)
func addBinary2(x string, y string) string {
	a := biu.ToBinaryString(x)

	return fmt.Sprintf("%b", a)
}
