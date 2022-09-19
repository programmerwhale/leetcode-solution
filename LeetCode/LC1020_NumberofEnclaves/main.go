package main

import (
	"fmt"
)

/*
给你一个大小为 m x n 的二进制矩阵 grid ，其中 0 表示一个海洋单元格、1 表示一个陆地单元格。

一次 移动 是指从一个陆地单元格走到另一个相邻（上、下、左、右）的陆地单元格或跨过 grid 的边界。

返回网格中 无法 在任意次数的移动中离开网格边界的陆地单元格的数量。



示例 1：


输入：grid = [[0,0,0,0],[1,0,1,0],[0,1,1,0],[0,0,0,0]]
输出：3
解释：有三个 1 被 0 包围。一个 1 没有被包围，因为它在边界上。
示例 2：


输入：grid = [[0,1,1,0],[0,0,1,0],[0,0,1,0],[0,0,0,0]]
输出：0
解释：所有 1 都在边界上或可以到达边界。


提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 500
grid[i][j] 的值为 0 或 1
*/
func main() {
	grid := [][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}
	fmt.Println(numEnclaves3(grid))
}

// 方法一：深度优先搜索dfs
// 从边界的陆地往内查找并标记，最后遍历没有查找标记的陆地即为所求
// 时间复杂度o(mn),空间复杂度o(mn)
type pair struct {
	x, y int
}

func numEnclaves(grid [][]int) (ans int) {
	// 往四个方向搜索
	dirs := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(grid), len(grid[0])
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}
	//深度优先搜索
	var dfs func(int, int)
	dfs = func(r, c int) {
		// 超出边界或者为海洋或者已被搜索过
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == 0 || vis[r][c] {
			return
		}
		// 已搜索
		vis[r][c] = true
		// 四周发散
		for _, d := range dirs {
			dfs(r+d.x, c+d.y)
		}
	}
	// 从边界开始搜索，最左一列和最右一列
	for i := range grid {
		dfs(i, 0)
		dfs(i, n-1)
	}
	// 出去四周四个点，搜索上下两行
	for j := 1; j < n-1; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}
	// 结束后遍历grid[i][j] == 1 但没有被搜索过但点的数量
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 1 && !vis[i][j] {
				ans++
			}
		}
	}
	return
}

// 方法二：广度优先搜索bfs
// 从边界的陆地往内查找并标记，最后遍历没有查找标记的陆地即为所求
// 时间复杂度o(mn),空间复杂度o(mn)
func numEnclaves2(grid [][]int) (ans int) {
	// 往四个方向搜索
	dirs := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(grid), len(grid[0])
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}
	q := []pair{}
	for i, row := range grid {
		if row[0] == 1 {
			vis[i][0] = true
			q = append(q, pair{i, 0})
		}
		if row[n-1] == 1 {
			vis[i][n-1] = true
			q = append(q, pair{i, n - 1})
		}
	}
	for i := 1; i < n-1; i++ {
		if grid[0][i] == 1 {
			vis[0][i] = true
			q = append(q, pair{0, i})
		}
		if grid[m-1][i] == 1 {
			vis[m-1][i] = true
			q = append(q, pair{m - 1, i})
		}
	}
	for len(q) > 0 {
		// 取队列头
		p := q[0]
		q = q[1:]
		for _, d := range dirs {
			if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == 1 && !vis[x][y] {
				vis[x][y] = true
				q = append(q, pair{x, y})
			}
		}
	}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 1 && !vis[i][j] {
				ans++
			}
		}
	}
	return
}

// 方法三：并查集( 不会)
//时间复杂度O(mn×α(mn)),空间复杂度o(mn)
type unionField struct {
	parent []int
	rank   []int
	onEdge []bool
}

func newUnionFind(grid [][]int) unionField {
	m, n := len(grid), len(grid[0])
	parent := make([]int, m*n)
	rank := make([]int, m*n)
	onEdge := make([]bool, m*n)
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				idx := i*n + j
				parent[idx] = idx
				if i == 0 || i == m-1 || j == 0 || j == n-1 {
					onEdge[idx] = true
				}
			}
		}
	}
	return unionField{parent, rank, onEdge}
}

func (uf unionField) merge(x, y int) {
	x, y = uf.find(x), uf.find(y)
	if x == y {
		return
	}
	if uf.rank[x] > uf.rank[y] {
		uf.parent[y] = x
		uf.onEdge[x] = uf.onEdge[x] || uf.onEdge[y]
	} else if uf.rank[x] < uf.rank[y] {
		uf.parent[x] = y
		uf.onEdge[y] = uf.onEdge[y] || uf.onEdge[x]
	} else {
		uf.parent[y] = x
		uf.onEdge[x] = uf.onEdge[x] || uf.onEdge[y]
		uf.rank[x]++
	}
}

func (uf unionField) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func numEnclaves3(grid [][]int) (ans int) {
	uf := newUnionFind(grid)
	m, n := len(grid), len(grid[0])
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				idx := i*n + j
				if j+1 < n && grid[i][j+1] == 1 {
					uf.merge(idx, idx+1)
				}
				if i+1 < m && grid[i+1][j] == 1 {
					uf.merge(idx, idx+n)
				}
			}
		}
	}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 1 && !uf.onEdge[uf.find(i*n+j)] {
				ans++
			}
		}
	}
	return
}
