package main

import (
	"strconv"
)

func main() {

}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 递归 n n
/*func binaryTreePaths(root *TreeNode) (result []string) {
	traversal(root,&result,"")
	return
}
func traversal(root *TreeNode, res *[]string, path string){
	// 先添加子串
	if len(path)!=0{
		path=path+"->"+strconv.Itoa(root.Val)
	}else{
		path=strconv.Itoa(root.Val)
	}

	// 加入res并中止
	if root.Left==nil&&root.Right==nil{
		if root.Left==nil&&root.Right==nil{
			*res=append(*res,path)
			return
		}
	}

	// 继续递归
	if root.Left!=nil{
		traversal(root.Left,res,path)
	}
	if root.Right!=nil{
		traversal(root.Right,res,path)
	}
}*/
// 回溯
func binaryTreePaths(root *TreeNode) (result []string) {
	var path []int
	traversal(root, &result, &path)
	return
}

func traversal(root *TreeNode, result *[]string, path *[]int) {
	*path = append(*path, root.Val)
	//判断是否为叶子节点
	if root.Left == nil && root.Right == nil {
		pathStr := strconv.Itoa((*path)[0])
		for i := 1; i < len(*path); i++ {
			pathStr = pathStr + "->" + strconv.Itoa((*path)[i])
		}
		*result = append(*result, pathStr)
		return
	}
	//左右
	if root.Left != nil {
		traversal(root.Left, result, path)
		*path = (*path)[:len(*path)-1] //回溯到上一个节点（因为traversal会加下一个节点值到path中）
	}
	if root.Right != nil {
		traversal(root.Right, result, path)
		*path = (*path)[:len(*path)-1] //回溯
	}
}
