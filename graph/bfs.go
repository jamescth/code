package main

import (
	"fmt"
)

func bfs(start int, nodes map[int][]int) {
	frontier := []int{start}
	explored := map[int]bool{}

	for len(frontier) > 0 {
		node := frontier[0]
		frontier = frontier[1:]
		if _, ok := explored[node]; ok {
			continue
		}

		frontier = append(frontier, nodes[node]...)
		// fmt.Println(frontier)
		//time.Sleep(2 * time.Second)
		explored[node] = true
	}

	fmt.Println(explored)
}

func dfs(start int, nodes map[int][]int) {
	frontier := []int{start}
	visited := map[int]bool{}
}

func dfs_recur(node int, nodes map[int]int, visited map[int]bool) {
	// current_node =
}

var nodes = map[int][]int{
	1:  []int{2, 3, 4},
	2:  []int{1, 5, 6},
	3:  []int{1},
	4:  []int{1, 7, 8},
	5:  []int{2, 9, 10},
	6:  []int{2},
	7:  []int{4, 11, 12},
	8:  []int{4},
	9:  []int{5},
	10: []int{5},
	11: []int{7},
	12: []int{7},
}

func main() {
	visited := []int{}
	bfs(1, nodes)
	fmt.Println(visited)
}
