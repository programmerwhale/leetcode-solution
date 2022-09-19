package main

import "fmt"

func main()  {
	fmt.Println(findKthNumber(15,2))
}

func getCount(prefix int,n int)int{
	cnt:=0
	for a,b:=prefix,prefix+1;a<=n;a,b=a*10,b*10{
		cnt+=min(n+1,b)-a
	}
	return cnt
}

func findKthNumber(n int, k int) int {
	prefix:=1
	p:=1
	for p<k{
		cnt:=getCount(prefix,n)
		//如果超出范围，扩大前缀，p++
		if p+cnt>k{
			prefix*=10
			// 记得p++，这时候已经相当于把prefix算进去了
			p++
		}else{
			// 否则继续循环
			prefix++
			p+=cnt
		}
	}
	return prefix
}

func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}