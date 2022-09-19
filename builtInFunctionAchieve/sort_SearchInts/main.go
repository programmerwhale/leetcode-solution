package main

import (
	"fmt"
	"sort"
)

func main()  {
	ans:=[]int{3,34,1,66,78,31}
	fmt.Println(ans)
	fmt.Println("---------------------")
	sort.SearchInts(ans,66)
	fmt.Println("---------------------")
}

// 二分查找找某个数
func searchInt(ans []int, n int){

}
