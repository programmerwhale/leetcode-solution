package main

import "fmt"

func main() {
	fruits := []int{1, 2, 3, 2, 2}
	fmt.Println(totalFruit(fruits))
}
func totalFruit(fruits []int) int {
	// key 第i种
	// value  棵树
	basket := map[int]int{}
	i := 0
	tmp, res := 0, 0
	for j := 0; j < len(fruits); j++ {
		basket[fruits[j]]++
		tmp++
		if len(basket) == 3 {
			tmp -= basket[i]
			delete(basket, i)
		}
		res = max(res, tmp)
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
