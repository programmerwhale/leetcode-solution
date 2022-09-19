/*
冒泡排序思路：
把值往后放
时间复杂度0(n^2),空间复杂度O(1)
*/

package main

import "fmt"

func main()  {
	ans:=[]int{3,8,7,3,4,4}
	fmt.Println(ans)
	BubbleSort(ans)
	fmt.Println(ans)
}


func BubbleSort(ans []int){
	for i:=0;i<=len(ans);i++{
		for j:=1;j<len(ans)-i;j++{
			if ans[j]>ans[j-1]{
				ans[j],ans[j-1]=ans[j-1],ans[j]
			}
		}
	}
}