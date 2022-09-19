package main

import "fmt"

func main()  {
	fmt.Println(sumNums2(3))
}
func sumNums(n int) int {
	if n==1{
		return 1
	}
	temp:=1+n
	times:=float64(n)/2
	return int(float64(temp)*times)
}
func sumNums2(n int) int {
	res:=0
	var sum func(int)
	sum=func(i int){
		if i<=0{
			return
		}
		res+=i
		sum(i-1)
	}
	sum(n)
	return res
}