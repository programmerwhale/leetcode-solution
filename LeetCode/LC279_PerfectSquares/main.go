package main

import (
	"fmt"
	"math"
)

/*
279. 完全平方数
给你一个整数 n ，返回 和为 n 的完全平方数的最少数量 。

完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。

 

示例 1：

输入：n = 12
输出：3
解释：12 = 4 + 4 + 4
示例 2：

输入：n = 13
输出：2
解释：13 = 4 + 9
 
提示：

1 <= n <= 104
*/
func main()  {
	fmt.Println(numSquares(12))
}
func numSquares(n int) int {
	// 申请n+1个内存空间
	dp:=make([]int,n+1)
	// 初始值为0
	dp[0]=0
	for i:=1;i<=n;i++{
		minnum:=math.MaxInt64
		for j:=1;j*j<=i;j++{
			// 从1开始，即dp[i-1]
			minnum=min(minnum,dp[i-j*j])
		}
		// 在这里统一加一
		dp[i]=minnum+1
	}
	return dp[n]
}

func min(a,b int) int {
	if a<b{
		return a
	}
	return b
}