package main

import "fmt"

func main()  {
	index:=1
	fmt.Scanf("%d\n", &index);
	for i:=1;i<=index;i++{
		n:=0
		fmt.Scanf("%d\n", &n);
		fmt.Println(findWinner(n))
	}
}

func findWinner(index int)int{
	n:=1
	for i:=1;i<=index;i++{
		n*=2
	}
	return n-1
}
