package main

import (
	"fmt"
	"math/bits"
)

func main()  {
	request:=[][]int{{0,1},{1,0},{0,1},{1,2},{2,0},{3,4}}
	fmt.Println(maximumRequests(5,request))
}
func maximumRequests(n int, requests [][]int) (ans int) {
next:
	for mask := 0; mask < 1<<len(requests); mask++ {
		cnt := bits.OnesCount(uint(mask))
		if cnt <= ans {
			continue
		}
		delta := make([]int, n)
		for i, r := range requests {
			if mask>>i&1 == 1 {
				delta[r[0]]++
				delta[r[1]]--
			}
		}
		for _, d := range delta {
			if d != 0 {
				continue next
			}
		}
		ans = cnt
	}
	return
}