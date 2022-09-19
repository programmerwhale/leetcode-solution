package main

import "fmt"

func main()  {
	ans:=[]int{9,8,10,7,2,15}
	fmt.Println(ans)
	InsertSort(ans)
	fmt.Println(ans)
}

/*插入排序
1. 外循环为交换次数，从第二个元素开始往前比较，所以下标从1开始
2. 设置一个前置下标，即当前结点前一个
3. 内循环处理当前结点及之前的元素，注意pre的边界处理
4. 如果前置结点比当前结点大，交换，再--前置元素下标，两两交换，为了不破坏以前的有序状态；否则按兵不动；*/
func InsertSort(ans []int){
	for i:=1;i<len(ans);i++{
		pre:=i-1
		cur:=ans[i]
		for pre>=0&&ans[pre]<cur{
			ans[pre+1],ans[pre]=ans[pre],ans[pre+1]
			pre--
		}
	}
}
