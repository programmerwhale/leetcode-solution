package main

import (
	"fmt"
	"sort"
)

func main() {
	g := []int{1, 2, 3}
	s := []int{1, 1}
	fmt.Println(findContentChildren(g, s))
}

func findContentChildren(g []int, s []int) (cnt int) {
	sort.Ints(g)
	sort.Ints(s)
	tmp := 0
	for _, v := range g {
		for i := tmp; i < len(s); i++ {
			if v <= s[i] {
				cnt++
				tmp = i + 1
			}
		}
	}
	return cnt
}
