package main

import "fmt"

func main()  {
	ans:=[]int{1}
	fmt.Println(searchRange(ans,1))
}
func searchRange(nums []int, target int) []int {
	var search func([]int,int,int,int)int
	search=func(nums []int,left,right,target int)int{
		for left<=right{
			mid:=(left+right)>>1
			if nums[mid]<target{
				left=mid+1
			}else if nums[mid]>=target{
				right=mid-1
			}
		}
		return left
	}
	first:=search(nums,0,len(nums)-1,target)
	last:=search(nums,0,len(nums)-1,target+1)-1
	// 没找到的条件
	if first<=last&&nums[first]==target&&nums[last]==target{
		return []int{first,last}
	}
	return []int{-1,-1}
}