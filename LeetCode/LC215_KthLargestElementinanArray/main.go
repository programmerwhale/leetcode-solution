package main
/*
215. 数组中的第K个最大元素
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4


提示：

1 <= k <= nums.length <= 104
-104 <= nums[i] <= 104
*/
import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	arr:=[]int{3,2,1,5,6,4}
	fmt.Println(findKthLargest(arr,2))
}

// 时间复杂度O(Nlgn)
// 主要考察的是快排和堆排
func findKthLargest(nums []int, k int) int {
	quickSort(nums,0,len(nums)-1)
	return nums[len(nums)-k]
}
func quickSort(nums []int,left,right int)  {
	if left>=right{
		return
	}
	mid:=partition(nums,left,right)
	// 把mid空出来
	quickSort(nums,left,mid-1)
	quickSort(nums,mid+1,right)
}
func partition(nums []int,left,right int) (mid int) {
	pivot:=nums[left]
	i,j:=left+1,right
	for {
		for i<=j&&nums[i]<=pivot{
			i++
		}
		for i<=j&&nums[j]>=pivot{
			j--
		}
		if i>=j{
			break
		}
		nums[i],nums[j]=nums[j],nums[i]
	}
	nums[left],nums[j]=nums[j],nums[left]
	return j
}

// 只找下标，前后是否有序不关心
// 时间复杂度O(N),空间复杂度O(Nlgn)
func findKthLargest2(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(a []int, l, r, index int) int {
	q := randomPartition(a, l, r)
	// 找到下标
	if q == index {
		return a[q]
	} else if q < index {
		return quickSelect(a, q + 1, r, index)
	}
	return quickSelect(a, l, q - 1, index)
}

func randomPartition(a []int, l, r int) int {
	i := rand.Int() % (r - l + 1) + l
	a[i], a[r] = a[r], a[i]
	return partition2(a, l, r)
}

func partition2(a []int, l, r int) int {
	x := a[r]
	i := l - 1
	for j := l; j < r; j++ {
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

// 堆排
// 时间复杂度O(Nlgn),空间复杂度O(lgn)
func findKthLargest3(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	// 删掉k个之后再重新构建，此时堆顶为结果
	for i := len(nums) - 1; i > len(nums) - k ; i-- {
		// 头尾交换
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		// 固定从下标0开始
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

// 构建最大堆
func buildMaxHeap(a []int, heapSize int) {
	//第一次建立大根堆，从后往前依次调整
	for i := heapSize/2; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	// l,r是当前下标下一层的左右下标
	l, r, largest := i * 2 + 1, i * 2 + 2, i
	// 先看是不是在范围内
	// 如果左右叶子节点比当前的大，则交换
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	// 如果largest没变就不用换了，变了就需要交换位置，重新构建堆
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapSize)
	}
}
