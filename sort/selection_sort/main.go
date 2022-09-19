/*
选择排序思路：
遍历找到最大/小的值
和第一遍定位的值交换
时间复杂度0(n^2),空间复杂度O(1)
*/


package main

import (
	"fmt"
)

func main()  {
	ans:=[]int{9,8,10,7,2,15}
	fmt.Println(ans)
	SelectionSort2(ans)
	fmt.Println(ans)
}

// 升序
func SelectionSort(ans []int){
	for i:=0;i<len(ans)-1;i++{
		minindex:=i
		for j:=i+1;j<len(ans);j++{
			if ans[j]<ans[minindex]{
				minindex=j
			}
		}
		ans[i],ans[minindex]=ans[minindex],ans[i]
	}
}

// 降序
func SelectionSort2(ans []int){
	for i:=0;i<len(ans)-1;i++{
		maxindex:=i
		for j:=i+1;j<len(ans);j++{
			if ans[j]>ans[maxindex]{
				maxindex=j
			}
		}
		ans[i],ans[maxindex]=ans[maxindex],ans[i]
	}
}

