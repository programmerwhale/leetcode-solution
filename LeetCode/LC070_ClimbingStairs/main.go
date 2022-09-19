package main

import "fmt"

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

 

示例 1：

输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶
示例 2：

输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶

*/
func main()  {
	fmt.Println(ClimbStairs2(45))
}
// 方法一：动态规划
// dp[i]含义为，台阶为i时有dp[i]种方法到楼顶
// 方程为dp[i]=dp[i-1]+dp[i-2]
// 结束条件：dp[n]
// 时间复杂度O(n),空间复杂度O(n)
func ClimbStairs(n int) int {
	if n<2{
		return n
	}
	dp:=make([]int,n+1)
	dp[0]=1
	dp[1]=1

	for i:=2;i<=n;i++{
		dp[i]=dp[i-1]+dp[i-2]
	}
	return dp[n]
}

// 方法二：斐波那契
// 根据dp[i]=dp[i-1]+dp[i-2]，发现可以用斐波那契算
// 时间复杂度O(n),空间复杂度O(1)
func ClimbStairs2(n int) int {
	if n<2{
		return n
	}
	n1,n2:=1,1

	for i:=2;i<=n;i++{
		n1,n2=n2,n1+n2
	}
	return n2
}