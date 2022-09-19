package main

import (
	"fmt"
	"sort"
)

type pair struct {
	target  string
	visited bool
}
type pairs []*pair

func main()  {
	tickets:=[][]string{{"JFK","SFO"},{"JFK","ATL"},{"SFO","ATL"},{"ATL","JFK"},{"ATL","SFO"}}
	fmt.Println(findItinerary(tickets))
}

func findItinerary(tickets [][]string) (result []string) {
	targets:=make(map[string]pairs)
	for _,ticket:=range tickets{
		if targets[ticket[0]] == nil {
			targets[ticket[0]] = make(pairs, 0)
		}
		targets[ticket[0]] = append(targets[ticket[0]], &pair{target: ticket[1], visited: false})
	}
	//fmt.Println(targets)
	// 按字典序排序
	for k, _ := range targets {
		sort.Slice(targets[k],func(i,j int)bool{return targets[k][i].target<targets[k][j].target})
	}
	fmt.Println(targets)

	result=append(result,"JFK")
	backtracing(len(tickets),&result,targets)
	return result
}

func backtracing(ticketNum int,result *[]string,targets map[string]pairs)bool{
	if len(*result)==ticketNum+1{
		return true
	}

	for _,pair:=range targets[(*result)[len(*result)-1]]{
		if !pair.visited{
			*result=append(*result,pair.target)
			pair.visited=true
			if backtracing(ticketNum,result,targets){
				return true
			}
			*result=(*result)[:len(*result)-1]
			pair.visited=false
		}
	}

	return false
}