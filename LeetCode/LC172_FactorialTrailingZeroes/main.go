package main

import "fmt"

/*
172. 阶乘后的零
给定一个整数 n ，返回 n! 结果中尾随零的数量。

提示 n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1



示例 1：

输入：n = 3
输出：0
解释：3! = 6 ，不含尾随 0
示例 2：

输入：n = 5
输出：1
解释：5! = 120 ，有一个尾随 0
示例 3：

输入：n = 0
输出：0


提示：

0 <= n <= 104


进阶：你可以设计并实现对数时间复杂度的算法来解决此问题吗？*/
func main()  {
	fmt.Println(trailingZeroes(10000))
}

// 方法一：数学
// 时间复杂度O(N)，空间复杂度O(1)
func trailingZeroes(n int) (cnt int) {
	for i:=5;i<=n;i+=5{
		for j:=i;j%5==0;j/=5{
			cnt++
		}
	}
	return
}

// 方法二：优化
// 时间复杂度O(logN)，空间复杂度O(1)
func trailingZeroes2(n int) (cnt int) {
	i:=5
	for i<=n{
		cnt+=n/i
		i*=5
	}
	return
}