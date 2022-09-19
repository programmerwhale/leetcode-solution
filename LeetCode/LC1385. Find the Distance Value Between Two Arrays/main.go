package main

import "sort"

func main()  {

}

// 模拟
// 时间复杂度O(M*N),空间复杂度0(1)
func findTheDistanceValue2(arr1 []int, arr2 []int, d int) int {
	res:=0
	for _,num1:=range arr1{
		for i,num2:=range arr2{
			if abs(num1-num2)<=d{
				break
			}
			if i==len(arr2)-1{
				res++
			}
		}
	}
	return res
}

// 二分法
// 时间复杂度O((N+M)logN),空间复杂度0(1)
// 可以理解为：在arr2中找与arr[i]最相近的数，如果连最相近的数绝对值差都大于d，那么其他就更不用考虑了
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	n:=len(arr2)
	res:=0
	for _,num1:=range arr1{
		if num1<=arr2[0]&&arr2[0]-num1>d||num1>=arr2[n-1]&&num1-arr2[n-1]>d{
			res++
		}else{
			i,j:=0,n-1
			for i<j{
				mid:=(i+j)/2
				if arr2[mid]==num1{
					i=mid
					break
				}else if arr2[mid]<num1{
					i=mid+1
				}else{
					// 为啥不-1
					j=mid
				}
			}
			if abs(arr2[i]-num1)>d&&i>0&&abs(arr2[i-1]-num1)>d{
				res++
			}
		}
	}
	return res
}

func abs(a int)int{
	if a<0{
		return -a
	}
	return a
}