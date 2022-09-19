package main

func main()  {

}

func mySqrt(x int) int {
	left,right:=0,x
	for left<right{
		mid:=(left+right)>>1
		if mid*mid<=x{
			left=mid+1
		}else{
			right=mid-1
		}
	}
	// 在最后的时候判断一下
	// 如果这个整数的平方 严格小于 输入整数，那么这个整数 可能 是我们要找的那个数。
	if left*left>x{
		return left-1
	}
	return left
}
1 2 3 4 5