package main

func main()  {
	
}
func letterCombinations(digits string) (res []string) {
	if len(digits)==0{
		return
	}
	digitsMap:= [10]string{
		"", // 0
		"", // 1
		"abc", // 2
		"def", // 3
		"ghi", // 4
		"jkl", // 5
		"mno", // 6
		"pqrs", // 7
		"tuv", // 8
		"wxyz", // 9
	}
	backtracking(digitsMap,digits,0,&res,"")
	return
}
func backtracking(digitsMap [10]string,digits string,index int,res *[]string,str string){
	if len(str)==len(digits){
		*res=append(*res,str)
		return
	}
	letter:=digitsMap[(digits[index]-'0')]
	for i:=0;i<len(letter);i++{
		str+=string(letter[i])
		backtracking(digitsMap,digits,index+1,res,str)
		str=str[:len(str)-1]
	}
}