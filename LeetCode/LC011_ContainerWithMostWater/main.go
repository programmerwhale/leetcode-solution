package main

import (
	"fmt"
	"math"
)

func main()  {
	nums:=[]int{1,3,2,5,25,24,5}
	fmt.Println(maxArea(nums))
}
func maxArea(nums []int) int {
	i,j:=0,len(nums)-1
	maxx:=math.MinInt64
	for i<j{
		if (j-i)*min(nums[i],nums[j])>maxx{
			maxx=(j-i)*min(nums[i],nums[j])
		}
		if nums[i]<nums[j]{
			i++
		}else{
			j--
		}
	}
	return maxx
}

func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}