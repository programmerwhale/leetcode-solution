package main

import (
	"fmt"
)

func main() {
	s := "abcdefg"
	fmt.Println(reverseLeftWords(s, 2))
}
func reverseLeftWords(s string, n int) string {
	length := len(s)
	bytes := []byte(s)
	for i := 0; i < n/2; i++ {
		bytes[i], bytes[n-i-1] = bytes[n-i-1], bytes[i]
	}
	for i := n; i < (n+length)/2; i++ {
		bytes[i], bytes[n+length-1-i] = bytes[n+length-1-i], bytes[i]
	}
	for i := 0; i < length/2; i++ {
		bytes[i], bytes[length-1-i] = bytes[length-1-i], bytes[i]
	}
	return string(bytes)
}
