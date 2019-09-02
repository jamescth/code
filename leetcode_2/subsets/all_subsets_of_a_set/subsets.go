package main

import "fmt"

func subsets(nums []int) [][]int {
	res := [][]int{[]int{}}
	for _, e := range nums {
		fmt.Printf("e %d\n", e)
		for _, ee := range res {
			fmt.Printf("ee %v\n", ee)
			res = append(res, append([]int{e}, ee...))
		}
		fmt.Printf("res %v\n", res)
	}
	return res
}
