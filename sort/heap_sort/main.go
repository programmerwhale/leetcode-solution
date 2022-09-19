package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	var length = 10
	var arr []int
	for i := 0; i < length; i++ {
		arr = append(arr, int(rand.Intn(1000)))
	}
	fmt.Println(arr)
	heapSort(arr)
	fmt.Println(arr)
}
// 堆排序的时间复杂度是O(N*logN)。
func heapSort(arr []int){
	n:=len(arr)
	//1.构建大顶堆
	for i:=n/2-1;i>=0;i--{
		//从第一个非叶子结点从下至上，从右至左调整结构
		adjustHeap(arr,i,n)
	}
	//2.调整堆结构+交换堆顶元素与末尾元素
	for j:=n/2-1;j>0;j-- {
		//将堆顶元素与末尾元素进行交换
		arr[0],arr[j]=arr[j],arr[0]
		//重新对堆进行调整
		adjustHeap(arr,0,j)
	}
}
// 调整大顶堆（仅是调整过程，建立在大顶堆已构建的基础上）
func adjustHeap(arr []int,i,length int){
	//先取出当前元素i
	temp := arr[i]
	for k:=i*2+1;k<length;k=k*2+1{//从i结点的左子结点开始，也就是2i+1处开始
		if k+1<length&&arr[k]>arr[k+1]{//如果左子结点小于右子结点，k指向右子结点
			k++
		}
		if arr[k]>temp{//如果子节点大于父节点，将子节点值赋给父节点（不用进行交换）
			arr[i]=arr[k]
			i=k
		}else {
			break
		}
	}
	arr[i]=temp//将temp值放到最终的位置
}