package main

import (
	"fmt"
	"math"
)

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}
func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for _, ch := range t {
		ori[byte(ch)]++
	}
	slen := len(s)
	length := math.MaxInt64
	ansL, ansR := -1, -1
	l, r := 0, 0
	for l <= r && r < slen {
		if ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check(ori, cnt) && l <= r {
			if r-l+1 < length {
				length = r - l + 1
				ansL = l
				ansR = r + 1
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]]--
			}
			l++
		}
		r++
	}
	return s[ansL:ansR]
}

//  满足魔术串返回true
func check(ori, cnt map[byte]int) bool {
	for k, v := range ori {
		if cnt[k] < v {
			return false
		}
	}
	return true
}
