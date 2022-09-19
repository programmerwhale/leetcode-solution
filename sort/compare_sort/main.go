
/*
思路：
两层循环：外层是元素个数，循环n次；内层是相邻两个值的比较
如果是升序，从第一个元素开始找最小的，如果前面的比后面大，则交换（把小的换到最前面）
如果是降序，从最后一个元素开始找最大的，如果后面的比前面的大，则交换
时间复杂度0(n^2),空间复杂度O(1)
*/
package main

import "fmt"

func main(){
	ans:=[]int{9,8,10,7,2,15}
	fmt.Println(ans)
	CompareSort(ans)
	fmt.Println(ans)
}

func  CompareSort(ans []int){
	for i:=0;i<=len(ans);i++ {
		for j:=len(ans)-1;j>=i+1;j--{
			if ans[j-1]<ans[j]{
				ans[j],ans[j-1]=ans[j-1],ans[j]
			}
		}
	}
}
