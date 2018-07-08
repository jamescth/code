package main

import "fmt"

func findDigits(n int32) int32 {
	var darray []int32
	t := n
	for t != 0 {
		d := t % 10
		darray = append(darray, d)
		t /= 10
	}

	fmt.Printf("%v\n", darray)

	var ret int32
	for _, d := range darray {
		if d == 0 {
			continue
		}
		if n%d == 0 {
			ret++
		}
	}
	return ret
}
