package main

import (
	"fmt"
	"math"
)

func main()  {
	nums:=[]int{4,1,7,9}
	fmt.Println(nums)
	quickSort(nums,0,len(nums)-1)
	fmt.Println(nums)
	fmt.Println(minimumDifference(nums,3))
}
func minimumDifference(nums []int, k int) int {
	if k==1{
		return 0
	}
	//mergeSort(nums,0,len(nums)-1)
	quickSort(nums,0,len(nums)-1)
	minn:=math.MaxInt64
	for i:=k-1;i<len(nums);i++{
		if nums[i]-nums[i-k+1]<minn{
			minn=nums[i]-nums[i-k+1]
		}
	}
	return minn
}

func quickSort(nums []int,start,end int){
	if start>=end{
		return
	}
	mid:=partition(nums,start,end)
	quickSort(nums,start,mid-1)
	quickSort(nums,mid+1,end)
}

func partition(nums []int,start,end int)int{
	pivot:=nums[start]
	l,r:=start+1,end
	for true{
		for l<=r&&nums[l]<=pivot{
			l++
		}
		for l<=r&&nums[r]>=pivot{
			r--
		}
		if l>=r{
			break
		}
		nums[l],nums[r]=nums[r],nums[l]
	}
	nums[start],nums[r]=nums[r],nums[start]
	return r
}

func mergeSort(nums []int,start,end int)  {
	if start>=end{
		return
	}
	mid:=(start+end)>>1
	mergeSort(nums,start,mid)
	mergeSort(nums,mid+1,end)
	merge(nums,start,mid,end)
}
func merge(nums []int,start,mid,end int){
	s1,s2:=start,mid+1
	temp:=[]int{}
	for s1<=mid&&s2<=end{
		if nums[s1]<nums[s2]{
			temp=append(temp,nums[s1])
			s1++
		}
		if nums[s2]<nums[s1]{
			temp=append(temp,nums[s2])
			s2++
		}
	}
	if s1<=mid{
		temp=append(temp,nums[s1:mid+1]...)
	}
	if s2<=end{
		temp=append(temp,nums[s2:end+1]...)
	}
	for i,v:=range temp{
		nums[start+i]=v
	}
}