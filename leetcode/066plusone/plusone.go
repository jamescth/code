package main

func plusOne(digits []int) []int {
	carry := 1 // plus one now
	l := len(digits)
	var ret []int

	i := l - 1
	for i >= 0 {
		v := digits[i] + carry
		if v >= 10 {
			v %= 10
			carry = 1
		} else {
			carry = 0
		}
		ret = append([]int{v}, ret...)
		v = 0
		i--
	}
	return ret
}
