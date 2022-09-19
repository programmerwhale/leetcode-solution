package main

import "fmt"

func main(){
	ans:=[]int{11,6,7,4,8,14}
	MergeSort(ans,0,len(ans)-1)
	fmt.Println(ans)
}

// 采用分治法（Divide and Conquer）的一个非常典型的应用。
// 将已有序的子序列合并，得到完全有序的序列；
// 即先使每个子序列有序，再使子序列段间有序。
// 若将两个有序表合并成一个有序表，称为2路归并。
// 时间复杂度O(nlogn）
// 空间复杂度O(n）
func MergeSort(ans []int,start,end int){
	if start>=end{
		return
	}
	mid:=(start+end)/2
	MergeSort(ans,start,mid)
	MergeSort(ans,mid+1,end)
	Merge(ans,start,mid,end)
}

// 将两个有序子序列归并
func Merge(ans []int,start,mid,end int){
	tmparr:=[]int{}
	s1,s2:=start,mid+1
	// 限定边界值
	for s1<=mid&&s2<=end{
		if ans[s1]>ans[s2]{
			tmparr=append(tmparr,ans[s2])
			s2++
		}else{
			tmparr=append(tmparr,ans[s1])
			s1++
		}
	}
	// 其中一个子序列已经排完了之后，其他的直接都加入tmparr
	if s1<=mid{
		tmparr=append(tmparr,ans[s1:mid+1]...)
	}
	if s2<=end{
		tmparr=append(tmparr,ans[s2:end+1]...)
	}
	// 放入最终结果集
	for pos,item:=range tmparr{
		ans[start+pos]=item
	}
}