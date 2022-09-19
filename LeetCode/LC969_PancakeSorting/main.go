package main

import "fmt"

func main() {
	arr := []int{3, 2, 4, 1}
	fmt.Println(pancakeSort(arr))
}

// 方法一：模拟法
// 重要思路：把后面的数字先排好序，这样再翻转前面的时候就不会影响到后面。
// （我不理解）
func pancakeSort(arr []int) (ans []int) {
	for n := len(arr); n > 1; n-- {
		index := 0
		for i, v := range arr[:n] {
			if v > arr[index] {
				index = i
			}
		}
		if index == n-1 {
			continue
		}
		for i, m := 0, index; i < (m+1)/2; i++ {
			arr[i], arr[m-i] = arr[m-i], arr[i]
		}
		for i := 0; i < n/2; i++ {
			arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
		}
		ans = append(ans, index+1, n)
	}
	return
}
