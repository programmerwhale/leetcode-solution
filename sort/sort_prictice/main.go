package main

import "fmt"

func main()  {
	ans:=[]int{3,4,2,7,6,3}
	quickSort(ans,0,len(ans)-1)
	fmt.Println(ans)
}

// 冒泡排序
// 时间复杂度O(N^2), 空间复杂度O(1)
// 稳定，原地
func BubbleSort(ans []int) {
	for i:=0;i<len(ans);i++{
		for j:=1;j<len(ans)-i;j++{
			if ans[j-1]>ans[j]{
				ans[j-1],ans[j]=ans[j],ans[j-1]
			}
		}
	}
}

// 比较排序
// 时间复杂度O(N^2), 空间复杂度O(1)
// 稳定，原地
func CompareSort(ans []int){
	for i:=0;i<len(ans);i++{
		for j:=len(ans)-1;j>=i+1;j-- {
			if ans[j-1] < ans[j] {
				ans[j], ans[j-1] = ans[j-1], ans[j]
			}
		}
	}
}

// 插入排序
// 时间复杂度O(N^2), 空间复杂度O(1)
// 稳定，原地
func InsertSort(ans []int){
	for i:=1;i<len(ans);i++{
		pre:=i-1
		cur:=ans[i]
		for pre>=0&&ans[pre]>cur{
			ans[pre],ans[pre+1]=ans[pre+1],ans[pre]
			// --，要往前比较
			pre--
		}
	}
}

// 归并排序,先分开排序再合在一起
// 时间复杂度O(nlgn),空间复杂度O(n)
// 稳定，非原地
func mergeSort(ans []int,start,end int){
	if start>=end{
		return
	}
	mid:=start+(end-start)/2
	mergeSort(ans,start,mid)
	mergeSort(ans,mid+1,end)
	merge(ans,start,mid,end)
}
func merge(ans []int,start,mid,end int){
	temp:=[]int{}
	s1,s2:=start,mid+1
	for s1<=mid&&s2<=end{
		if ans[s1]<ans[s2]{
			temp=append(temp,ans[s1])
			s1++
		}
		if ans[s1]>=ans[s2]{
			temp=append(temp,ans[s2])
			s2++
		}
	}
	if s1<mid{
		temp=append(temp,ans[s1:mid+1]...)
	}
	if s2<end{
		temp=append(temp,ans[s2:end+1]...)
	}
	for i,v:=range temp {
		ans[start+i]=v
	}
}

// 快排，根据基准点排序
// 时间复杂度O(nlgn),空间复杂度O(lgn)
// 不稳定，原地
func quickSort(ans []int,start,end int){
	if start>=end{
		return
	}
	mid:=partition(ans,start,end)
	//不包括mid
	quickSort(ans,start,mid-1)
	quickSort(ans,mid+1,end)
}
func partition(ans []int,start,end int)int{
	pivot:=ans[start]
	i,j:=start+1,end
	for true{
		for i<=j&&ans[i]<pivot{
			i++
		}
		for i<=j&&ans[j]>pivot{
			j--
		}
		if i>=j{
			break
		}
		ans[i],ans[j]=ans[j],ans[i]
	}
	ans[start],ans[j]=ans[j],ans[start]
	return j
}