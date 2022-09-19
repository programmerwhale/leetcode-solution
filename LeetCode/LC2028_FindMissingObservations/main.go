package main

import "fmt"

func main() {
	rolls := []int{4, 2, 2, 5, 4, 5, 4, 5, 3, 3, 6, 1, 2, 4, 2, 1, 6, 5, 4, 2, 3, 4, 2, 3, 3, 5, 4, 1, 4, 4, 5, 3, 6, 1, 5, 2, 3, 3, 6, 1, 6, 4, 1, 3}
	fmt.Println(missingRolls(rolls, 2, 53))
}
func missingRolls(rolls []int, mean int, n int) (res []int) {
	m := len(rolls)
	total := 0
	for _, v := range rolls {
		total += v
	}
	meantotal := (m+n)*mean - total
	//注意不符合的条件
	if meantotal < n || meantotal > 6*n {
		return res
	}
	tmp1 := meantotal / n
	tmp2 := meantotal % n
	for i := 1; i <= n; i++ {
		if i <= tmp2 {
			res = append(res, tmp1+1)
		} else {
			res = append(res, tmp1)
		}
	}
	return
}
