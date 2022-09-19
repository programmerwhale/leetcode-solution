package main

func main()  {

}

func findAnagrams(s string, p string) (ans []int) {
	sLen,pLen:=len(s),len(p)
	if sLen<pLen{
		return
	}
	var sCount,pCount [26]int
	// 构造初始串
	for i,ch:=range p{
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}
	/*for i,ch:=range s[:sLen-pLen]{
		sCount[ch-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount==pCount{
			ans=append(ans,i+1)
		}
	}*/
	for i:=0;i<sLen-pLen;i++{
		sCount[s[i]-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount==pCount{
			ans=append(ans,i+1)
		}
	}
	return
}