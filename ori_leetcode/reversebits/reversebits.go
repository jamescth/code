package main

import (
	"fmt"
	"strconv"
)

func reverseBits(n int) int {
	tmp := ""
	b := 32
	for b > 0 {
		d := n & 1
		tmp = tmp + strconv.Itoa(d)
		n = n >> 1
		b--
	}
	fmt.Printf("%s\n", tmp)

	i := len(tmp) - 1
	ret := 0
	var exp uint
	for i >= 0 {
		d := tmp[i]
		v, _ := strconv.Atoi(string(d))
		ret += v * (1 << exp)
		i--
		exp++
	}
	return ret
}
