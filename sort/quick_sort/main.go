/*
它的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
*/
package main

import "fmt"

func main()  {
	ans:=[]int{9,8,10,7,2,15}
	fmt.Println(ans)
	QuickSort(ans,0,len(ans)-1)
	fmt.Println(ans)
}

/*
与归并类似，分治思想；
限定好左右界限大小关系，先找基准线，后递归

*/
func QuickSort(ans []int,left,right int)  {
	// 限定好左右索引关系
	if left<right{
		// 以基准线为界
		mid:=partition(ans,left,right)
		QuickSort(ans,left,mid-1)
		QuickSort(ans,mid+1,right)
	}
}

func partition(ans []int,left,right int)int{
	//选中基准元素，可以随机选
	pivot:=ans[left]
	i,j:=left+1,right
	for true{
		// 向右找到第一个大于 pivot 的元素位置
		for i<=j&&ans[i]<=pivot{
			i++
		}
		// 向左找到第一个小于 pivot 的元素位置
		for i<=j&&ans[j]>=pivot{
			j--
		}
		// 交叉了，跳出循环
		if i>=j{
			break
		}
		//交换两个元素的位置，使得左边的元素不大于pivot,右边的不小于pivot
		ans[i],ans[j]=ans[j],ans[i]
	}

	//把ans[j]和主元交换（i会破坏顺序）
	ans[left]=ans[j]
	ans[j]=pivot
	// 下一个基准元素
	return j
}