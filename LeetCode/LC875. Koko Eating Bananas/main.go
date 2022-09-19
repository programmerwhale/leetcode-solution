package main

func main()  {
	
}
// 二分查找
func minEatingSpeed(piles []int, h int) int {
	maxx:=0
	for _,pile:=range piles{
		if pile>maxx{
			maxx=pile
		}
	}
	l,r:=1,maxx
	//fmt.Println(r)
	for l<r{
		mid:=l+(r-l)/2
		sum:=calculate(piles,mid)
		if sum>h{
			l=mid+1
		}else{
			r=mid
		}
	}
	return l
}
func calculate(piles []int,k int)(res int){
	for _,pile:=range piles{
		res+=(pile+k-1)/k
	}
	//fmt.Println(res)
	return
}