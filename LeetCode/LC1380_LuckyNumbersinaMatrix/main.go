package main

import (
	"fmt"
	"math"
)

func main()  {
	ans:=[][]int{{3,7,8},{9,11,13},{15,16,17}}
	fmt.Println(luckyNumbers(ans))
}

func luckyNumbers (matrix [][]int) (res []int) {
	m,n:=len(matrix),len(matrix[0])
	minlist:=[]int{}
	maxlist:=[]int{}
	for i:=0;i<m;i++{
		minn:=math.MaxInt64
		for j:=0;j<n;j++{
			minn=min(minn,matrix[i][j])
		}
		minlist=append(minlist,minn)
	}

	for i:=0;i<n;i++{
		maxx:=math.MinInt64
		for j:=0;j<m;j++{
			maxx=max(maxx,matrix[j][i])
		}
		maxlist=append(maxlist,maxx)
	}

	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if matrix[i][j]==minlist[i]&&matrix[i][j]==maxlist[j]{
				res=append(res,matrix[i][j])
			}
		}
	}

	return res
}
func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}
