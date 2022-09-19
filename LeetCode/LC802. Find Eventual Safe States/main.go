package main

import "fmt"

func main(){
	graph := [][]int{{1,2},{2,3},{5},{0},{5},{},{}}
	fmt.Println(eventualSafeNodes(graph))
}
func eventualSafeNodes(graph [][]int) []int {
	n:=len(graph)
	// queue bfs需要的队列
	queue:=make([]int,0)
	// topology 入度数
	topology := make([]int, n)
	// edge 出边数组
	edge := make([][]int, n)
	// l:起始点；vv:终点列表
	for l,vv:=range graph{
		for _,v:=range vv{
			// 索引加在
			edge[v]=append(edge[v],l)
		}
		topology[l] = len(vv)
	}
	// 将所有入度为0的加入到对列
	for i, v := range topology {
		if v == 0 {
			queue = append(queue, i)
		}
	}

	// bfs
	for len(queue)!=0{
		q:=queue[0]
		queue=queue[1:]
		for _,ints:=range edge[q]{
			topology[ints]--
			if topology[ints]==0{
				queue=append(queue,ints)
			}
		}
	}
	ans:=[]int{}
	for i,v:=range topology{
		if v==0{
			ans=append(ans,i)
		}
	}
	return ans
}