package main

func main()  {
	
}

// 字典树
func lexicalOrder(n int) (res []int){
	res=make([]int,n)
	prefix:=1
	for i:=range res{
		res[i]=prefix
		if prefix*10<=n{
			prefix*=10
		}else{
			for prefix%10==9||prefix+1>n{
				prefix/=10
			}
			prefix++
		}
	}
	return
}