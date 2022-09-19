package main

import "fmt"

func main() {
	strs := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
}

// 方法一：纵向
func longestCommonPrefix(strs []string) string {
	for i := 0; i < len(strs[0]); i++ {
		tmp := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != tmp {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
