/*
1765. 地图中的最高点 | Mid

给你一个大小为 m x n 的整数矩阵 isWater ，它代表了一个由 陆地 和 水域 单元格组成的地图。

如果 isWater[i][j] == 0 ，格子 (i, j) 是一个 陆地 格子。
如果 isWater[i][j] == 1 ，格子 (i, j) 是一个 水域 格子。
你需要按照如下规则给每个单元格安排高度：

每个格子的高度都必须是非负的。
如果一个格子是是 水域 ，那么它的高度必须为 0 。
任意相邻的格子高度差 至多 为 1 。当两个格子在正东、南、西、北方向上相互紧挨着，就称它们为相邻的格子。（也就是说它们有一条公共边）
找到一种安排高度的方案，使得矩阵中的最高高度值 最大 。

请你返回一个大小为 m x n 的整数矩阵 height ，其中 height[i][j] 是格子 (i, j) 的高度。如果有多种解法，请返回任意一个。

示例一：
输入：isWater = [[0,1],[0,0]]
输出：[[1,0],[2,1]]
解释：上图展示了给各个格子安排的高度。
蓝色格子是水域格，绿色格子是陆地格。

示例二：
输入：isWater = [[0,0,1],[1,0,0],[0,0,0]]
输出：[[1,1,0],[0,1,1],[1,2,2]]
解释：所有安排方案中，最高可行高度为 2 。
任意安排方案中，只要最高高度为 2 且符合上述规则的，都为可行方案。

思路：
1. 动态规划
2. BFS

*/

package main

import "fmt"

func main()  {
	isWater:=[][]int{{0,1},{0,0}}
	res:=highestPeak(isWater)
	fmt.Println(res)
}

func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	ans, d := make([][]int, m), [][]int{}
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				d = append(d, []int{i, j})
				ans[i][j] = 0
			} else {
				ans[i][j] = -1
			}
		}
	}
	dirs := [][]int{{1,0}, {-1,0}, {0,1}, {0,-1}}
	for h := 1; len(d) > 0; h++ {
		for size := len(d); size > 0; size--{
			info := d[0]
			d = d[1:]
			x, y := info[0], info[1]
			for _, di := range dirs {
				nx, ny := x + di[0], y + di[1]
				if nx >= 0 && nx < m && ny >= 0 && ny < n && ans[nx][ny] == -1 {
					ans[nx][ny] = h
					d = append(d, []int{nx, ny})
				}
			}
		}
	}
	return ans
}
