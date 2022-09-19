package main

func main()  {
	
}
// 快慢指针
// 时间复杂度O(LogN),空间复杂度O(1)
func isHappy(n int) bool {
	slow,fast:=n,next(n)
	for slow!=1&&fast!=1{
		if slow==fast{
			return false
		}
		slow,fast=next(slow),next(next(fast))
	}
	return true
}
func next(n int)int{
	sum:=0
	for n>0{
		sum+=(n%10)*(n%10)
		n/=10
	}
	return sum
}