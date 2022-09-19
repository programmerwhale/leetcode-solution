package main
import "fmt"

func main()  {
	index:=1
	fmt.Scanf("%d\n", &index);
	for i:=1;i<=index;i++{
		n:=0
		fmt.Scanf("%d\n", &n);
		ans:=[]int{}
		for j:=1;j<=n;j++{
			x:=1
			fmt.Scanf("%d", &x);
			ans=append(ans,x)
		}
		xor(ans)
	}
}

func xor(ans []int){
	res:=ans[0]
	for i:=1;i<len(ans);i++{

		fmt.Println(res)
		res^=ans[i]
	}
	fmt.Println(res)
}