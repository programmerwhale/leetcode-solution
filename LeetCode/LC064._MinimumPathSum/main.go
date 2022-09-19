package main

import "fmt"

func main()  {
	grid:=[][]int{{1,3,1},{1,5,1},{4,2,1}}
	fmt.Println(minPathSum(grid))
}

type pair struct {
	x,y int
}

// 原地修改
func minPathSum(grid [][]int) int {
	m,n:=len(grid),len(grid[0])
	for i:=1;i<n;i++{
		grid[0][i]+=grid[0][i-1]
	}
	for i:=1;i<m;i++{
		grid[i][0]+=grid[i-1][0]
	}
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			grid[i][j]+=min(grid[i-1][j],grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}
func min(a,b int) int {
	if a<b{
		return a
	}
	return b
}