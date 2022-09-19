package main

import "fmt"

func main() {
	img := [][]int{{100, 200, 100}, {200, 50, 200}, {100, 200, 100}}
	fmt.Println(imageSmoother(img))
}
func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	res := make([][]int, m)
	for i, row := range img {
		res[i] = make([]int, n)
		for j, _ := range row {
			sum := 0
			cnt := 0
			for _, ans := range img[max(i-1, 0):min(i+2, m)] {
				for _, v := range ans[max(j-1, 0):min(j+2, n)] {
					sum += v
					cnt++
				}
			}
			res[i][j] = sum / cnt
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
