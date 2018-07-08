package main

import "fmt"

func intToRoman(num int) string {

	var ret string

	fn := func(n int, b byte, s string) string {
		for i := 0; i < n; i++ {
			s = s + string(b)
		}

		return s
	}

	t := num / 1000
	if t > 0 {
		ret = fn(t, 'M', ret)
	}

	num = num % 1000
	fmt.Println(num)
	h := num / 100
	switch h {
	case 9:
		ret = ret + "IM"
	}
	return ret
}
