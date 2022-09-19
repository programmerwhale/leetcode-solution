package main

import "fmt"

/*
121. 买卖股票的最佳时机
给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

 

示例 1：

输入：[7,1,5,3,6,4]
输出：5
解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
示例 2：

输入：prices = [7,6,4,3,1]
输出：0
解释：在这种情况下, 没有交易完成, 所以最大利润为 0。
 

提示：

1 <= prices.length <= 105
0 <= prices[i] <= 104
*/
func main()  {
	prices:=[]int{7,6,4,3,1}
	fmt.Println(maxProfit2(prices))
}
// dp代表在当前天的最高利润
// 需要定义一个最小成本cost
// 动态转移方程为dp[i]=max(dp[i-1],prices[i]-cost)
// 时间复杂度O(n)，空间复杂度O(n)
func maxProfit(prices []int) int {
	cost:=prices[0]
	n:=len(prices)
	dp:=make([]int,n)
	dp[0]=0
	for i:=1;i<n;i++{
		if prices[i]<cost{
			cost=prices[i]
		}
		dp[i]=max(dp[i-1],prices[i]-cost)
	}
	return dp[n-1]
}

// 优化，因为只用到了dp[i-1]，所以可以只用一个值去记录然后更新
// 时间复杂度O(n)，空间复杂度O(1)
func maxProfit2(prices []int) int {
	cost:=prices[0]
	n:=len(prices)
	dp:=0
	for i:=1;i<n;i++{
		if prices[i]<cost{
			cost=prices[i]
		}
		dp=max(dp,prices[i]-cost)
	}
	return dp
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}