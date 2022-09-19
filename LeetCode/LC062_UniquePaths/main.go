package main

import "fmt"
/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

问总共有多少条不同的路径？
*/
func main()  {
	fmt.Println(uniquePaths(3,7))
}

// 方法一：动态规划
// dp[i][j]含义为到点（i，j）可能的路径数
// 方程为dp[i][j]=dp[i-1][j]+dp[i][j-1]
// 终点为dp[m-1][n-1]
// 空间复杂度O(mn),时间复杂度O(mn)
func uniquePaths(m int, n int) int {
	dp:=make([][]int,m)
	for i,_:=range dp{
		dp[i]=make([]int,n)
		dp[i][0]=1
	}
	for i:=0;i<n;i++{
		dp[0][i]=1
	}
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			// 当前点的路径总数=上一个节点的+左一个节点的
			dp[i][j]=dp[i-1][j]+dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}