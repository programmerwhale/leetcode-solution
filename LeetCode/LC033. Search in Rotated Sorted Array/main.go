package main

func main()  {

}

func search(nums []int, target int) int {
	l,r:=0,len(nums)-1
	for l<=r{
		mid:=l+(r-l)/2
		if nums[mid]==target{
			return mid
		}
		if nums[l]<nums[mid]{
			// 前半段有序
			if nums[l]<=target&&target<nums[mid]{
				r=mid-1
			}else{
				l=mid+1
			}
		}else{
			// 后半段有序
			if nums[mid]<target&&target<=nums[r]{
				l=mid+1
			}else{
				r=mid-1
			}
		}
	}
	return -1
}