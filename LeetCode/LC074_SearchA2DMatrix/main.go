package main

import "fmt"

func main()  {
	matrix:=[][]int{{1,3,5,7},{10,11,16,20},{20,30,34,60}}
	fmt.Println(matrix)
	fmt.Println(searchMatrix(matrix,13))
}

func searchMatrix(matrix [][]int, target int) bool {
	m,n:=len(matrix),len(matrix[0])
	i,j:=m-1,0
	for i>=0&&j<n{
		if matrix[i][j]==target{
			return true
		}else if j<n&&matrix[i][j]<target{
			j++
		}else {
			i--
		}
	}
	return false
}