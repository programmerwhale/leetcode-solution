package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(solveNQueens(4))
}
var res [][]string
func solveNQueens(n int) [][]string {
	// 初始化
	res=[][]string{}
	board:=make([][]string,n)
	for i:=0;i<n;i++{
		board[i]=make([]string,n)
	}
	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
			board[i][j]="."
		}
	}
	//回溯,从0开始
	backtracking(board,0)
	return res
}

func backtracking(board [][]string,row int){
	size:=len(board)
	if row==size{
		temp:=make([]string,size)
		for i:=0;i<size;i++{
			// 变string
			temp[i]=strings.Join(board[i],"")
		}
		res=append(res,temp)
		// 这个return啊！！
		return
	}
	for col:=0;col<size;col++{
		if !isValid(board, row, col){
			continue
		}
		board[row][col] = "Q"
		backtracking(board, row+1)
		board[row][col] = "."
	}
}

func isValid(board [][]string,row,col int)bool{
	n:=len(board)
	// 同一列是否有皇后
	for i:=0;i<row;i++{
		if board[i][col]=="Q"{
			return false
		}
	}
	// 同一行是否有皇后
	for i:=0;i<col;i++{
		if board[row][i]=="Q"{
			return false
		}
	}
	// 两个对角线
	for i,j:=row,col;i>=0&&j>=0;i,j=i-1,j-1{
		if board[i][j]=="Q"{
			return false
		}
	}
	for i,j:=row,col;i>=0&&j<n;i,j=i-1,j+1{
		if board[i][j]=="Q"{
			return false
		}
	}
	return true
}