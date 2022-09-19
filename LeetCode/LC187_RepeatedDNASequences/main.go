package main

import "fmt"

func main() {
	fmt.Println(findRepeatedDnaSequences2("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
}

// 方法一：哈希表
// 时间空间复杂度o(NL)
func findRepeatedDnaSequences(s string) (res []string) {
	if len(s) < 10 {
		return res
	}
	mp := map[string]int{}
	// 终止条件i <= len(s)
	// i可以等于0，因为i作为右边区开间
	for i := 10; i <= len(s); i++ {
		tmp := s[i-10 : i]
		p, _ := mp[tmp]
		mp[tmp]++
		if p == 1 {
			res = append(res, tmp)
		}
	}
	return res
}

// 方法二：位运算
// 时间空间复杂度o(N)
const L = 10

var bin = map[byte]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}

func findRepeatedDnaSequences2(s string) (ans []string) {
	n := len(s)
	if n <= L {
		return
	}
	x := 0
	// 构造初始字符串，到下标8，共9位
	for _, ch := range s[:L-1] {
		// bin[ch] 为字符 ch 的对应二进制
		x = x<<2 | bin[byte(ch)]
	}
	cnt := map[int]int{}
	for i := 0; i <= n-L; i++ {
		// 滑动，两位两位滑；再把新的接在后面
		// 只考虑 x的低 20 位比特，需要将其余位置零，即与上 (1 << 20) - 1="00001111"，即后20位不变，高位全位0
		x = (x<<2 | bin[s[i+L-1]]) & (1<<(L*2) - 1)
		cnt[x]++
		if cnt[x] == 2 {
			ans = append(ans, s[i:i+L])
		}
	}
	return ans
}

func findRepeatedDnaSequences3(s string) (ans []string) {
	n := len(s)
	if n <= L {
		return
	}
	x := 0
	// 构造初始字符串，到下标9
	for _, ch := range s[:L] {
		// bin[ch] 为字符 ch 的对应二进制
		x = x<<2 | bin[byte(ch)]
	}
	cnt := map[int]int{}
	cnt[x]++
	for i := L + 1; i <= n; i++ {
		// 滑动，两位两位滑；再把新的接在后面
		// 只考虑 x的低 20 位比特，需要将其余位置零，即与上 (1 << 20) - 1="00001111"，即后20位不变，高位全位0
		x = (x<<2 | bin[s[i-1]]) & (1<<(L*2) - 1)
		cnt[x]++
		if cnt[x] == 2 {
			ans = append(ans, s[i-L:i])
		}
	}
	return ans
}
